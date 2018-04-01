package main

import (
	"context"
	"fmt"
	"net/http"
	"github.com/graphql-go/handler"
	models "github.com/SilentMouse/stock_market_data/pkg/models"
	st_schema "github.com/SilentMouse/stock_market_data/pkg/graphql"
	"github.com/Sirupsen/logrus"
)

func main() {
	db := models.InitDB()
	st_schema.InitSchema(&db)

	h := handler.New(&handler.Config{
		Schema:   &st_schema.StSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	fmt.Println("Now server is running on port 5000")
	logrus.Infoln("logrus info")
	fmt.Println("Test with Get      : curl -g 'http://localhost:5000/graphql?query={spheres{name}}'")
	http.Handle("/graphql", add_headers(h))

	http.ListenAndServe(":6000", nil)

}

func add_headers(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("X-Token")
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "X-Token, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx := context.WithValue(r.Context(), "auth", authorizationHeader)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}