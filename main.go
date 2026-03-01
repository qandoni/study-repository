package main

import (
	simpleconnection "Study/feature_postgres/simple_connection"
	"Study/feature_postgres/simple_sql"
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	conn, err := simpleconnection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	simple_sql.InsertRow(ctx, conn, simple_sql.TaskModel{
		Title:       "Сделать домашку",
		Description: "Сделать домашку по русскому языку",
		Completed:   false,
		CreatedAt:   time.Now(),
	})

	fmt.Println("Успешно")
}
