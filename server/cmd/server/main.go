package main

import (
	"flag"
	"fmt"
	"github.com/rs/cors"
	"github.com/rusmDocs/rusmDocs/internal/configs"
	"github.com/rusmDocs/rusmDocs/internal/handlers"
	"github.com/rusmDocs/rusmDocs/pkg/config"
	"github.com/rusmDocs/rusmDocs/pkg/database"
	"log"
	"net/http"
)

func InitServer(config configs.ServerConfig) error {
	r := handlers.InitRoute(config)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	var isMigrate bool
	flag.BoolVar(&isMigrate, "migrate", false, "make migrations")
	flag.Parse()

	if isMigrate {
		database.MakeMigration()
	}

	return http.ListenAndServe(
		config.App.Host+":"+config.App.Port,
		c.Handler(r),
	)
}

func main() {
	config.Configure("config.yml")

	fmt.Println("Server is running on port", config.Config.App.Port)
	if err := InitServer(config.Config); err != nil {
		log.Panic(err)
	}

	fmt.Print("serverConfig", config.Config)
}
