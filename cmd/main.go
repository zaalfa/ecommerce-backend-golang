package main

import (
    "ecommerce-backend-golang/internal/routes"
)

func main() {
    r := routes.SetupRouter()

    r.Run(":8080") // Run server on port 8080
}