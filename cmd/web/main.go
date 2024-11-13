package main

func main() {
	// Create a new server instance
	server := newServer()

	// Start the server and listen for incoming HTTP requests
	server.ListenAndServe()
}
