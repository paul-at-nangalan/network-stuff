package api_auth

import (
	"github.com/paul-at-nangalan/errorhandler/handlers"
	"net/http"
	"net/url"
	"testing"
)

func generate()(testfunc http.HandlerFunc, validkey, name string, httphndlr *MockHttpHandler ){

	mockauthmod := &MockSetterGetter{
		apikeys: make(map[string]string),
	}
	name = "test-key"
	authenticator := NewAuthenticator()
	validkey = authenticator.Generate(name, mockauthmod)

	httphndlr = &MockHttpHandler{}

	authfunc := AuthFunc(httphndlr.ServeHTTP, mockauthmod, false)
	return authfunc, validkey, name, httphndlr
}

func Test_ValidKey(t *testing.T){
	testfunc, validye, name, httphandler := generate()

	mockrespwrtr := &MockResponseWriter{}
	testurl := "https://nowhere?" + "authname=" + name + "&authorization=" + url.QueryEscape(validye)
	req, err := http.NewRequest(http.MethodGet, testurl, nil)
	handlers.PanicOnError(err)
	testfunc(mockrespwrtr, req)
	reqrecvdurl := httphandler.reqrecvd.URL
	if req.URL != reqrecvdurl{
		t.Error("Incorrect req in valid key case")
	}
	if mockrespwrtr.status != 200{
		t.Error("Mismatch status, expected 200")
	}
}

func HandleExpectedPanic(t *testing.T){
	if r := recover(); r == nil{
		t.Error("Expected panic, got non")
	}
}
func Test_Invalidkey(t *testing.T){
	testfunc, validye, name, httphandler := generate()

	mockrespwrtr := &MockResponseWriter{}
	[]byte(validye)[3] = 'd'
	[]byte(validye)[6] = 'd'
	testurl := "https://nowhere?" + "authname=" + name + "&authorization=" + url.QueryEscape(validye)
	req, err := http.NewRequest(http.MethodGet, testurl, nil)
	handlers.PanicOnError(err)
	testfunc(mockrespwrtr, req)
	if req.URL != httphandler.reqrecvd.URL{
		t.Error("Incorrect req in valid key case")
	}
}
func Test_Invalidname(t *testing.T){
	defer HandleExpectedPanic(t)
	testfunc, validye, name, httphandler := generate()

	name = "invalid"
	mockrespwrtr := &MockResponseWriter{}
	testurl := "https://nowhere?" + "authname=" + name + "&authorization=" + url.QueryEscape(validye)
	req, err := http.NewRequest(http.MethodGet, testurl, nil)
	handlers.PanicOnError(err)
	testfunc(mockrespwrtr, req)
	if req.URL != httphandler.reqrecvd.URL{
		t.Error("Incorrect req in valid key case")
	}
}
func Test_Noname(t *testing.T){
	defer HandleExpectedPanic(t)
	testfunc, validye, name, httphandler := generate()

	name = ""
	mockrespwrtr := &MockResponseWriter{}
	testurl := "https://nowhere?" + "authname=" + name + "&authorization=" + url.QueryEscape(validye)
	req, err := http.NewRequest(http.MethodGet, testurl, nil)
	handlers.PanicOnError(err)
	testfunc(mockrespwrtr, req)
	if req.URL != httphandler.reqrecvd.URL{
		t.Error("Incorrect req in valid key case")
	}
}
func Test_Nokey(t *testing.T){
	defer HandleExpectedPanic(t)
	testfunc, validye, name, httphandler := generate()

	validye = ""
	mockrespwrtr := &MockResponseWriter{}
	testurl := "https://nowhere?" + "authname=" + name + "&authorization=" + url.QueryEscape(validye)
	req, err := http.NewRequest(http.MethodGet, testurl, nil)
	handlers.PanicOnError(err)
	testfunc(mockrespwrtr, req)
	if req.URL != httphandler.reqrecvd.URL{
		t.Error("Incorrect req in valid key case")
	}
}
func Test_NoNameParam(t *testing.T){
	testfunc, validye, name, _ := generate()

	mockrespwrtr := &MockResponseWriter{}
	testurl := "https://nowhere?" + "auth=" + name + "&authorization=" + url.QueryEscape(validye)
	req, err := http.NewRequest(http.MethodGet, testurl, nil)
	handlers.PanicOnError(err)
	testfunc(mockrespwrtr, req)
	if mockrespwrtr.status != 403{
		t.Error("Mismatch status, expected 404 got ", mockrespwrtr.status)
	}
}
func Test_NoAuthParam(t *testing.T){
	testfunc, validye, name, _ := generate()

	mockrespwrtr := &MockResponseWriter{}
	testurl := "https://nowhere?" + "authname=" + name + "&tion=" + url.QueryEscape(validye)
	req, err := http.NewRequest(http.MethodGet, testurl, nil)
	handlers.PanicOnError(err)
	testfunc(mockrespwrtr, req)
	if mockrespwrtr.status != 403{
		t.Error("Mismatch status, expected 404 got ", mockrespwrtr.status)
	}
}
