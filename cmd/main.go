package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/mapfumo/go-websockets-htmx/internal/hardware"
	websocket "nhooyr.io/websocket"
)

// server struct to handle subscribers and HTTP routing
type server struct {
	subscriberMessageBuffer int                 // buffer size for each subscriber's message channel
	mux                     http.ServeMux       // HTTP request multiplexer
	subscribersMutex        sync.Mutex          // mutex to ensure safe access to the subscribers map
	subscribers             map[*subscriber]struct{} // map to keep track of subscribers
}

// subscriber struct to represent a single WebSocket connection
type subscriber struct {
	msgs chan []byte // channel to send messages to the subscriber
}

// NewServer initializes a new server instance
func NewServer() *server {
	s := &server{
		subscriberMessageBuffer: 10,                        // set buffer size for subscriber message channel
		subscribers:             make(map[*subscriber]struct{}), // initialize the subscribers map
	}
	s.mux.Handle("/", http.FileServer(http.Dir("./htmx"))) // serve static files from the "htmx" directory
	s.mux.HandleFunc("/ws", s.subscribeHandler)           // handle WebSocket connections at the "/ws" endpoint
	return s
}

// subscribeHandler handles HTTP requests for WebSocket connections
func (s *server) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	err := s.subscribe(r.Context(), w, r) // call the subscribe method to handle the WebSocket connection
	if err != nil {
		fmt.Println(err) // print any errors
		return
	}
}

// addSubscriber adds a new subscriber to the subscribers map
func (s *server) addSubscriber(subscriber *subscriber) {
	s.subscribersMutex.Lock()   // lock the mutex to ensure thread-safe access to the subscribers map
	s.subscribers[subscriber] = struct{}{} // add the subscriber to the map
	s.subscribersMutex.Unlock() // unlock the mutex
}

// subscribe handles the WebSocket connection for a new subscriber
func (s *server) subscribe(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var c *websocket.Conn
	subscriber := &subscriber{
		msgs: make(chan []byte, s.subscriberMessageBuffer), // create a new subscriber with a buffered message channel
	}
	s.addSubscriber(subscriber) // add the subscriber to the server
	c, err := websocket.Accept(w, r, nil) // accept the WebSocket connection
	if err != nil {
		return err // return the error if the connection fails
	}
	defer c.CloseNow() // ensure the connection is closed when the function exits

	ctx = c.CloseRead(ctx) // set the context to be done when the WebSocket connection is closed for reading
	for {
		select {
		case msg := <-subscriber.msgs: // receive a message from the subscriber's message channel
			ctx, cancel := context.WithTimeout(ctx, time.Second) // create a new context with a timeout
			defer cancel()
			err := c.Write(ctx, websocket.MessageText, msg) // write the message to the WebSocket connection
			if err != nil {
				return err // return the error if the write fails
			}
		case <-ctx.Done(): // if the context is done (connection closed or timeout)
			return ctx.Err() // return the context error
		}
	}
}

// broadcast sends a message to all subscribers
func (s *server) broadcast(msg []byte) {
	s.subscribersMutex.Lock() // lock the mutex to ensure thread-safe access to the subscribers map
	for subscriber := range s.subscribers {
		subscriber.msgs <- msg // send the message to each subscriber's message channel
	}
	s.subscribersMutex.Unlock() // unlock the mutex
}

func main() {
	fmt.Println("Starting system monitor...")
	srv := NewServer() // create a new server instance
	go func(s *server) {
		for {
			systemData, err := hardware.GetSystemSection() // get system data
			handleError(err)
			cpuSection, err := hardware.GetCpuSection() // get CPU data
			handleError(err)
			diskSection, err := hardware.GetDiskSection() // get disk data
			handleError(err)

			timeStamp := time.Now().Format("02-01-2006 15:04:05") // get the current timestamp

			// create an HTML string with the collected data
			html := `
			<div hx-swap-oob="innerHTML:#update-timestamp">` + timeStamp + `</div>
			<div hx-swap-oob="innerHTML:#system-data">` + systemData + `</div>
			<div hx-swap-oob="innerHTML:#cpu-data">` + cpuSection + `</div>
			<div hx-swap-oob="innerHTML:#disk-data">` + diskSection + `</div>
			`

			s.broadcast([]byte(html)) // broadcast the HTML to all subscribers

			time.Sleep(3 * time.Second) // wait for 3 seconds before repeating
		}
	}(srv)

	err := http.ListenAndServe(":8080", &srv.mux) // start the HTTP server on port 8080
	handleError(err)
}

// handleError prints the error if it is not nil
func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
