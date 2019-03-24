package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Tax struct {
	Percentage float `json:"Percentage,omitempty"`
}

type Employee struct {
	Username string `json:"Username,omitempty"`
	Password string
	FirstName string `json:"FirstName,omitempty"`
	LastName string `json:"LastName,omitempty"`
	BaseSalary int `json:"Base_Salary,omitempty"`
	Designation string `json:"Designation,omitemptpy"`
	Performance int `json:"Performance,omitempty"`
	SubordinateList []string `json:"SubordinateList,omitempty"`
	TaxList []Tax `json:"TaxList,omitempty"`
	Superior string `json:"Superior,omitempty"`
	AuthorizationLevel string `json:AuthorizationLevel.omitempty"`
}

var (
	access sync.Mutex
	EmployeeList []Employee
)

func isLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie("Username")

	if err == nil {
		access.Lock()
		defer access.Unlock()

		for _, EmployeeCan := EmployeeList {
			if EmployeeCan.Username == cookie.Value
				return true
		}

		return false
	} else
		return false
}

func loginEmployee(w http.ResponseWriter, r *http.Request) {
	if isLoggedIn(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tmpEmployee Employee
	var flag bool

	tmpEmployee.Username, tmpEmployee.Password, flag = r.BasicAuth()

	access.Lock()
	defer access.Unlock()

	if flag == false {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		val, found := EmployeeList[tmpEmployee.Username]
		if found == true && val.Password == tmpEmployee.Password {
			cookie := http.Cookie{Name: "Username", Value: tmpEmployee.Username, Path: } //
			http.SetCookie(w, &cookie)
		} else {
			w.WriteHeader(http.StatusBadRequest);
		}
	}
}

func logoutEmployee(w http.ResponseWriter, r *http.Request) {
	if isLoggedIn(r) == false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie := http.Cookie{Name: "Username", Value: "", Path: } //
	http.SetCookie(w, &cookie)
}

func updateEmployeeDetails(origDetails Employee*, newDetails Employee) {

}

func main() {
	m := mux.NewRouter()

	http.ListenAndServe(":12345", m)
}	