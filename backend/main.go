package main

func main() {
	server := NewAPIServer(":1337")
	server.Run()
}
