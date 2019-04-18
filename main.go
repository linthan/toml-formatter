package main

import (
	"os"

	"github.com/pelletier/go-toml"
)

func main() {
	config, _ := toml.LoadFile("example.toml")
	// fmt.Println()
	f, _ := os.Create("aa.toml")
	config.WriteTo(f)
}
