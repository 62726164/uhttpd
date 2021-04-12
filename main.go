package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
)

// basic logging to see the method and path requested by clients
func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func main() {
	var socket = flag.String("socket", "/tmp/httpd.sock", "the path to the unix domain socket")
	var webroot = flag.String("webroot", "/tmp/webroot.", "the path to the webroot")
	var help = flag.Bool("help", false, "show help.")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	os.Remove(*socket)

	server := http.Server{
		Handler: logHandler(http.FileServer(http.Dir(*webroot))),
	}

	udsListener, err := net.Listen("unix", *socket)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.Chmod(*socket, 0777); err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.Serve(udsListener))
}
