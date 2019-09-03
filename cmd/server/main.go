package main

import (
	"fmt"
	"os"

	"github.com/duythien0912/go_todo/pkg/cmd/server"
)

// server -grpc-port=9090 -db-host=<HOST>:3306 -db-user=<USER> -db-password=<PASSWORD> -db-schema=<SCHEMA>

// server -grpc-port=9090 -db-host=127.0.0.1:3306 -db-user=root -db-password=Keypro46@ -db-schema=todo

// with rest

// server -grpc-port=9090 -http-port=8080 -db-host=127.0.0.1:3306 -db-user=root -db-password=Keypro46@ -db-schema=todo

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
