package main

import (
    "fmt"
)

// Struct
type Produk struct {
    Nama  string
    Harga int
}

// Fungsi
func TambahProduk(p Produk) string {
    return fmt.Sprintf("Nama produk: %s, Harga: %d",p.Nama,p.Harga)
}

func main() {
    p := Produk{
        Nama:  "Teh Sibo",
        Harga: 7000,
    }

    hasil := TambahProduk(p)
    fmt.Println(hasil)
}
