package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync"
	"time"
)

// Server interface defines methods for a load balancer to interact with backend servers
type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

// simpleserver struct represents a simple backend server with an address and reverse proxy
type simpleserver struct {
	addr  string
	proxy *httputil.ReverseProxy
}

// Constructor function for creating a new simpleserver instance
func newsimplesever(addr string) *simpleserver {
	serverurl, err := url.Parse(addr)
	handleErr(err)
	return &simpleserver{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverurl),
	}
}

// LoadBalancer struct contains the logic for distributing requests to servers
type LoadBalancer struct {
	port       string
	roundrobin int
	servers    []Server
	mutex      sync.Mutex
}

// NewLoadBalancer constructor initializes a LoadBalancer with a given port and list of servers
func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:    port,
		servers: servers,
	}
}

func (s *simpleserver) Address() string {
	return s.addr
}

func (s *simpleserver) IsAlive() bool {
	healthURL := fmt.Sprintf("%s/health", s.addr)
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	resp, err := client.Get(healthURL)
	if err != nil || resp.StatusCode >= 500 {
		return false
	}
	return true
}

// Serve method forwards the HTTP request to the backend server using the reverse proxy
func (s *simpleserver) Serve(rw http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(rw, r)
}

// getNextAvailableServer selects the next available server using round-robin
func (lb *LoadBalancer) getNextAvailableServer() Server {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	serverCount := len(lb.servers)
	for i := 0; i < serverCount; i++ {
		server := lb.servers[lb.roundrobin%serverCount]
		if server.IsAlive() {
			lb.roundrobin = (lb.roundrobin + 1) % serverCount
			return server
		}
		lb.roundrobin = (lb.roundrobin + 1) % serverCount
	}
	return nil // All servers are down
}

// serveProxy method forwards the request to the selected backend server
func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvailableServer()
	if targetServer == nil {
		http.Error(rw, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}
	fmt.Printf("Forwarding request to %s\n", targetServer.Address())
	targetServer.Serve(rw, r)
}

// handleErr function handles any errors, printing the error message and exiting the program
func handleErr(err error) {
	if err != nil {
		fmt.Println("Error occurred:", err)
		os.Exit(1)
	}
}

func main() {
	// List of backend servers to forward traffic to
	servers := []Server{
		newsimplesever("http://localhost:8081"), // Backend server 1
		newsimplesever("http://localhost:8082"), // Backend server 2
		newsimplesever("http://localhost:8083"), // Backend server 3
	}

	lb := NewLoadBalancer("8080", servers) // Create a new load balancer listening on port 8080

	// Handle incoming requests and forward them to the backend servers
	handleRedirect := func(rw http.ResponseWriter, r *http.Request) {
		lb.serveProxy(rw, r) // Use the load balancer to serve the request
	}

	http.HandleFunc("/", handleRedirect)                        
	fmt.Printf("Load Balancer running at http://localhost:%s\n", lb.port) 
	err := http.ListenAndServe(":"+lb.port, nil)                           
	if err != nil {
		fmt.Println("Failed to start load balancer:", err)
	}
}
