package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
    "strconv"
	_ "github.com/go-sql-driver/mysql"
)

type emp struct {
	Eno        int    `json:"eno"`
	Ename      string `json:"ename"`
	Age        string `json:"age"`
	Department string `json:"department"`
}

func dbCon() (db *sql.DB) {
	db, error := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/employee")
	if error != nil {
		panic(error)
	}
	return db
}

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

		db := dbCon()

		var query = "INSERT INTO emp (eno,ename,age,department) VALUES (?,?,?,?)"

		insert, err := db.Query(query, data.Eno, data.Ename, data.Age, data.Department)

		defer insert.Close()

		t := data
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(t); err != nil {
			panic(err)
		}
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Println("Endpoint Hit: register ", r.Method)
}

func getEmp(w http.ResponseWriter, r *http.Request) {

	db := dbCon()
	query := "SELECT * FROM emp"
	id := strings.TrimPrefix(r.URL.Path, "/getInfo/")
	if id != "" {
    query=query+" where eno = "+id
	fmt.Println("query", query)
	}
    
	empRes, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	res := []emp{}

	for empRes.Next() {
		data := emp{}
		err = empRes.Scan(&data.Eno, &data.Ename, &data.Age, &data.Department)
		if err != nil {
			panic(err)
		}
		res = append(res, data)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
	return
	fmt.Println("Endpoint Hit: register ", r.Method)

}

func updateEmp(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		data := emp{}
		json.Unmarshal([]byte(string(body)), &data)

		fmt.Println("Endpoint Hit: update ", data.Eno, data.Ename, data.Age, data.Department)

		db := dbCon()
		var query = "UPDATE emp SET ename = ? ,age = ? ,department = ? WHERE eno = ?"
		empRes, err := db.Query(query, data.Ename, data.Age, data.Department, data.Eno)
		if err != nil {
			panic(err)
		}
		defer empRes.Close()
		t := data
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//	w.WriteHeader(http.OK)
		if err := json.NewEncoder(w).Encode(t); err != nil {
			panic(err)
		}
		return
	}
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
	        data.Eno,err =strconv.Atoi(id)
        }

		db := dbCon()
		var query = "DELETE FROM emp WHERE eno = ? "
		
		empRes, err := db.Query(query, data.Eno)
		if err != nil {
			panic(err)
		}
		defer empRes.Close()
		t := data
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//	w.WriteHeader(http.OK)
		if err := json.NewEncoder(w).Encode(t); err != nil {
			panic(err)
		}
		return
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
