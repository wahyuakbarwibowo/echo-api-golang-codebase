package main

import "log"

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	server := app.NewServer(cfg)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
