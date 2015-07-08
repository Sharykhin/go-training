package main

import (
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	users := make(map[string]string)
	server.On("connection", func(so socketio.Socket) {
		users[so.Id()] = ""
		log.Println("on connection", so.Id())
		log.Println(users)
		so.Join("chat")
		so.On("enter name", func(msg string) {
			log.Println("emit:", so.Emit("enter name", msg))
			users[so.Id()] = msg
			//so.BroadcastTo("chat", "chat message", msg)
			log.Println(users)
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
			delete(users, so.Id())
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public/")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
