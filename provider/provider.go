package provider

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/bitcodr/avn-promotion/config"
	"github.com/bitcodr/avn-promotion/handler"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Init() {
	app := new(config.App)
	app.Init()

	err := make(chan error, 3)

	go func() {
		grpcListener(app, err)
	}()

	go func() {
		httpListener(app, err)
	}()

	natsListener(app)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		err <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("terminated %s", <-err)
}

func httpListener(app *config.App, err chan<- error) {
	router := mux.NewRouter()

	handler.HTTP(app, router)

	router.Use(mux.CORSMethodMiddleware(router))

	fmt.Println("Listening HTTP on port " + os.Getenv("HTTP_PORT"))

	err <- http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), router)
}

func grpcListener(app *config.App, err chan<- error) {
	listener, er := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if er != nil {
		log.Fatal(er)
	}
	server := grpc.NewServer()

	handler.GRPC(app, server)

	reflection.Register(server)

	fmt.Println("Listening TCP on port " + os.Getenv("GRPC_PORT"))

	err <- server.Serve(listener)
}

func natsListener(app *config.App) {
	c, er := config.NATSClient()
	if er != nil {
		log.Fatal(er)
	}

	fmt.Println("Listening NATS...")

	handler.NATS(app, c)

	runtime.Goexit()
}
