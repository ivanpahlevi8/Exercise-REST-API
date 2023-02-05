package service

import (
	"errors"
	"exercise-web/pkg/model"
	"exercise-web/pkg/repository"
	"fmt"
	"log"
)

var MyItemService ItemService

type ItemService struct {
	ItemRepo repository.ItemRepo
}

func InitItemService(itemRepo repository.ItemRepo) ItemService {
	MyItemService.ItemRepo = itemRepo

	return MyItemService
}

// logic about service item based on database start here

func (s *ItemService) AddItemData(item model.Item) (model.Item, error) {
	var getItem model.Item
	var err error

	getItem, err = s.ItemRepo.AddData(item)

	if err != nil {
		str := fmt.Sprintf("Getting error when adding data in service : %s\n", err)
		log.Println(str)
		return getItem, err
	} else {
		log.Println("SUccess adding item data in service")
		return getItem, nil
	}
}

func (s *ItemService) GetItemById(id string) (model.Item, error) {
	var getItem model.Item
	var err error

	getItem, err = s.ItemRepo.GetData(id)

	if err != nil {
		// if there is an error
		str := fmt.Sprintf("getting error when getting item using item id : %s\n", err)
		log.Println(str)
		return getItem, err
	} else {
		// if there is no error
		log.Println("Success getting data using item id")
		return getItem, nil
	}
}

func (s *ItemService) GetitemByName(name string) (model.Item, error) {
	// make variable
	var getItem model.Item
	var err error

	// make logic to find id based on item name
	var id string

	allData, err := s.ItemRepo.GetAllData()

	if err != nil {
		log.Println(err)
		return getItem, err
	}

	check := false // to check if data exist or not
	for _, data := range allData {
		if name == data.GetItemName() {
			// if data exist
			id = data.GetItemId()
			check = true
			break
		}
	}

	if check {
		// if data exist, get data from database
		getItem, err = s.ItemRepo.GetData(id)
		if err != nil {
			str := fmt.Sprintf("error happen when getting item by name : %s\n", err)
			log.Println(str)
			return getItem, err
		} else {
			return getItem, nil
		}
	} else {
		// if data does'nt exist
		errs := errors.New("error when getting item by name")
		return getItem, errs
	}
}

func (s *ItemService) GetItemByPrice(price string) (model.Item, error) {
	return model.Item{}, nil
}
