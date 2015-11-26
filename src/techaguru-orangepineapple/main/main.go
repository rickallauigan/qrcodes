package main

import (
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/astaxie/beegae"
	//"golang.org/x/net/context"
	"log"
	"net/http"
	"sync"
	"techaguru-orangepineapple/controllers"
	_ "techaguru-orangepineapple/routers"
	//"time"
)

var lock = sync.RWMutex{}

func init() {
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("../static/css"))))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", http.FileServer(http.Dir("../static/js"))))
	http.Handle("/static/components/js/", http.StripPrefix("/static/components/js/", http.FileServer(http.Dir("../static/components/js"))))
	http.Handle("/static/fonts/", http.StripPrefix("/static/fonts/", http.FileServer(http.Dir("../static/fonts"))))

	//beegae.Router("/api/v1/create", &controllers.RestController{}, "post:PostData")

	beegae.Router("/api/v1/create", &controllers.MainController{}, "post:Post")
	beegae.Router("/", &controllers.MainController{}, "get:Get")
	beegae.Run()

}

func datastore_INSERT(w http.ResponseWriter, r *http.Request, data interface{}, tableName string) {
	c := appengine.NewContext(r)
	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, tableName, nil), data)
	if err != nil {
		print(key)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func memcache_INSERT(w http.ResponseWriter, r *http.Request, data interface{}, tableName string) {
	c := appengine.NewContext(r)
	this_item, err := memcache.Get(c, "myKey")
	var myKey = "this_key"
	if err == nil {
		myKey = string(this_item.Value)
	}
	log.Fatal(myKey)
	memcache.Set(c, &memcache.Item{
		Key:   "myKey",
		Value: []byte("Go"),
	})
}

func apiIndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.(http.ResponseWriter).Write([]byte("{\"details\":\"access denied\"}"))
}

func apiHandler(w rest.ResponseWriter, r *rest.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.(http.ResponseWriter).Write([]byte("{\"details\":\"access denied\"}"))
}
