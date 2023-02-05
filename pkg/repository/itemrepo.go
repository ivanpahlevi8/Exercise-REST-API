package repository

import (
	"database/sql"
	"errors"
	"exercise-web/pkg/model"
	"fmt"
	"log"
)

var MyItemRepo ItemRepo

type ItemRepo struct {
	Item *sql.DB
}

func InitItemRepo() ItemRepo {
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

	MyItemRepo.Item = result
	return MyItemRepo
}

func (r *ItemRepo) AddData(data model.Item) (model.Item, error) {
	getDataId := data.GetItemId()
	getDataName := data.GetItemName()
	getDataPrice := data.GetItemPrice()
	getDataDate := data.GetItemDate()

	fmt.Printf("Adding Data From Database, \nData Id : %s \nData Name : %s \nData Price : %s \nDataDate : %s \n\n",
		getDataId, getDataName, getDataPrice, getDataDate)

	myQuery := `insert into "Item"("item_id", "item_name", "item_price", "item_date") values($1, $2, $3, $4)`

	// insert to database
	_, err := r.Item.Exec(myQuery, getDataId, getDataName, getDataPrice, getDataDate)

	if err != nil {
		log.Println(err)
		var getData model.Item
		return getData, err
	} else {
		return data, nil
	}
}

func (r *ItemRepo) GetData(id string) (model.Item, error) {
	var getItem model.Item

	myQuery := `SLEECT * FROM "Item" where "item_id"=$1`

	allData, err := r.Item.Query(myQuery, id)

	if err != nil {
		log.Println(err)
		return getItem, err
	}

	// Check if data exist in database
	check := false

	for allData.Next() {
		var getItemId string
		var getItemName string
		var getItemPrice string
		var getItemDate string

		err = allData.Scan(&getItemId, &getItemName, &getItemPrice, &getItemDate)

		if err != nil {
			log.Println(err)
			return getItem, err
		}

		if id == getItemId {
			// data exist
			check = true
			getItem.SetItemId(getItemId)
			getItem.SetItemName(getItemName)
			getItem.SetItemPrice(getItemPrice)
			getItem.SetItemDate(getItemDate)
			break
		}
	}

	if check {
		// if data exist
		return getItem, nil
	} else {
		errs := errors.New("error happen when getting data from database")
		return getItem, errs
	}
}

func (r *ItemRepo) UpdateData(id string, item model.Item) (model.Item, error) {
	var getItem model.Item

	// get update credensial from input
	getItemId := item.GetItemId()
	getItemName := item.GetItemName()
	getItemPrice := item.GetItemPrice()
	getItemDate := item.GetItemDate()

	// check if there is an user with cretain id
	check := false

	checkQuery := `SELECT * FROM "Item" where "item_id"=$1`

	allData, err := r.Item.Query(checkQuery, id)

	if err != nil {
		log.Println(err)
		return getItem, err
	}

	for allData.Next() {
		// check every data in database, to check if there is data with certain id
		var itemId string
		var itemName string
		var itemPrice string
		var itemDate string

		err := allData.Scan(&itemId, &itemName, &itemPrice, &itemDate)

		if err != nil {
			log.Println(err)
			return getItem, err
		}

		if id == itemId {
			// if there is data with same id, then change checkn to true and break
			check = true
			break
		}
	}

	if check {
		// if data with certain id exist
		// do update logic
		myQuery := `update "Item" set "item_id"=$1, "item_name"=$2, "item_price"=$3, "item_date"=$4 where "item_id"=$5`

		_, e := r.Item.Query(myQuery, getItemId, getItemName, getItemPrice, getItemDate, id)

		if e != nil {
			log.Println(e)
			return getItem, err
		}
		getItem = item
		return getItem, nil
	} else {
		// if data is not found, then should return error
		errs := errors.New("there is an error when updating data in database")
		return getItem, errs
	}
}

func (i *ItemRepo) DeleteData(id string) (model.Item, error) {
	var getItem model.Item

	// check if there is data with certain id
	check := false

	checkQuery := `SELECT * FROM "Item"`

	allData, err := i.Item.Query(checkQuery)

	if err != nil {
		log.Println(err)
		return getItem, err
	}

	for allData.Next() {
		var itemId string
		var itemName string
		var itemPrice string
		var itemDate string

		e := allData.Scan(&itemId, &itemName, &itemPrice, &itemDate)

		if e != nil {
			log.Println(e)
			return getItem, e
		}

		if itemId == id {
			// if data with certain id founded
			check = true
			getItem.SetItemId(itemId)
			getItem.SetItemName(itemName)
			getItem.SetItemPrice(itemPrice)
			getItem.SetItemDate(itemDate)
			break
		}
	}

	// do some delete logic in here
	if check {
		// if there is an data with certain id
		deleteQuery := `delete from "Item" where "item_id"=$1`

		_, e := i.Item.Query(deleteQuery, id)

		if e != nil {
			log.Println(e)
			var falseItem model.Item
			return falseItem, e
		}

		return getItem, nil
	} else {
		// if datan was'n found
		errs := errors.New("error happen when deleted item in database")
		var falseItem model.Item
		return falseItem, errs
	}
}
