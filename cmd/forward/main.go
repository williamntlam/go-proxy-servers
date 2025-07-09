package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	var port *string = flag.String("port", "8080", "Proxy Port")
	var configFile *string = flag.String("config", "", "Path to config file")
	var verbose *bool = flag.Bool("verbose", false, "Enable verbose logging")

	flag.Parse()

}