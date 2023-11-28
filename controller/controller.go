package controller

import (
	"encoding/json"
	"fmt"
	"go-web-echo-app/config"
	"go-web-echo-app/models"
	"html/template"
	"io"
	"log"
	"net/http"

	"go-web-echo-app/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request: ", r)
	switch method := r.Method; {
	case method == "GET":
		fmt.Println("Request: ", r)
		// msg := fmt.Sprintf("Bro, ur so close....")
		t, err := template.ParseFiles("assets/templates/home.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		t.Execute(w, map[string]interface{}{
			// "MessageToUser": msg,
		})
		return
	case method == "POST":
		byteSlice, err := io.ReadAll(r.Body)
		//		fmt.Println("bytes: ", byteSlice)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to read the request body.")
			http.Error(w, errMsg, http.StatusInternalServerError)
		}
		defer r.Body.Close()
		w.Header().Add("x-go-web-echo-app", "custom header")
		w.WriteHeader(200)
		w.Write(byteSlice)
		return
	default:
		errMsg := fmt.Sprintf("Invalid HTTP Method: %s", method)
		http.Error(w, errMsg, http.StatusInternalServerError)

	}

}

func PageTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request: ", r)
	t, err := template.ParseFiles("assets/templates/page2.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
	return
}

func PageThree(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request: ", r)
	t, err := template.ParseFiles("assets/templates/page3.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
	return
}

type User struct {
	ID       uint   `gorm:"primaryKey; unique; autoIncrement; not null" json:"id"`
	Email    string `gorm:"column:email; unique; not null" json:"email"`
	Username string `gorm:"column:username; unique; not null" json:"username"`
	Password string `gorm:"column:password; not null" json:"password"`
}

type RegisterUserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=4"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		res := utils.BuildErrorResponse("Failed to register user. Endpoint accepts only POST requests.", "", utils.EmptyObj{})
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}

	fmt.Println("r.Body: ", r.Body)

	var req RegisterUserDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := utils.BuildErrorResponse("Failed to decode body", err.Error(), utils.EmptyObj{})
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}

	// Save user to database
	user := models.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}
	if err := config.DB.Create(&user).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := utils.BuildErrorResponse("Failed to register user", err.Error(), utils.EmptyObj{})
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}

	res := utils.BuildResponse(true, "OK!", "User registered successfully")
	fmt.Println("res", res)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := utils.BuildErrorResponse("Failed to register user", err.Error(), utils.EmptyObj{})
		json.NewEncoder(w).Encode(res)
		return
	}
	return
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusBadRequest)
		res := utils.BuildErrorResponse("Failed to register user. Endpoint accepts only POST requests.", "", utils.EmptyObj{})
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		res := utils.BuildErrorResponse("Failed to register user. Endpoint accepts only GET requests.", "", utils.EmptyObj{})
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}

	params := r.URL.Query()
	var queryUsername string
	if _, ok := params["username"]; ok != true {
		res := utils.BuildErrorResponse("Failed to get User. Malformed query string.", "", utils.EmptyObj{})
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	} else {
		queryUsername = params["username"][0]
	}

	var user models.User
	if err := config.DB.Where("username = ?", queryUsername).First(&user).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := utils.BuildErrorResponse("Failed to get user", err.Error(), utils.EmptyObj{})
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}

	res := utils.BuildResponse(true, "OK!", user)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := utils.BuildErrorResponse("Failed to register user", err.Error(), utils.EmptyObj{})
		json.NewEncoder(w).Encode(res)
		return
	}
	return
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusBadRequest)
		res := utils.BuildErrorResponse("Failed to register user. Endpoint accepts only POST requests.", "", utils.EmptyObj{})
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Printf("Error sending response: %v", err)
		}
		return
	}
}
