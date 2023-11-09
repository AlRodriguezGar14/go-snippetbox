package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type config struct {
	Addr      string
	StaticDir string
}

func main() {
	var cfg config

	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP port Addres")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static/", "Path to the static assets")
	// This reads the command line arg and adds it to the declared variable
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(cfg.StaticDir))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", cfg.Addr)
	err := srv.ListenAndServe()
	errLog.Fatal(err)
}
