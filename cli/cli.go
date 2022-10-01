package cli

import (
	"awesomeProject/explorer"
	"awesomeProject/rest"
	"flag"
	"fmt"
	"runtime"
)

func usage() {
	fmt.Printf("Welcome to go-coin\n")
	fmt.Printf("Please use the following flags: \n\n")
	fmt.Printf("-restPort: Set the port of REST API\n")
	fmt.Printf("-explorerPort: Set the port of Explorer\n")
	runtime.Goexit()
}

func Start() {

	restPort := flag.Int("restPort", 4000, "Set the port of REST API")
	explorerPort := flag.Int("explorerPort", 3000, "Set the port of Explorer")

	flag.Parse()

	go rest.Start(*restPort)
	explorer.Start(*explorerPort)
}
