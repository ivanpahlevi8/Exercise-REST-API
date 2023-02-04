package handler

import (
	"encoding/json"
	"exercise-web/pkg/model"
	"exercise-web/pkg/service"
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
	This package where response to handler request happen
*/

var MyHandler Handle

type Handle struct {
	serve service.Service
}

func InitHandler(serv service.Service) Handle {
	MyHandler.serve = serv
	return MyHandler
}

func (h *Handle) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userData model.UserData

	_ = json.NewDecoder(r.Body).Decode(&userData)

	fmt.Printf("Get Account From Handler, \nId User : %s \nUsername : %s \nPassword : %s\n", userData.GetId(), userData.GetUsername(), userData.GetPassword())

	var user model.User
	user.SetId(userData.GetId())
	user.SetUsername(userData.GetUsername())
	user.SetPassword(userData.GetPassword())

	data := h.serve.AddData(user)

	json.NewEncoder(w).Encode(data)
	log.Println("Success Adding Data")
	w.WriteHeader(201)
}

func (h *Handle) GetDataById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var getUser model.User
	var err error

	params := r.URL.Query()

	getId := params.Get("id")

	getUser, err = h.serve.GetDataById(getId)

	if err != nil {
		str := fmt.Sprintf("Error Happen When Get Data Using Id")
		log.Println(str)
		// error json respones error
		var errorUser model.ErrorUser
		t := time.Now().Format("2006-01-02 15:04:05")
		errorUser.SetErrMessage(str)
		errorUser.SetTimeStamp(t)

		json.NewEncoder(w).Encode(errorUser)
	} else {
		// make json user
		var user model.UserData
		user.SetId(getUser.GetId())
		user.SetUsername(getUser.GetUsername())
		user.SetPassword(getUser.GetPassword())

		json.NewEncoder(w).Encode(user)
	}
}

func (h *Handle) GetDataByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var getUser model.User
	var err error

	params := r.URL.Query()

	getUsername := params.Get("username")

	getUser, err = h.serve.GetDataByUsername(getUsername)

	if err != nil {
		str := fmt.Sprintf("There Is an error when get user by username : %s\n", err)
		log.Println(str)
		// what happen if error
		var userError model.ErrorUser
		t := time.Now().Format("2006-01-02 15:04:05")
		userError.SetErrMessage(str)
		userError.SetTimeStamp(t)

		json.NewEncoder(w).Encode(userError)
	} else {
		// make user data to response as json
		var user model.UserData
		user.SetId(getUser.GetId())
		user.SetUsername(getUser.GetUsername())
		user.SetPassword(getUser.GetPassword())

		json.NewEncoder(w).Encode(user)
	}
}

func (h *Handle) GetDataByPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// variable to get
	var getUser model.User
	var err error

	// input parameter from url
	params := r.URL.Query()

	getPassword := params.Get("password")

	// get user based on password
	getUser, err = h.serve.GetDataByPassword(getPassword)

	if err != nil {
		// if there is an error happern
		str := fmt.Sprintf("There is an error happen when Get User By Password : %s\n", err)
		log.Println(str)
		// make error response as json to body
		var errorUser model.ErrorUser
		t := time.Now().Format("2006-01-02 15:04:05")
		errorUser.SetErrMessage(str)
		errorUser.SetTimeStamp(t)

		json.NewEncoder(w).Encode(errorUser)
	} else {
		// if there is no error happen
		var user model.UserData
		user.SetId(getUser.GetId())
		user.SetUsername(getUser.GetUsername())
		user.SetPassword(getUser.GetPassword())

		json.NewEncoder(w).Encode(user)
	}
}

func (h *Handle) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// make variable to get from body response
	var getUserFromBody model.UserData
	var err error

	// get Data from body
	json.NewDecoder(r.Body).Decode(&getUserFromBody)

	// get id from url
	params := r.URL.Query()

	getId := params.Get("id")

	// update user based on id, make user to fit with database
	var userDatabase model.User
	userDatabase.SetId(getUserFromBody.GetId())
	userDatabase.SetUsername(getUserFromBody.GetUsername())
	userDatabase.SetPassword(getUserFromBody.GetPassword())

	getUser, err := h.serve.UpdateUserById(getId, userDatabase)
	log.Println(err)

	if err != nil {
		// what happen is there is an error
		str := fmt.Sprintf("Getting Error When Update Data Using Id : %s\n", err)
		log.Println(str)
		// Make error response to user
		var errorUser model.ErrorUser
		t := time.Now().Format("2006-01-02 15:04:05")
		errorUser.SetErrMessage(str)
		errorUser.SetTimeStamp(t)

		json.NewEncoder(w).Encode(errorUser)
	} else {
		// update data user variable from body input
		getUserFromBody.SetId(getUser.GetId())
		getUserFromBody.SetUsername(getUser.GetUsername())
		getUserFromBody.SetPassword(getUser.GetPassword())

		json.NewEncoder(w).Encode(getUserFromBody)
	}
}

