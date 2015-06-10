package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection", so.Id())

		so.Join("chat")
		so.Emit("users", so.Id())
		so.BroadcastTo("chat", "users", so.Id())
		so.Join(so.Id())
		so.On("chat message", func(msg string) {
			log.Println("emit:", msg)
			//so.BroadcastTo("chat message", msg, msg)
			so.Emit("chat message", msg)
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
			so.Leave(so.Id())
			so.BroadcastTo("chat", "leave", so.Id())
		})

		so.On("create", func(room string) {
			log.Println(room)

			so.Join(room)
			so.Emit("create", room)
			so.BroadcastTo("chat", "create", room)
			log.Print(room + " message")
			so.On(room+" message", func(msg string) {
				log.Println(room, "message", msg)
			})
		})
		so.On("join", func(room string) {
			log.Println("Join to", room)
			so.Join(room)
		})

	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
