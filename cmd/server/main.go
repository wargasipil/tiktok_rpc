package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/urfave/cli/v2"
)

func CorsMiddleware(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Header("Access-Control-Allow-Methods", "HEAD,PATCH,OPTIONS,GET,POST,PUT,DELETE")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}

func runServer(host string) error {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	// Create Setup
	router := gin.Default()
	router.Use(CorsMiddleware)
	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		for {
			conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
			time.Sleep(time.Second)
		}
	})

	log.Println("running on", host)
	err := router.Run(host)

	return err
}

func main() {
	app := &cli.App{
		Name: "wargasipil rpc server tiktok",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "address",
				Aliases:     []string{"addr"},
				DefaultText: "localhost:8081",
				Value:       "localhost:8081",
			},
		},
		Action: func(ctx *cli.Context) error {
			host := ctx.String("address")
			return runServer(host)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
