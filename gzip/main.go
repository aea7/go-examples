package main

import (
  "net/http"
  "compress/gzip"
  "io"
  "strings"
)

func main()  {
  http.Handle("/", http.FileServer(http.Dir("static")))

  go func() {
			http.ListenAndServe(":4443", nil)
	}()

  http.ListenAndServe(":4445", new(gzipHandler))
}

type gzipHandler struct {

}

func (gzh *gzipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
  if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip"){
    http.DefaultServeMux.ServeHTTP(w, r)
    return
  }
  w.Header().Add("Content-Encoding", "gzip")
  gzw := gzip.NewWriter(w)
  defer gzw.Close()
  gzrw := &gzipResponseWriter{
    ResponseWriter: w,
    Writer: gzw,
  }
  http.DefaultServeMux.ServeHTTP(gzrw, r)
}

type gzipResponseWriter struct {
  http.ResponseWriter
  io.Writer
}

func (gzrw gzipResponseWriter) Write(data []byte) (int, error){
  return gzrw.Writer.Write(data)
}
