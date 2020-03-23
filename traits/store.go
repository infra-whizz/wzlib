package wzlib_traits

import (
	"log"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type WzTraitsContainer struct {
	content map[string]interface{}
}

// Constructor
func NewWzTraitsContainer() *WzTraitsContainer {
	wt := new(WzTraitsContainer)
	wt.content = make(map[string]interface{})
	return wt
}

// LoadFromFile traits
func (wt *WzTraitsContainer) LoadFromFile(fpath string) *WzTraitsContainer {
	buff, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Println("Unable to read from file:", err.Error())
	}
	err = yaml.Unmarshal(buff, wt.content)
	if err != nil {
		panic(err)
	}
	return wt
}

// Save traits to the file
func (wt *WzTraitsContainer) SaveToFile(fpath string) {
	data, err := yaml.Marshal(wt.content)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(fpath, data, 0640); err != nil {
		panic(err)
	}
}

// Get a trait
func (wt *WzTraitsContainer) Get(trait string) interface{} {
	val, ex := wt.content[trait]
	if !ex {
		return nil
	} else {
		return val
	}
}

// Set a trait to a permanent storage
func (wt *WzTraitsContainer) Set(trait string, value interface{}) {
	wt.content[trait] = value
}
