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

	"github.com/amiraliio/avn-promotion/config"
	"github.com/amiraliio/avn-promotion/handler"
	"github.com/gorilla/mux"
	nats "github.com/nats-io/nats.go"
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

	natsListener()

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

	fmt.Println("Listening HTTP on port " + config.AppConfig.GetString("APP.HTTP_PORT"))

	err <- http.ListenAndServe(":"+config.AppConfig.GetString("APP.HTTP_PORT"), router)
}

func grpcListener(app *config.App, err chan<- error) {
	listener, er := net.Listen("tcp", ":"+config.AppConfig.GetString("APP.TCP_PORT"))
	if er != nil {
		log.Fatal(er)
	}
	server := grpc.NewServer()

	handler.GRPC(app, server)

	reflection.Register(server)

	fmt.Println("Listening TCP on port " + config.AppConfig.GetString("APP.TCP_PORT"))

	err <- server.Serve(listener)
}

func natsListener() {
	fmt.Println("Listening NATS...")
	c, er := config.NATSClient()
	if er != nil {
		log.Fatal(er)
	}
	c.Subscribe("promotion.*", func(m *nats.Msg) {
		fmt.Println("gb")
		fmt.Println(string(m.Data))
		fmt.Println(m.Reply)
		fmt.Println(m.Sub)
		fmt.Println(m.Subject)
		c.Publish(m.Reply, "I can help!")
	})
	runtime.Goexit()
}
