package main

import (
	"fmt"
	"github.com/syhlion/gwspack"
	"log"
	"net/http"
)

type Hello struct{}

func (h *Hello) Receive(tag string, s gwspack.Sender, b []byte, data map[string]interface{}) {
	log.Println(tag)
	s.SendAll(b)
}
func main() {

	h := &Hello{}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		app := gwspack.New("key", h, 10)
		ws, err := app.Register("Frank", w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		ws.Listen()
		log.Println("socket end")
	})
	log.Fatal(http.ListenAndServe(":8888", nil))
}