func (h *Handle) UpdateUserByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// make variable to get data from body response
	var getDataFromBody model.UserData
	var err error

	json.NewDecoder(r.Body).Decode(&getDataFromBody)

	// get variable username from url
	params := r.URL.Query()

	getUsername := params.Get("username")

	// update data using username in database, but make user databse to fot with database column
	var getDatabase model.User
	getDatabase.SetId(getDataFromBody.GetId())
	getDatabase.SetUsername(getDataFromBody.GetUsername())
	getDatabase.SetPassword(getDataFromBody.GetPassword())

	getUser, err := h.serve.UpdateUserByUsername(getUsername, getDatabase)

	// check if there is an error
	if err != nil {
		// there is an error
		str := fmt.Sprintf("Error Happen When Update Data Using Username : %s\n", err)
		log.Println(str)
		// make error response
		var userError model.ErrorUser
		t := time.Now().Format("2006-01-02 15:04:05")
		userError.SetErrMessage(str)
		userError.SetTimeStamp(t)

		json.NewEncoder(w).Encode(userError)
	} else {
		// uddate user data from body to return as json to user
		getDataFromBody.SetId(getUser.GetId())
		getDataFromBody.SetUsername(getUser.GetUsername())
		getDataFromBody.SetPassword(getUser.GetPassword())

		json.NewEncoder(w).Encode(getDataFromBody)
	}
}

func (h *Handle) UpdateUserByPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// create variable
	var getUserFromBodyRequest model.UserData
	var err error

	// get user from body
	json.NewDecoder(r.Body).Decode(&getUserFromBodyRequest)

	// get variable from URL
	params := r.URL.Query()

	getPassword := params.Get("password")

	// update user from database, make user to fit with database
	var userDatabase model.User
	userDatabase.SetId(getUserFromBodyRequest.GetId())
	userDatabase.SetUsername(getUserFromBodyRequest.GetUsername())
	userDatabase.SetPassword(getUserFromBodyRequest.GetPassword())

	getUser, err := h.serve.UpdateUserByPassword(getPassword, userDatabase)

	// error check
	if err != nil {
		// if there is an error
		str := fmt.Sprintf("Error Happen When Update Data Using Password : %s\n", err)
		log.Println(str)
		var userError model.ErrorUser
		userError.SetErrMessage(str)
		t := time.Now().Format("2006-01-02 15:04:05")
		userError.SetTimeStamp(t)

		json.NewEncoder(w).Encode(userError)
	} else {
		// update data user for body response
		getUserFromBodyRequest.SetId(getUser.GetId())
		getUserFromBodyRequest.SetUsername(getUser.GetUsername())
		getUserFromBodyRequest.SetPassword(getUser.GetPassword())

		json.NewEncoder(w).Encode(getUserFromBodyRequest)
	}
}

func (h *Handle) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// make variable to get from database
	var getUserDatabase model.User
	var err error

	// get variable id from url
	params := r.URL.Query()

	getId := params.Get("id")

	// delete user based on id
	getUserDatabase, err = h.serve.DeleteUserById(getId)

	if err != nil {
		// if error happen
		str := fmt.Sprintf("Error Happen When Delete User Using Id : %s\n", err)
		log.Println(str)
		// make error response
		var errorUser model.ErrorUser
		t := time.Now().Format("2006-01-02 15:04:05")
		errorUser.SetErrMessage(str)
		errorUser.SetTimeStamp(t)

		json.NewEncoder(w).Encode(errorUser)
	} else {
		// make user data for json response in body response
		var userResponse model.UserData
		userResponse.SetId(getUserDatabase.GetId())
		userResponse.SetUsername(getUserDatabase.GetUsername())
		userResponse.SetPassword(getUserDatabase.GetPassword())

		json.NewEncoder(w).Encode(userResponse)
	}
}

func (h *Handle) DeleteUserByUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// Create variable for get data from database
	var getUserDatabase model.User
	var err error

	// get variable username from url
	params := r.URL.Query()

	getUsername := params.Get("username")

	// get user from database
	getUserDatabase, err = h.serve.DeleteUserByUsername(getUsername)

	// check error
	if err != nil {
		// if error happen
		str := fmt.Sprintf("Error Happen When Deleted User By Username : %s\n", err)
		log.Println(str)
		// make error response
		var errorUser model.ErrorUser
		t := time.Now().Format("2006-01-02 15:04:05")
		errorUser.SetErrMessage(str)
		errorUser.SetTimeStamp(t)

		json.NewEncoder(w).Encode(errorUser)
	} else {
		// make user for body request
		var userBody model.UserData
		userBody.SetId(getUserDatabase.GetId())
		userBody.SetUsername(getUserDatabase.GetUsername())
		userBody.SetPassword(getUserDatabase.GetPassword())

		json.NewEncoder(w).Encode(userBody)
	}
}

func (h *Handle) DeleteUserByPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// make variable
	var getUserDatabase model.User
	var err error

	// get password username from url
	params := r.URL.Query()

	getPassword := params.Get("password")

	// delete user base on password
	getUserDatabase, err = h.serve.DeleteUserByPassword(getPassword)

	// error check
	if err != nil {
		// thre is an error
		str := fmt.Sprintf("There is an error when deleted data by password : %s\n", err)
		log.Println(str)
		// make error response
		var errorUser model.ErrorUser
		t := time.Now().Format("2006-01-02 15:04:05")
		errorUser.SetErrMessage(str)
		errorUser.SetTimeStamp(t)

		json.NewEncoder(w).Encode(errorUser)
	} else {
		// create data user for body response
		var userBody model.UserData
		userBody.SetId(getUserDatabase.GetId())
		userBody.SetUsername(getUserDatabase.GetUsername())
		userBody.SetPassword(getUserDatabase.GetPassword())

		json.NewEncoder(w).Encode(userBody)
	}
}
