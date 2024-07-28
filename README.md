# System Monitor with WebSockets and HTMX

## Overview

This project is a web-based system monitoring tool that uses WebSockets to stream live system information to a web interface. It is built with Go, leveraging the `gopsutil` library to gather system metrics and `nhooyr.io/websocket` for WebSocket communication. The web interface is powered by HTMX to handle the live updates.

## Features

- Real-time system monitoring.
- WebSocket-based live data updates.
- Displays system information including CPU, memory, and disk usage.
- Simple and extensible architecture.

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/mapfumo/go-websockets-htmx.git
   cd go-websockets-htmx
   ```

2. Install the required dependencies:

```sh
    go mod tidy
```

## Usage

### Running the Application

To start the application, run:

```sh
make run
```

The server will start on port 8080. Open your web browser and navigate to http://localhost:8080 to view the system monitor.

## Hardware Information Gathering

The `internal/hardware/hardware.go` file contains functions for gathering system information using the gopsutil library:

- GetSystemSection: Gathers system information such as hostname, total and used memory, and OS.
- GetCpuSection: Gathers CPU information such as model name and the number of cores.
- GetDiskSection: Gathers disk usage information such as total and free disk space.

### WebSocket Server

The WebSocket server is implemented in cmd/main.go and handles real-time communication with the clients. The server broadcasts system information to all connected clients every 3 seconds.

### Key Components

- server: Manages subscribers and broadcasts messages.
- subscriber: Represents a WebSocket connection.
- subscribeHandler: Handles new WebSocket connections.
- broadcast: Sends messages to all subscribers.

## Benefits of HTMX

HTMX allows you to extend HTML with attributes to handle AJAX requests, WebSocket messages, and other dynamic content without needing a full-fledged framework like React or Angular. It offers the following benefits:

- Simplicity: Minimalistic approach to dynamic content, reducing complexity.
- Performance: Directly manipulates the DOM, leading to faster interactions.
- Flexibility: Can be used incrementally, making it easy to integrate into existing projects.
- Lightweight: No need for extensive client-side JavaScript frameworks, keeping the client-side footprint small.

## Conclusion

This project demonstrates how to use Go, WebSockets, and HTMX to create a real-time system monitoring tool. The project provides a simple and extensible architecture that can be easily extended to include additional system information and functionality. By leveraging the power of Go and HTMX, developers can create dynamic and interactive web applications that provide real-time updates and improved user experiences.

## Resources

1. [HTMX.org](https://htmx.org/)
2. [Go HTMX Websockets example](https://github.com/sigrdrifa/go-htmx-websockets-example)
