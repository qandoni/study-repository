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

	// tasks, err := simple_sql.SelectRows(ctx, conn)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, task := range tasks {
	// 	if task.ID == 4 {
	// 		task.Title = "Покормить кошку"
	// 		task.Description = "Отсыпать кошке 30 грамм корма"
	// 		task.Completed = true
	// 		now := time.Now()
	// 		task.CompletedAt = &now

	// 		if err := simple_sql.UpdateTask(ctx, conn, task); err != nil {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// }

	fmt.Println("Успешно")
}
