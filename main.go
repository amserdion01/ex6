package main

import (
	"encoding/json"
	entity "ex6/Entity"
	"ex6/db"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-chi/chi"
)

func main() {
	db.InitDatabase()
	router := chi.NewRouter()
	router.Route("/api", func(r chi.Router) {
		r.Post("/location", PostLocation)
		r.Get("/location", GetLocation)
		r.Post("/employee", PostEmployee)
		r.Post("/company", PostCompany)
		r.Get("/company", GetCompany)

	})
	http.ListenAndServe(":8080", router)
}

func PostLocation(w http.ResponseWriter, r *http.Request) {
	location := entity.Location{}
	body := r.Body
	decodedBody, _ := io.ReadAll(body)
	json.Unmarshal(decodedBody, &location)
	value := fmt.Sprintf("My location is %v, %v", location.Latitude, location.Longitude)
	db.GetDB().Create(&location)
	fmt.Fprintf(w, value)

}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	fmt.Fprint(w, city)
	var location = entity.Location{}
	db.GetDB().Where("city=?", city).Find(&location)
	fmt.Fprintln(w, location)
}

func PostEmployee(w http.ResponseWriter, r *http.Request) {
	employee := entity.Employee{}
	body := r.Body
	decodedBody, _ := io.ReadAll(body)
	json.Unmarshal(decodedBody, &employee)
	db.GetDB().Create(&employee)
	fmt.Fprintf(w, employee.Name)
}

func PostCompany(w http.ResponseWriter, r *http.Request) {
	company := entity.Company{}
	body := r.Body
	decodedBody, _ := io.ReadAll(body)
	json.Unmarshal(decodedBody, &company)
	db.GetDB().Create(&company)
	fmt.Print(company)
	fmt.Fprintf(w, company.Name)
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, name)
	var company = entity.Company{}
	db.GetDB().Where("city=?", name).Find(&company)
	fmt.Fprintln(w, company)
}
