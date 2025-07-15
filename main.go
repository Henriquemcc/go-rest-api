package main

import (
	"fmt"

	"github.com/Henriquemcc/go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
