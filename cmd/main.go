package main

import (
    "ecommerce-backend-golang/internal/config"
    "ecommerce-backend-golang/internal/routes"
)

func main() {
    config.ConnectDatabase()
    
    r := routes.SetupRouter()
    r.Run(":8080") // Run server on port 8080
}