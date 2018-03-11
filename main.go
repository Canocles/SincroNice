package main

import (
	"fmt"
	"os"
	"sincronice/client"
	"sincronice/server"
)

func main() {
	fmt.Println("Un ejemplo de server/cliente mediante TLS/HTTP en Go.")
	s := "Introduce srv para funcionalidad de servidor y cli para funcionalidad de cliente"

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "srv":
			fmt.Println("Entrando en modo servidor...")
			server.Run()
		case "cli":
			fmt.Println("Entrando en modo cliente...")
			client.Run()
		default:
			fmt.Println("Parámetro '", os.Args[1], "' desconocido. ", s)
		}
	} else {
		fmt.Println(s)
	}
}
