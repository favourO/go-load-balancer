package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)

	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type LoadBalancer struct {
	port           string
	rounRobinCount int
	servers        []Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:           port,
		rounRobinCount: 0,
		servers:        servers,
	}
}
func handleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func Address(s *simpleServer) string { return s.addr }

func (s *simpleServer) IsAlive() bool { return true }


func (lb *LoadBalancer) getNextAvailableServer() Server {}

func (lb *LoadBalancer) serverProxy(rw http.ResponseWriter, r *http.Request) {}

func main() {
	servers := []Server{
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("https://www.bing.com"),
		newSimpleServer("https://www.apple.com"),
	}

	lb := NewLoadBalancer("8080", servers)

	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serverProxy(rw, req)
	}

	http.HandlerFunc("/", handleRedirect)

	fmt.Printf("serving request at 'localhost:%s'\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}