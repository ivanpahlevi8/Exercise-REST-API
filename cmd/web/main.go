package main

import (
	"exercise-web/pkg/handler"
	"exercise-web/pkg/repository"
	"exercise-web/pkg/service"
	"net/http"
)

func main() {
	// Initializing routing for user
	myRepo := repository.InitRepo()

	myService := service.InitService(myRepo)

	myHandler := handler.InitHandler(myService)

	http.HandleFunc("/", myHandler.AddData)

	http.HandleFunc("/getuserid", myHandler.GetDataById)

	http.HandleFunc("/getuserusername", myHandler.GetDataByUsername)

	http.HandleFunc("/getuserpassword", myHandler.GetDataByPassword)

	http.HandleFunc("/updateuserid", myHandler.UpdateUserById)

	http.HandleFunc("/updateuserusername", myHandler.UpdateUserByUsername)

	http.HandleFunc("/updateuserpassword", myHandler.UpdateUserByPassword)

	http.HandleFunc("/deleteuserid", myHandler.DeleteUserById)

	http.HandleFunc("/deleteuserusername", myHandler.DeleteUserByUsername)

	http.HandleFunc("/deleteuserpassword", myHandler.DeleteUserByPassword)

	http.ListenAndServe(":6060", nil)
}
