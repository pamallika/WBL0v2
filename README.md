Демонстрация работы программы: https://drive.google.com/file/d/1kW4PxZsy4igL1ng86IhxbTleqyfnoaQD/view?usp=sharing  
База накатывается миграциями:  
migrate -path ./schema/migrations -database "postgres://root:root@localhost:5432/ordersDB?sslmode=disable" up  
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate  
