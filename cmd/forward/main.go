package main

import (
	"flag"
	"fmt"
	"log"
	"http"
)

func main() {

	var port *string = flag.String("port", "8080", "Proxy Port")
	var configFile *string = flag.String("config", "", "Path to config file")
	var verbose *bool = flag.Bool("verbose", false, "Enable verbose logging")

	flag.Parse()

	fmt.Printf("Starting proxy on port: %s\n", *port)

	if *configFile != "" {

		fmt.Printf("Config file at: %s\n", *configFile)

	} else {
		log.Println("No config file found.")
	}

	if *verbose {
		log.Println("Verbose logging enabled.")
	} else {
		log.Println("Verbose logging disabled.")
	}

	var server *http.Server = &http.Server{
		Addr: ":" + *port,
		Handler: &ProxyHandler;
	}

}