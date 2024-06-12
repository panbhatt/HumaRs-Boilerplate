package main

import (
	"fmt"

	"github.com/panbhatt/HumaRs-Boilerplate/src/internal/config"
)

func main() {

	config.Init()
	fmt.Println(config.Cfg)

}
