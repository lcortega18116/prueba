package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main2() {
	dsn := "postgresql://luis:luk65vrf5QovnAW3Vltp-w@tiring-muskox-10220.jxf.gcp-us-east1.cockroachlabs.cloud:26257/dbluis?sslmode=verify-full"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	var now time.Time
	err = db.QueryRow("SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatal("failed to execute query", err)
	}

	fmt.Println(now)
}
