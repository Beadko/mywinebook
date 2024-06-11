/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/Beadko/mywinebook/endpoints"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "atarts a server",
	Run: func(cmd *cobra.Command, args []string) {
		RunServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func getCorsAllowedOrigin() string {
	envContent := os.Getenv("CORS_ALLOWED_ORIGIN")
	if envContent == "" {
		envContent = "http://localhost:8080"
	}
	return envContent
}

func RunServer() {
	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{getCorsAllowedOrigin()},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	endpoints.AddRouterEndpoints(router)
	err := http.ListenAndServe(":8081", handler)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
