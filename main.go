package main

import (
	"awesomeProject/cli"
	"awesomeProject/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
