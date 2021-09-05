package main

import (
	"fmt"
	"strings"
	"strconv"
)

func nomorsatu() {
    fmt.Print("Masukkan sebuah kata: ")
	var input string
	fmt.Scanln(&input)

	m := make(map[byte]int)
	for i := 0; i < len(input); i++ {
		m[input[i]] = m[input[i]] + 1
	}
	
	fmt.Printf("Panjang karakter %s yaitu %d \n", input, len(input))
	for key, value := range m {
	    fmt.Printf("Karakter %s yang sama ada %d \n", string(key), value)
	}
}

func nomordua() {
    fmt.Print("Masukkan suhu celcius: ")
    var c int
    fmt.Scanln(&c)
    
    f := (c * 9/5) + 32
    k := c + 273
    
    fmt.Printf("Suhu Celcius: %d, jika Fahrenheit yaitu %d, jika Kelvin yaitu %d \n", c, f, k)
}

func nomortiga() {
    fmt.Print("Masukkan sebuah angka: ")
    var angka int
    fmt.Scanln(&angka)
    
    fmt.Print("\nDeret ganjil: ")
    for i := 0; i <= angka; i++ {
        if i%2 != 0 {
            fmt.Printf("%d ", i)
        }
    }
    
    fmt.Print("\nDeret genap: ")
    for i := 0; i <= angka; i++ {
        if i%2 == 0 {
            fmt.Printf("%d ", i)
        }
    }
}

func nomorempat() {
    fmt.Print("Masukan sebuah angka: ")
    var uang int64
    fmt.Scanln(&uang)
    
    parts := []string{"", "", "", "", "", "", ""}
	i := len(parts) - 1

	for uang > 999 {
		parts[i] = strconv.FormatInt(uang%1000, 10)
		switch len(parts[i]) {
		case 2:
			parts[i] = "0" + parts[i]
		case 1:
			parts[i] = "00" + parts[i]
		}
		uang = uang / 1000
		i--
	}
	parts[i] = strconv.Itoa(int(uang))
	rupiah := "Rp" + strings.Join(parts[i:], ".")
	dollar := "$" + strings.Join(parts[i:], ".")

    fmt.Printf("Jika menjadi Rupiah: %s \n", rupiah)
    fmt.Printf("Jika menjadi Dollar: %s \n", dollar)
}

func nomorlima() {
    fmt.Println("Masukkan panjang: ")  
    var panjang int
    fmt.Scanln(&panjang)
    
    fmt.Println("Masukan lebar: ")
    var lebar int
    fmt.Scanln(&lebar)
    
    k := 2 * (panjang + lebar)
    l := panjang * lebar
    fmt.Printf("Luas persegi panjang adalah %d, dan Keliling persegi panjang adalah %d", l, k)
}

func main() {
    nomorsatu()
    nomordua()
    nomortiga()
    nomorempat()
    nomorlima()
}
