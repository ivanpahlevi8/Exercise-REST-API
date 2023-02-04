package service

import (
	"errors"
	"exercise-web/pkg/model"
	"exercise-web/pkg/repository"
	"fmt"
	"log"
)

/*
	This package where all logiuc business happen. All method to do REST API method are written in here
	This pacakage uising repository as database
*/

var MyService Service

type Service struct {
	repo repository.Repo
}

func InitService(repo repository.Repo) Service {
	MyService.repo = repo
	return MyService
}

func (s *Service) AddData(data model.User) model.User {
	s.repo.AddData(data)
	return data
}

func (s *Service) GetDataById(id string) (model.User, error) {
	var getUser model.User
	var err error

	getUser, err = s.repo.GetData(id)

	if err != nil {
		// return failed validation in the next time
		log.Println(err)
		return getUser, err
	}

	return getUser, nil
}

func (s *Service) GetDataByUsername(username string) (model.User, error) {
	var allUser []model.User
	var getUser model.User
	var err error
	check := false

	allUser, err = s.repo.GetAllData()

	if err != nil {
		// return failed validation in next update
		log.Println(err)
		return getUser, err
	}

	for _, user := range allUser {
		fmt.Println("username input : ", username)
		fmt.Println("username from iterate : ", user.GetUsername())
		if username == user.GetUsername() {
			getUser = user
			check = true
			break
		}
	}

	if check {
		// jika data ditemukan berdasarkan usernam
		return getUser, nil
	} else {
		// jika tidak ditemukan data
		var trans model.User
		errs := errors.New("error happen when getting user by username")
		return trans, errs
	}

}

func (s *Service) GetDataByPassword(password string) (model.User, error) {
	var allUser []model.User
	var getUser model.User
	var err error
	check := false

	allUser, err = s.repo.GetAllData()

	if err != nil {
		// return failed validation in next update
		log.Println(err)
		return getUser, err
	}

	for _, user := range allUser {
		if password == user.GetPassword() {
			getUser = user
			check = true
			break
		}
	}

	if check {
		// jika data ditemukan berdasarkan usernam
		return getUser, nil
	} else {
		// jika tidak ditemukan data
		var trans model.User
		errs := errors.New("error happen when getting user using password")
		return trans, errs
	}

}

func (s *Service) UpdateUserById(id string, user model.User) (model.User, error) {
	var getUser model.User
	var err error

	getUser, err = s.repo.UpdateData(id, user)

	if err != nil {
		log.Println(err)
		return getUser, err
	}

	return getUser, nil
}

func (s *Service) UpdateUserByUsername(username string, user model.User) (model.User, error) {
	var allUser []model.User
	var getUser model.User
	var getIdFromUser string
	var err error

	allUser, err = s.repo.GetAllData()

	if err != nil {
		log.Println(err)
		return getUser, nil
	}

	check := false

	for _, user := range allUser {
		if username == user.GetUsername() {
			getIdFromUser = user.GetId()
			check = true
			break
		}
	}

	// Update data
	if check {
		// if data found or check is true
		getUser, err = s.repo.UpdateData(getIdFromUser, user)

		if err != nil {
			log.Println(err)
			return getUser, err
		}
		return getUser, nil

	} else {
		// when data not found or check is false
		errs := errors.New("there is an error when update user using username")
		return getUser, errs
	}

}

func (s *Service) UpdateUserByPassword(password string, user model.User) (model.User, error) {
	var allUser []model.User
	var getUser model.User
	var getIdFromUser string
	var err error

	allUser, err = s.repo.GetAllData()

	if err != nil {
		log.Println(err)
		return getUser, nil
	}

	for _, user := range allUser {
		if password == user.GetPassword() {
			getIdFromUser = user.GetId()
			break
		}
	}

	// Updatedata
	getUser, err = s.repo.UpdateData(getIdFromUser, user)

	if err != nil {
		log.Println(err)
		return getUser, err
	}

	return getUser, nil
}

func (s *Service) DeleteUserById(id string) (model.User, error) {
	var getUser model.User
	var err error

	getUser, err = s.repo.DeleteData(id)

	if err != nil {
		log.Println(err)
		return getUser, err
	}

	return getUser, nil
}

func (s *Service) DeleteUserByUsername(username string) (model.User, error) {
	var allUser []model.User
	var getUser model.User
	var err error

	// logic to find id at certain username
	allUser, err = s.repo.GetAllData()

	if err != nil {
		log.Println(err)
		return getUser, err
	}

	for _, user := range allUser {
		if username == user.GetUsername() {
			getUser = user
			break
		}
	}

	getUser, err = s.repo.DeleteData(getUser.GetUsername())

	if err != nil {
		log.Println(err)
		return getUser, err
	}

	return getUser, err
}

func (s *Service) DeleteUserByPassword(password string) (model.User, error) {
	var allUser []model.User
	var getUser model.User
	var err error

	// logic to find id at certain username
	allUser, err = s.repo.GetAllData()

	if err != nil {
		log.Println(err)
		return getUser, err
	}

	for _, user := range allUser {
		if password == user.GetPassword() {
			getUser = user
			break
		}
	}

	getUser, err = s.repo.DeleteData(getUser.GetUsername())

	if err != nil {
		log.Println(err)
		return getUser, err
	}

	return getUser, err
}
