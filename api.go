package main

import (
    "github.com/gin-gonic/gin"
)

type Produk struct {
    Nama  string `json:"nama"`
    Harga int    `json:"harga"`
}

func main() {
    r := gin.Default()

    r.GET("/produk", func(c *gin.Context) {
        produk := Produk{
            Nama:  "Teh Sibo",
            Harga: 7000,
        }

        c.JSON(200, produk)
    })

    r.Run() //Menjalankan server di localhost:8080
}
