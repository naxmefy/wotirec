package main

import (
    "flag"
    "log"
    "net/http"
    "time"

    "github.com/gorilla/mux"
)

func main() {
    var entry string
    var static string

    var host string
    var port string

    flag.StringVar(&entry, "entry", "./client/build/index.html", "the entrypoint to serve.")
    flag.StringVar(&static, "static", "./client/build/", "the directory to serve static files from.")
    flag.StringVar(&port, "host", "", "the `host` to listen on.")
    flag.StringVar(&port, "port", "8000", "the `port` to listen on.")
    flag.Parse()

	r := mux.NewRouter()

    // This will serve files under http://localhost:8000/static/<filename>
    r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(static))))

	server := &http.Server{
	    Handler: r,
	    Addr: host+":"+port,

	    WriteTimeout: 15 * time.Second,
	    ReadTimeout: 15 * time.Second,
    }

	log.Print("listening on ", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func IndexHandler(entryPoint string) func(w http.ResponseWriter, r *http.Request) {
    fn := func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, entryPoint)
    }

    return fn
}
