package main

import (
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/astaxie/beegae"
	"log"
	"net/http"
	"sync"
	"techaguru-orangepineapple/controllers"
	_ "techaguru-orangepineapple/routers"
	"time"

)

type values struct {
	TimeCreated         time.Time `json:"time created"`
	Name                string    `json:"name"`
	Age                 string    `json:"age"`
	Municipality        string    `json:"municipality"`
	Family              string    `json:"family"`
	FarmLocation        string    `json:"farm location"`
	VarietyOfRicestring string    `json:"variety of rice"`
	AppliedFertilizer   string    `json:"applied fertilizer"`
	LandPreparation     string    `json:"land preparation"`
	Seedling            string    `json:"seedling"`
	Planting            string    `json:"planting"`
	Harvesting          string    `json:"harvesting"`
}

var lock = sync.RWMutex{}

func init() {
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("../static/css"))))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", http.FileServer(http.Dir("../static/js"))))
	http.Handle("/static/components/js/", http.StripPrefix("/static/components/js/", http.FileServer(http.Dir("../static/components/js"))))

	beegae.Router("/", &controllers.MainController{})
	beegae.Run()

	//api := rest.NewApi()
	//api.Use(rest.DefaultDevStack...)
	//router, err := rest.MakeRouter(
	//rest.Post("/api/v1/", PostDataHandler),
	//)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// api.SetApp(router)
	// http.Handle("/", api.MakeHandler())
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

}

func PostDataHandler(w rest.ResponseWriter, r *rest.Request) {
	vals := values{
		TimeCreated: time.Now(),
	}
	err := r.DecodeJsonPayload(&vals)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	lock.Lock()
	datastore_INSERT(w.(http.ResponseWriter), r.Request, &vals, "tblData")
	lock.Unlock()
	w.WriteJson(&vals)
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
