package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", HandleRoot)
	// encadenando middlewares
	server.Handle("/api", server.AddMiddleware(HandleHome, CheckAuth(), Loggin()))
	server.Listen()
}
