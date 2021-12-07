package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	// encadenando middlewares
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, CheckAuth(), Loggin()))
	server.Listen()
}
