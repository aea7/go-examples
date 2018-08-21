package main

import(
  "net/http"
  "github.com/gorilla/websocket"
  "time"
)

var upgrader = websocket.Upgrader{}

func main(){
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w, r, "index.html")
  })

  http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
    var conn, _ = upgrader.Upgrade(w, r, nil)
    go func(conn *websocket.Conn){
      for {
        mType, msg, err :=conn.ReadMessage()
        println(string(msg)) // gönderilen mesajı görmek için
        conn.WriteMessage(mType, msg)

        if err != nil { //client wants to close the websocket
          conn.Close()
        }

        ch := time.Tick(5 *time.Second)
        for range ch {
          conn.WriteJSON(user{
            Name:"Xd",
            Surname:"Xp",
          })
        }

      }
    }(conn)
  })

  http.ListenAndServe(":2000", nil)
}

type user struct {
  Name string `json:"name"`
  Surname string `json:"surname"`
}

//var ws = new WebSocket("ws://localhost:2000/ws")
//ws.addEventListener("message", function(e) {console.log(e);});
//or:  //ws.addEventListener("message", function(e) {console.log(e.data);});
//ws.send("foo")
//ws.close()
