package main

import (
	"fmt"

	cachelru "github.com/IsaqueAmorim/cachelru"
)

func main() {
	cachelru.NewCacheLRU(3)

	cachelru.SetCacheLRU("Isaque", "I-Amorim")
	cachelru.SetCacheLRU("Ameliane", "A-Amorim")
	cachelru.SetCacheLRU("Elienai", "E-Amorim")

	fmt.Println(cachelru.GetCacheLRU("Isaque"))

	cachelru.SetCacheLRU("Gabi", "G-Amorim")

	fmt.Println(cachelru.GetCacheLRU("Isaque"))
}
