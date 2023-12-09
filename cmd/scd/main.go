package main

import (
	"fmt"
	"github.com/Q00/go-chat/config"
	"github.com/Q00/go-chat/graphql"
	"github.com/Q00/go-chat/graphql/docs"
	graphqlGo "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	PORT = 8080
)

func main() {
	env := os.Getenv("GO_ENV")
	fmt.Printf("chat node starts, env:  %v\n", env)
	// load config
	cfg := config.LoadConfig()

	e := echo.New()
	schemaCompiled := graphqlGo.MustParseSchema(graphql.Schema, graphql.NewResolver(&cfg.AWS))
	graphqlHandler := graphqlws.NewHandlerFunc(schemaCompiled, &relay.Handler{Schema: schemaCompiled})

	echoGraphqlHandler := func(c echo.Context) error {
		graphqlHandler(c.Response().Writer, c.Request())
		return nil
	}

	g := docs.NewGraphiql(PORT)
	e.Use(m.Recover())
	e.Use(CORS)
	e.Any("/graphql", func(c echo.Context) error {
		if c.Request().Header.Get("Upgrade") == "websocket" {
			// Handle WebSocket GraphQL connection
			return echoGraphqlHandler(c)
		} else if c.Request().Method == "GET" {
			return g.Start()(c)
		} else {
			// Handle regular HTTP GraphQL request
			return echoGraphqlHandler(c)
		}
	})

	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", PORT),
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 120,
	}

	log.Fatal(e.StartServer(server))
	log.Fatal(server.ListenAndServe())
}

// CORS allows us to customize the CORS headers
func CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		CORSMiddleware := m.CORSWithConfig(m.CORSConfig{
			AllowOrigins: []string{"http://localhost:3000"},
		})
		CORSHandler := CORSMiddleware(next)
		return CORSHandler(c)
	}
}
