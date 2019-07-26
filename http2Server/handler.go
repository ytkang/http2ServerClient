package main

import (
	"http2ServerClient/http2Protocol"
	"io/ioutil"
	"log"
	"net/http"
)

type Multiplexer struct {
	handlers map[string]func(w http.ResponseWriter, r *http.Request)
	Handler http.HandlerFunc
}

func NewMultiplexer() *Multiplexer {
	multiplexer := &Multiplexer{
		handlers: make(map[string] func(w http.ResponseWriter, r *http.Request)),
	}
	multiplexer.init()
	return multiplexer
}

func (mux *Multiplexer) init() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// you can allow only http2 request
		//if r.ProtoMajor != 2 {
		//	log.Println("HTTP/2 request is rejected")
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}

		f := mux.GetHandler(r.URL.Path)
		if f == nil {
			log.Println("unknown path", r.URL.Path)
			return
		}
		f(w,r)
	})
	mux.Handler = handler
}

func (mux *Multiplexer) GetHandler(path string) func(w http.ResponseWriter, r *http.Request) {
	f, ok := mux.handlers[path]
	if ok {
		return f
	}
	return nil
}

func (mux *Multiplexer) HandleFunc(path string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.handlers[path] = f
}

func TestRequestHandler(w http.ResponseWriter, r *http.Request) {
	req := http2Protocol.TestRequest{}
	res := http2Protocol.TestResponse{}

	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("read error", err)
		res.Msg = "socket read error"
	} else {
		err = req.UnmarshalJSON(contents)
		if err != nil {
			res.Msg = "invalid request"
		} else {
			res.Msg = "Hello!, " + req.Msg
		}
	}
	bytes, _ := res.MarshalJSON()
	w.Write(bytes)
}
