package main

import (
	_ "github.com/lib/pq"
	"github.com/pamallika/WBL0v2/configs"
	"github.com/pamallika/WBL0v2/internal/service/worker"
)

func main() {
	config := new(configs.Config)
	config.InitFile()
	worker := worker.InitApp(*config)
	worker.Run()
}
