package api_auth

import (
	"fmt"
	"log"
	"net/http"
)

type MockSetterGetter struct{
	apikeys map[string]string
}

func (p *MockSetterGetter) Get(name string) string {
	fmt.Println("Looking for name ", name)
	key, ok := p.apikeys[name]
	if !ok{
		fmt.Println("Not ok, panic")
		log.Panicln("Invalid api key")
	}
	fmt.Println("Found key ", key)
	return key
}

func (p *MockSetterGetter) Set(name, apikey string) {
	p.apikeys[name] = apikey
}

type MockHttpHandler struct{
	reqrecvd *http.Request
}

func (p *MockHttpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Recvd request ", request)
	p.reqrecvd = request
	writer.WriteHeader(200)
}

type MockResponseWriter struct{
	 status int
}

func (p *MockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (p *MockResponseWriter) Write(bytes []byte) (int, error) {
	return len(bytes), nil
}

func (p *MockResponseWriter) WriteHeader(statusCode int) {
	p.status = statusCode
}



