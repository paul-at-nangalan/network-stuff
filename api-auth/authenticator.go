package api_auth

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/paul-at-nangalan/errorhandler/handlers"
	"log"
	"net/http"
)
import "golang.org/x/crypto/bcrypt"

type AuthModel interface {
	Get(name string) string
}
type AuthModelSetter interface {
	Set(name, apikey string)
}

/**
Usage:
import "github.com/paul-at-nangalan/network-stuff/api-auth"
import "github.com/paul-at-nangalan/db-util/connect"

//Create a http handler
type Handler struct{

}

///Implement the handler func
func (p *Handler)HandleReq(w http.ResponseWriter, req *http.Request) {

}

////somewhere , normally the constructor, set up the handler func wrapped in the authenticaotr
/// replace Handler with the name of your type
func NewHandler()*Handler{

	handler := &Handler{}

	//// Create the connection to the DB - or use your own DB pointer
	db := connect.Connect()

	////Create the auth DB model - or use your own model - it just needs to obey the interface
	authmod := models.NewApikeysModel(db)

	////set up the route
	http.HandleFunc("/your/webpage", api_auth.AuthFunc(handler.HandleReq, authmod, true))
}
 */
type Authenticator struct{
	handlerfunc http.HandlerFunc
	authmod AuthModel
	keylen int
}

////ONLY use this for generating keys, for checking use AuthFunc (see comments above)
func NewAuthenticator()*Authenticator{
	return &Authenticator{
		keylen: 90,
	}
}

func AuthFunc(handlerFunc http.HandlerFunc, authmod AuthModel, usehdr bool)http.HandlerFunc{
	auth := &Authenticator{
		handlerfunc: handlerFunc,
		authmod: authmod,
		keylen: 120,
	}
	if usehdr {
		return auth.HandleReq
	}else{
		return auth.HandleReqByQuery
	}
}

func (p *Authenticator)HandleReq(w http.ResponseWriter, req *http.Request) {
	authheader := req.Header.Get("Authorization")
	name := req.Header.Get("Authname")
	p.handleReq(name, authheader, w, req)
}
func (p *Authenticator)HandleReqByQuery(w http.ResponseWriter, req *http.Request) {
	authheader := req.URL.Query().Get("authorization")
	name := req.URL.Query().Get("authname")
	p.handleReq(name, authheader, w, req)
}

func (p *Authenticator)handleReq(name, authheader string, w http.ResponseWriter, req *http.Request){
	if authheader == ""{
		http.Error(w, "No authorization header", http.StatusForbidden)
		return
	}
	if name == ""{
		http.Error(w, "No authname header", http.StatusForbidden)
		return
	}
	apikey := p.authmod.Get(name)
	err := bcrypt.CompareHashAndPassword([]byte(apikey), []byte(authheader))
	if err != nil{
		log.Println(err)
		http.Error(w, "Invalid token", http.StatusForbidden)
		return
	}
	/// if we get here it should be ok
	p.handlerfunc(w, req)
}

func (p *Authenticator)Try(name, key string)error{
	apikey := p.authmod.Get(name)
	err := bcrypt.CompareHashAndPassword([]byte(apikey), []byte(key))
	return err
}

func (p *Authenticator)Generate(name string, authmod AuthModelSetter)string{
	b := make([]byte, p.keylen)
	n, err := rand.Read(b)
	if n != len(b) || err != nil{
		log.Panicln("Failed to gen random string:", n, err)
	}
	str := base64.StdEncoding.EncodeToString(b)
	encstr, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	handlers.PanicOnError(err)
	authmod.Set(name, string(encstr))

	return str
}