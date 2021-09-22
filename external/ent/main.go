package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"go_packages/external/ent/ent"
	"log"
)

func main() {
	client, err := ent.Open("mysql", "hjm_dev:hjm_dev@tcp(h5m:3306)/gb?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to splite: %v", err)
	}
	defer client.Close()

	// run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
