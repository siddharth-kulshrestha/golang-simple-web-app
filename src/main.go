package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		loginTemplate.ExecuteTemplate(response, "base.gohtml", nil)
	case "POST":
		err := request.ParseForm()
		if err != nil {
			fmt.Println(err)
		}
		user := UserDetail{Name: "",
			Username: request.PostFormValue("username"),
			Password: request.PostFormValue("password")}
		data := make(map[string]string)
		if ValidatePassword(user) {
			data["message"] = "User has been logged in successfully."
			data["details"] = "Click here to see details of all users."
			data["detail_route"] = "/details"
			loginTemplate.ExecuteTemplate(response, "base.gohtml", data)
		} else {
			data["message"] = "Incorrect username or password."
			loginTemplate.ExecuteTemplate(response, "base.gohtml", data)
		}
	default:
		fmt.Println("error request method is not supported.")
	}
}

func RegisterHandler(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		registerTemplate.ExecuteTemplate(response, "base.gohtml", nil)
	case "POST":
		err := request.ParseForm()
		if err != nil {
			fmt.Println(err)
		}
		user := UserDetail{Name: request.PostFormValue("name"),
			Username: request.PostFormValue("username"),
			Password: request.PostFormValue("password")}
		data := make(map[string]string)
		if CommitUser(user) {
			data["message"] = "User has been created successfully."
		} else {
			data["message"] = "User has not been created. Please try again later"
		}
		registerTemplate.ExecuteTemplate(response, "base.gohtml", data)

	default:
		fmt.Println("error request method is not supported.")
	}

}

func DetailsHandler(response http.ResponseWriter, request *http.Request) {
	data := GetAllUsers()
	detailsTemplate.ExecuteTemplate(response, "base.gohtml", data)
}

var detailsTemplate, loginTemplate, registerTemplate *template.Template

func init() {
	loginTemplate = template.Must(template.ParseFiles("../templates/components/login.gohtml", "../templates/containers/base.gohtml"))
	registerTemplate = template.Must(template.ParseFiles("../templates/components/register.gohtml", "../templates/containers/base.gohtml"))
	detailsTemplate = template.Must(template.ParseFiles("../templates/components/details.gohtml", "../templates/containers/base.gohtml"))
}

func main() {
	fmt.Println("server started")
	http.HandleFunc("/", LoginHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/details", DetailsHandler)
	http.ListenAndServe(":8080", nil) //using default mux

}
