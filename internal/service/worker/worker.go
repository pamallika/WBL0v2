package worker

import (
	"context"
	"github.com/nats-io/stan.go"
	"github.com/pamallika/WBL0v2/configs"
	"github.com/pamallika/WBL0v2/internal/repository/database"
	"github.com/pamallika/WBL0v2/internal/service/cache"
	"github.com/pamallika/WBL0v2/internal/service/consumer"
	"github.com/pamallika/WBL0v2/internal/service/server"
	"github.com/pamallika/WBL0v2/internal/service/store"
	"log"
	"net/http"
	"time"
)

type App struct {
	cfg configs.Config
}

func InitApp(cfg configs.Config) *App {
	app := App{}
	app.cfg = cfg
	return &app
}

func (app *App) Run() {
	db, err := database.InitDBConn(app.cfg)
	if err != nil {
		log.Fatalf("Error connnecting to database : %s", err)
	}

	cacheService := cache.CacheInit()

	storeService := store.InitStore(*cacheService, *db)

	err = storeService.RestoreCache()

	if err != nil {
		log.Println("error restoring cache: db is empty")
	}

	sc := consumer.CreateSub(*storeService)
	err = sc.Connect(app.cfg.Nats_server.Cluster_id, app.cfg.Nats_server.Client_id, app.cfg.Nats_server.Host+":"+app.cfg.Nats_server.Port)

	if err != nil {
		log.Printf("Error connecting to STAN: %s", err)
	}

	sub, err := sc.SubscribeToChannel(app.cfg.Nats_server.Channel, stan.StartWithLastReceived())

	if err != nil {
		log.Printf("Error subscribing to channel : %s", err)
	}

	server := server.InitServer(*storeService, app.cfg.Http_server.Host+":"+app.cfg.Http_server.Port)

	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err)
		}
	}()
	log.Print("Server Started")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
		defer server.Stop()
		defer sub.Unsubscribe()
		defer sc.Close()
		defer db.Close()
	}()

	if err := server.Srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed : %s", err)
	}
	log.Print("Server Stopped")
}
