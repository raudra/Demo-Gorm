package controllers


import (
		"net/http"
    	//"encoding/json"
    	//"log"
    	//"io/ioutil"
		"fmt"
	)

type employeeController struct{}

var EmployeeController employeeController

func init(){
	EmployeeController = employeeController{}
}

func (e *employeeController)Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome!....")
}