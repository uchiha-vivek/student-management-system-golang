package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

type Student struct {
	// id , name , email , phone_no
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Phone  int     `json:"phone"`
	Branch *Branch `json:"branch"`
}

type Branch struct {
	Name string `json:"name"`
}

var students []Student

func main() {
	r := mux.NewUser()
	r.HandleFunc("/students", getStudents).Methods("GET")
	r.HandleFunc("/students/{id}", getStudentByID).Methods("GET")
	r.HandleFunc("/students", createStudent).Methods("POST")
	r.HandleFunc("/students/{id}", updateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")
	fmt.Printf("Server is running on 8000 \n")

}
