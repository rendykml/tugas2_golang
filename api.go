package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Produk struct {
    ID    int     `json:"id"`
    Nama  string  `json:"nama"`
    Harga float64 `json:"harga"`
}

func AmbilProduk() []Produk {
    return []Produk{
        {ID: 1, Nama: "Air Mineral", Harga: 5000},
        {ID: 2, Nama: "Teh Botol", Harga: 6000},
    }
}

func main() {
    r := gin.Default()

    r.GET("/produk", func(c *gin.Context) {
        data := AmbilProduk()
        c.JSON(http.StatusOK, gin.H{
            "status":  "sukses",
            "data":    data,
            "message": "Berhasil mengambil produk",
        })
    })

    r.Run() 
}
