package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"proxyserver/config"
	"syscall"
	"time"
)

func Create() *http.Server {
	handleSigTerms()

	var port = "8080"
	if len(config.Port) > 0 {
		port = config.Port
	}
	//parse templates
	mux := http.NewServeMux()
	fmt.Println("Starting server on port " + port)
	return &http.Server{
		Addr:           ":" + port,
		Handler:        routes(mux),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func routes(m *http.ServeMux) *http.ServeMux {
	m.HandleFunc("/", response)
	return m
}

func response(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("route")
	body := r.Body
	if key == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}
	// find route
	var route config.Route
	for _, r := range config.Routes {
		if r.Key == key {
			route = r
			break
		}
	}
	if route.Key == "" {
		http.Error(w, "Route not found", http.StatusNotFound)
		return
	}
	// check method
	if r.Method != route.From.Method {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	resp, err := proxyTo(route.To, body)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

}

func proxyTo(route config.To, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(route.Method, route.Url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	return client.Do(req)
}

func handleSigTerms() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("received SIGTERM, exiting")
		os.Exit(1)
	}()
}
