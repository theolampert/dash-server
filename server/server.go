package server

import (
	"github.com/goware/cors"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"log"
	"net/http"
)

func Run(config map[string]string) {

	mediaDir := config["workDir"]

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "OPTIONS"},
		AllowedHeaders: []string{
			"Access-Control-Request-Headers",
			"Origin",
			"Range",
		},
		ExposedHeaders:   []string{"Link", "Server", "range"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler)

	r.FileServer("/", http.Dir(mediaDir))

	if config["pem"] != "" && config["key"] != "" {
		err := http.ListenAndServeTLS(config["port"], config["pem"], config["key"], r)

		if err != nil {
			log.Fatal("ListenAndServeTLS: ", err)
		}
	} else {
		err := http.ListenAndServe(config["port"], r)

		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}

}
