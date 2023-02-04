package repository

import (
	"database/sql"
	"errors"
	"exercise-web/pkg/model"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var MyRepo Repo

type Repo struct {
	account *sql.DB
}

func InitRepo() Repo {
	//inisialisasi data
	host := "127.0.0.1"
	port := "5432"
	user := "postgres"
	password := "03052001ivan"
	dbname := "MyDatabase"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//inisialisasi koneksi
	result, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("error = ", err)
	}

	// close when program stop

	MyRepo.account = result
	return MyRepo
}

func (r *Repo) AddData(account model.User) {
	getId := account.GetId()
	getPassword := account.GetPassword()
	getUsername := account.GetUsername()

	fmt.Printf("Get Account, \nId User : %s \nUsername : %s \nPassword : %s\n", getId, getPassword, getUsername)

	myQuery := `insert into "Student"("student_id", "student_username", "student_password") values($1, $2, $3)`

	_, err := r.account.Exec(myQuery, getId, getUsername, getPassword)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Success Adding Data To Database")
}

func (r *Repo) GetData(id string) (model.User, error) {
	var getUser model.User

	myQuery := `SELECT * FROM "Student" where "student_id"=$1`

	all, err := r.account.Query(myQuery, id)

	if err != nil {
		log.Println(err)
		return getUser, err
	}

	check := false

	for all.Next() {
		var getId string
		var getUsername string
		var getPassword string

		err = all.Scan(&getId, &getUsername, &getPassword)

		if err != nil {
			log.Println(err)
			return getUser, nil
		}

		if id == getId {
			getUser.SetId(getId)
			getUser.SetUsername(getUsername)
			getUser.SetPassword(getPassword)
			check = true
			break
		}
	}

	if check {
		return getUser, nil
	} else {
		errs := errors.New("there is an error when getting data")
		return getUser, errs
	}
}

func (r *Repo) UpdateData(id string, user model.User) (model.User, error) {
	var getData model.User

	// Credensial from user
	getId := user.GetId()
	getUsername := user.GetUsername()
	getPassword := user.GetPassword()

	// check first if data with certain id exist
	checkQuery := `SELECT * FROM "Student" where "student_id"=$1`
	all, err := r.account.Query(checkQuery, id)
	if err != nil {
		log.Println(err)
	}
	check := false
	for all.Next() {
		var getId string
		var getUsername string
		var getPassword string

		err = all.Scan(&getId, &getUsername, &getPassword)

		if err != nil {
			log.Println(err)
		}

		if id == getId {
			check = true
			break
		}
	}

	// do updating logic
	if check {
		myQuery := `update "Student" set "student_id"=$1, "student_username"=$2, "student_password"=$3 where "student_id"=$4`
		fmt.Println("Student Id in database input : ", id)

		_, e := r.account.Query(myQuery, getId, getUsername, getPassword, id)

		if e != nil {
			log.Println(e)
			return getData, e
		}

		getData = user

		return getData, nil
	} else {
		errs := errors.New("there is an error when updating data")
		return getData, errs
	}
}

func (r *Repo) DeleteData(id string) (model.User, error) {
	var getData model.User

	// check first if data with certain id exist
	checkQuery := `SELECT * FROM "Student" where "student_id"=$1`
	all, err := r.account.Query(checkQuery, id)
	if err != nil {
		log.Println(err)
	}
	check := false
	for all.Next() {
		var getId string
		var getUsername string
		var getPassword string

		err = all.Scan(&getId, &getUsername, &getPassword)

		if err != nil {
			log.Println(err)
		}

		if id == getId {
			check = true
			break
		}
	}

	if check {
		myQuery := `delete from "Student" where "student_id"=$1`

		_, err2 := r.account.Exec(myQuery, id)

		if err2 != nil {
			log.Println(err)
			return getData, err
		}

		getData, err = r.GetData(id)

		if err != nil {
			log.Println(err)
			return getData, err
		}

		return getData, nil
	} else {
		errs := errors.New("there is an error when delete data")
		return getData, errs
	}

}

func (r *Repo) GetAllData() ([]model.User, error) {
	var allData []model.User

	myQuery := `SELECT * FROM "Student" `

	all, err := r.account.Query(myQuery)

	if err != nil {
		log.Println(err)
		return allData, err
	}

	for all.Next() {
		var getUser model.User

		err = all.Scan(&getUser.Id, &getUser.Username, &getUser.Password)

		if err != nil {
			log.Println(err)
			return allData, err
		}

		allData = append(allData, getUser)
	}

	return allData, nil
}
