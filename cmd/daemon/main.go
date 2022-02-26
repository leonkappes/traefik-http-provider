package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"syscall"

	"github.com/leonkappes/go-traefik-daemon/internal/routes"
	"github.com/leonkappes/go-traefik-daemon/pkg/config"
	"github.com/leonkappes/go-traefik-daemon/pkg/http"
	"github.com/sevlyar/go-daemon"
	"github.com/valyala/fasthttp"
)

var (
	signal = flag.String("s", "", "")
	stop   = make(chan struct{})
	done   = make(chan struct{})
)

func main() {
	flag.Parse()
	cfg, _ := config.ParseConfig()
	daemon.AddCommand(daemon.StringFlag(signal, "quit"), syscall.SIGQUIT, termHandler)

	cntxt := &daemon.Context{
		PidFileName: "go-traefik.pid",
		PidFilePerm: 0644,
		LogFileName: "go-traefik.log",
		LogFilePerm: 0644,
		WorkDir:     "./",
		Umask:       027,
	}

	if len(daemon.ActiveFlags()) > 0 {
		d, err := cntxt.Search()
		if err != nil {
			log.Fatalf("Unable send signal to the daemon: %s", err.Error())
		}
		daemon.SendCommands(d)
		return
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatalln(err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	go worker(cfg)

	err = daemon.ServeSignals()
	if err != nil {
		log.Fatalln(err)
	}
}

func worker(cfg *config.Config) {
	ln, err := net.Listen("tcp4", fmt.Sprintf("127.0.0.1:%d", cfg.Port))
	if err != nil {
		log.Fatalf("Error in net.Listen: %S", err)
	}

	router := http.New()

	registerRoutes(*router)

	go func() {
		log.Printf("Listening on Port %d", cfg.Port)
		if err := fasthttp.Serve(ln, router.Handler); err != nil {
			log.Fatalf("error in Serve: %s", err)
		}
	}()

	log.Println("Waiting for interrupt")
	<-stop
	log.Println("Interrupt received")
	ln.Close()
	done <- struct{}{}
}

func registerRoutes(router http.Router) {
	router.AddRoute(http.NewRoute("GET", "/config", routes.ProviderHandler))
}

func termHandler(sig os.Signal) error {
	stop <- struct{}{}
	if sig == syscall.SIGQUIT {
		<-done
		log.Println("Gracefully shutdown")
	}
	return daemon.ErrStop
}
