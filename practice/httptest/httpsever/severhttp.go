package httpsever

import (
	"io"
	"net/http"
)

var HandlersMap = make(map[string]HandlersFunc)

type HandlersFunc func(http.ResponseWriter, *http.Request)
type WebSever struct{}

func (ws *WebSever) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := HandlersMap[r.URL.String()]; ok {
		h(w, r)
	}
}
func F1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "111111111111")
}

func F2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "2222222222222")
}
func HelloW(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello China!"))
}
