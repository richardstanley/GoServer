package main

import (
	"io"
  "fmt"
	"net/http"
  "io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "./client")
}

func fetch(w http.ResponseWriter, r *http.Request) {
  resp, err := http.Get("http://api.eia.gov/series/?api_key=B868A3CF252ABB4CB57A2976DB6B5999&series_id=TOTAL.PAIMPSA.A")
  
  if err != nil{
    fmt.Println(err)
  }

  defer resp.Body.Close()
  content, err := ioutil.ReadAll(resp.Body)

  if err != nil{
    fmt.Println(err)
  }

  fmt.Println("%s\n", string(content))

}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server {
    Addr: ":8000",
    Handler: &myHandler{},
  }

  mux = make(map[string]func(http.ResponseWriter, *http.Request))
  mux["/"] = handler
  mux["/fetch"] = fetch

  fmt.Println("Server is up")

  server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if h, ok := mux[r.URL.String()]; ok {
    h(w, r)
    return
  }

  io.WriteString(w, "My server: "+r.URL.String())
}
