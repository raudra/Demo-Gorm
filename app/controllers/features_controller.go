package controllers


import (
		"net/http"
        "encoding/json"
		"demo_gorm/app/models"
		// "log"
		"io/ioutil"
	)

type featureController struct{
	params map[string]interface{}	
}

var FeatureController featureController



func init(){
	FeatureController = featureController{}
}

func (e *featureController)Index(w http.ResponseWriter, r *http.Request){
	data, err:=json.Marshal(models.FeatureModel.All())
	if err != nil{
		panic(err)
	}
	w.Write(data)
}

func (self *featureController)Create(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &self.params)
	data, err := json.Marshal(models.FeatureModel.Create(self.params))
	if err != nil{
		panic(err)
	}
	w.Write(data)
}