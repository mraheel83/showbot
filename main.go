package main

import (
  "log"
  "net/http"
  "fmt"

  "github.com/gorilla/mux"
)

func HomeEndpoint(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello from TradeTracker!")
}

/*func VerificationEndPoint(w http.ResponseWritter, r *http.Request)  {
  challenge := r.URL.Query().Get("hub.challenge")
  mode := r.URL.Query().Get("hub.mode")
  token := r.URL.Query().Get("hub.verify_token")

  if mode != "" && token == os.Getenv("VERIFY_TOKEN") {
    w.WriteHeader(200)
    w.Write([]byte(challenge))
  } else {
    w.WriteHeader(400)
    w.Write([]byte("Error", "Wrong validation token"))
  }
}*/

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", HomeEndpoint)
  //r.HandleFunc("/webhook", VerificationEndPoint).Method("GET")
  //r.HanleFunc("/webhook", MessagesEndPoint).Method("POST")
  log.Fatal(http.ListenAndServe(":8080", r))
}
