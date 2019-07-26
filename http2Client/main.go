package main

import (
	"bytes"
	"fmt"
	"http2ServerClient/http2Protocol"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	host := "127.0.0.1:9998"
	// http2 request
	client := NewClient()
	sendHttp2Requeest(client, host, "/test", "YT")
	sendHttp2Requeest(client, host, "/test", "SJ")
	sendHttp2Requeest(client, host, "/test", "JH")

	// http1.1 request
	sendHttp1Request(host, "/test", "GOGO")
}

func sendHttp1Request(host, path string, name string) {
	url := fmt.Sprintf("http://%s%s", host, path)
	req := http2Protocol.TestRequest{}
	req.Msg = name
	data, _ := req.MarshalJSON()

	response, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	defer func(){
		if response != nil {
			response.Body.Close()
		}
	}()

	if err != nil {
		log.Println("Http request failed", err.Error())
		return
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	res := http2Protocol.TestResponse{}
	err = res.UnmarshalJSON(contents)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("[http/1.1] I'm "+name, res.Msg)
	}
}

func sendHttp2Requeest(client *Client, host, path, name string) {
	req := http2Protocol.TestRequest{}
	req.Msg = name
	bytes, _ := req.MarshalJSON()
	resp, err := client.Post(host, path, bytes)

	if err != nil {
		log.Println(err)
	} else {
		res := http2Protocol.TestResponse{}
		err = res.UnmarshalJSON(resp)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("[http/2] I'm "+ name, res.Msg)
		}
	}
}
