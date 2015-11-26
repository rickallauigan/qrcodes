package models

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type Value struct {
	Id                int64  `json:"id" datastore:"-"`
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

func (t *Value) Save(c context.Context) (*Value, error) {
	k, err := datastore.Put(c, t.key(c), t)
	if err != nil {
		return nil, err
	}
	t.Id = k.IntID()
	return t, nil
}

func (this Value) Add(c context.Context) error {
	key := datastore.NewIncompleteKey(c, "tblData", nil)
	_, err := datastore.Put(c, key, &this)
	return err
}

func ListData(c context.Context) *datastore.Key {
	return datastore.NewKey(c, "List", "data", 0, nil)
}

func (v *Value) key(c context.Context) *datastore.Key {
	return datastore.NewKey(c, "Value", "", 0, ListData(c))
}
