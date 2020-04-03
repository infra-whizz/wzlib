package wzlib_traits

import (
	"io/ioutil"

	wzlib_logger "github.com/infra-whizz/wzlib/logger"
	"gopkg.in/yaml.v2"
)

type WzTraitsContainer struct {
	content map[string]interface{}
	wzlib_logger.WzLogger
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
		wt.GetLogger().Errorln("Unable to read existing traits from file:", err.Error())
		return wt
	}

	err = yaml.Unmarshal(buff, wt.content)
	if err != nil {
		wt.GetLogger().Fatalln(err)
	}
	return wt
}

// Save traits to the file
func (wt *WzTraitsContainer) SaveToFile(fpath string) {
	data, err := yaml.Marshal(wt.content)
	if err != nil {
		wt.GetLogger().Fatalln(err)
	}

	if err := ioutil.WriteFile(fpath, data, 0640); err != nil {
		wt.GetLogger().Fatalln(err)
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
