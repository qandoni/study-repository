package main

import (
	httpserver "Study/http_server"
	"fmt"
)

func main() {
	fmt.Println("HTTP server started!")

	err := httpserver.StartHTTPServer()
	if err != nil {
		fmt.Println("HTTP server error:", err)
	} else {
		fmt.Println("HTTP stopped")
	}
}
