package main

import (
	"github.com/nguyenbatam/op-contract-test/backend/api"
	"github.com/nguyenbatam/op-contract-test/backend/services"
)

func main() {
	server := services.NewService()
	api.StartServer(server, "8080")
}
