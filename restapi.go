package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type emp struct {
	Eno        string  `json:"eno"`
	Ename      string `json:"ename"`
	Age        string `json:"age"`
	Department string `json:"department"`
}

var Emps []emp

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func register(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		data := emp{}
		json.Unmarshal([]byte(string(body)), &data)

		fmt.Println("Endpoint Hit: register ", data.Eno, data.Ename, data.Age, data.Department)

	Emps = append(Emps, data)
	json.NewEncoder(w).Encode(data)

		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Println("Endpoint Hit: register ", r.Method)
}

func getEmp(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/getInfo/")
	if id != "" {
    	for _, emp := range Emps {
		if id == emp.Eno {
			json.NewEncoder(w).Encode(emp)
		}
		
	}
	return
}
    	json.NewEncoder(w).Encode(Emps)
}

func updateEmp(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		data := emp{}
		json.Unmarshal([]byte(string(body)), &data)
    	id := strings.TrimPrefix(r.URL.Path, "/updateInfo/")
	    if id != "" {
        
		}



		fmt.Println("Endpoint Hit: update ", data.Eno, data.Ename, data.Age, data.Department)

	for i, looper := range Emps {
		if looper.Eno == id {
			looper.Eno = data.Eno
			looper.Ename = data.Ename
			looper.Age = data.Age
			looper.Department = data.Department

			Emps = append(Emps[:i], looper)
			json.NewEncoder(w).Encode(looper)
		}
	}	}
}

func deleteEmp(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {
  
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		data := emp{}
		json.Unmarshal([]byte(string(body)), &data)
		fmt.Println("Endpoint Hit: delete ", data.Eno)
        id := strings.TrimPrefix(r.URL.Path, "/delInfo/")
	    if id != "" {
	        data.Eno =id
        }

	for index, emp := range Emps {
		if emp.Eno == id {
			Emps = append(Emps[:index], Emps[index+1:]...)
		}
	}
	}
}

func handleRequests() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/createInfo", register)
	http.HandleFunc("/updateInfo/", updateEmp)
	http.HandleFunc("/delInfo/", deleteEmp)
	http.HandleFunc("/getInfo/", getEmp)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Sever Start at 8000")
	handleRequests()
}
