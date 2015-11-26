package controllers

import (
	//"fmt"
	"encoding/json"
	"github.com/astaxie/beegae"
	"io"
	//"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"techaguru-orangepineapple/models"
)

type MainController struct {
	beegae.Controller
}

type Value struct {
	TimeCreated       string `json:"time created"`
	Name              string `json:"name"`
	Age               int    `json:"age"`
	Municipality      string `json:"municipality"`
	Family            string `json:"family"`
	FarmLocation      string `json:"farm location"`
	VarietyOfRice     string `json:"variety of rice"`
	AppliedFertilizer string `json:"applied fertilizer"`
	LandPreparation   string `json:"land preparation"`
	Seedling          string `json:"seedling"`
	Planting          string `json:"planting"`
	Harvesting        string `json:"harvesting"`
}

func (this *MainController) Post() {
	profile, err := decodeProfile(this.Ctx.Input.Request.Body)
	if err != nil {
		this.Data["json"] = err
		return
	}
	t, err := profile.Save(this.AppEngineCtx)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = &t
	}
}

func decodeProfile(r io.ReadCloser) (*models.Value, error) {
	defer r.Close()
	var profile models.Value
	err := json.NewDecoder(r).Decode(&profile)
	return &profile, err
}

func (this *MainController) Get() {
	profile := []models.Value{}
	ks, err := datastore.NewQuery("Value").GetAll(this.AppEngineCtx, &profile)
	if err != nil {
		this.Data["Profile"] = err
		return
	}
	for i := 0; i < len(profile); i++ {
		profile[i].Id = ks[i].IntID()
	}
	this.Data["Profile"] = &profile
	//this.ServeJson()
	this.TplNames = "main.html"
}

// func (this *MainController) Get() {
// 	//var values []Value
// 	values := []models.Value{}
// 	res, err := datastore.NewQuery("tblData").Ancestor(models.ListData(this.AppEngineCtx)).GetAll(this.AppEngineCtx, &values)
// 	//res, err := datastore.NewQuery("tblData").GetAll(this.AppEngineCtx, &values)

// 	if err != nil {
// 		// this.Data["Profile"] = err
// 		return
// 	}

// 	// for i := 0; i < len(values); i++ {
// 	// 	values[i].Name = "rick"
// 	// }

// 	// 	this.Data["Profile"] = &Value{
// 	// 		TimeCreated:       "oras",
// 	// 		Name:              "Marcelo",
// 	// 		Age:               50,
// 	// 		Municipality:      "Balanga",
// 	// 		Family:            "1 Son",
// 	// 		FarmLocation:      "Balanga City",
// 	// 		VarietyOfRice:     "Fancy",
// 	// 		AppliedFertilizer: "Organic",
// 	// 		LandPreparation:   "05/01/2015",
// 	// 		Seedling:          "07/01/2015",
// 	// 		Harvesting:        "10/01/2015",
// 	// 	}

// 	// this.Data["Profile"] = &res

// 	this.Data["json"] = &res
// 	this.ServeJson()
// }
