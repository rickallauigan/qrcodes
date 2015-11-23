package controllers

import (
	"github.com/astaxie/beegae"
	//"google.golang.org/appengine/datastore"
)

type MainController struct {
	beegae.Controller
}
type Values struct {
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

func (this *MainController) Get() {

	this.Data["Profile"] = &Values{
		TimeCreated:       "oras",
		Name:              "Marcelo",
		Age:               50,
		Municipality:      "Balanga",
		Family:            "1 Son",
		FarmLocation:      "Balanga City",
		VarietyOfRice:     "Fancy",
		AppliedFertilizer: "Organic",
		LandPreparation:   "05/01/2015",
		Seedling:          "07/01/2015",
		Harvesting:        "10/01/2015",
	}
	this.TplNames = "main.html"
}
