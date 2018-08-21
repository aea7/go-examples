package main

import (
	"fmt"
	"net/http"
)

type hotdog int

//handler oluşturmuş oluyorsun ve listenAndServe'e istediğin handlerı(burada hotdog) verince istediğin şekil
//response döndürebiliyorsn.
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
