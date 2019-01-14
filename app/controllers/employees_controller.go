package controllers


import (
		"fmt"
		"net/http"
        "encoding/json"
		"demo_gorm/app/models"
	)

type employeeController struct{
	params map[string]interface{}
}

var EmployeeController employeeController

func init(){
	EmployeeController = employeeController{}
}

func (e *employeeController)Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome!....")
}


func(e *employeeController)Fetch(w http.ResponseWriter, r *http.Request){
	// body, err := ioutil.ReadAll(r.Body)
	// err = json.Unmarshal(body, &self.params)
	data, err := json.Marshal(models.UserModel.FetchByEmail("pratap.raudra@gmail.com"))
	if err != nil{
		panic(err)
	}
	w.Write(data)
}
