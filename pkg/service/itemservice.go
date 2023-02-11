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
	var allItem []model.Item
	var getItem model.Item
	var err error

	// make logci to find id based on price
	var id string
	allItem, err = s.ItemRepo.GetAllData()

	if err != nil {
		log.Println(err)
		return getItem, err
	}

	check := false

	for _, item := range allItem {
		if price == item.GetItemPrice() {
			check = true
			id = item.GetItemId()
			break
		}
	}

	if check {
		// if data exist
		getItem, err = s.ItemRepo.GetData(id)
		if err != nil {
			log.Println(err)
			return getItem, err
		} else {
			return getItem, nil
		}
	} else {
		// if data not exist
		errs := errors.New("get error when getting data with price")
		return getItem, errs
	}
}

func (s *ItemService) GetItemByDate(date string) (model.Item, error) {
	// make variable
	var allItem []model.Item
	var err error
	var getItem model.Item

	// make logic to get id at certain date of data
	var getId string

	allItem, err = s.ItemRepo.GetAllData()

	if err != nil {
		log.Println(err)
		return getItem, err
	}

	check := false

	for _, item := range allItem {
		if date == item.GetItemDate() {
			check = true
			getId = item.GetItemId()
			break
		}
	}

	if check {
		// if data with certain date founded
		getItem, err = s.ItemRepo.GetData(getId)

		if err != nil {
			log.Println(err)
			return getItem, err
		} else {
			return getItem, nil
		}
	} else {
		errs := errors.New("error happen when getting data using date")
		return getItem, errs
	}
}

func (s *ItemService) UpdateDataById(id string, inputItem model.Item) (model.Item, error) {
	// make variable
	var getItem model.Item
	var err error

	// do update
	getItem, err = s.ItemRepo.UpdateData(id, inputItem)

	if err != nil {
		log.Println(err)
		return getItem, err
	} else {
		return getItem, nil
	}
}

func (s *ItemService) UpdateDataByName(name string, inputItem model.Item) (model.Item, error) {
	// make variable
	var allItem []model.Item
	var getItem model.Item
	var err error

	// make logic to get data based on name
	var getId string

	allItem, err = s.ItemRepo.GetAllData()

	if err != nil {
		log.Println(err)
		return getItem, err
	}

	check := false

	for _, item := range allItem {
		if name == item.GetItemName() {
			check = true
			getId = item.GetItemId()
			break
		}
	}

	if check {
		// if item exist
		getItem, err = s.ItemRepo.UpdateData(getId, inputItem)

		if err != nil {
			log.Println(err)
			return getItem, err
		} else {
			return getItem, nil
		}
	} else {
		errs := errors.New("error when updating data using item name")
		return getItem, errs
	}
}

func (s *ItemService) UpdateItemByPrice(price string, inputItem model.Item) (model.Item, error) {
	// make variable
	var allItem []model.Item
	var getItem model.Item
	var err error

	// make logic to get item using price
	var getId string

	allItem, err = s.ItemRepo.GetAllData()

	if err != nil {
		log.Println(err)
		return getItem, err
	}

	check := false

	for _, item := range allItem {
		if price == item.GetItemPrice() {
			check = true
			getId = item.GetItemId()
			break
		}
	}

	if check {
		// if data founded
		getItem, err = s.ItemRepo.UpdateData(getId, inputItem)

		if err != nil {
			log.Println(err)
			return getItem, err
		} else {
			return getItem, nil
		}
	} else {
		errs := errors.New("error when updating item using price")
		return getItem, errs
	}
}

func (s *ItemService) UpdateItemByDate(date string, item model.Item) (model.Item, error) {
	// create variable
	var allItem []model.Item
	var getItem model.Item
	var err error

	// make logic to get item using date
	var getId string

	allItem, err = s.ItemRepo.GetAllData()

	if err != nil {
		log.Println(err)
		return getItem, err
	}

	check := false

	for _, item := range allItem {
		if date == item.GetItemDate() {
			check = true
			getId = item.GetItemId()
			break
		}
	}

	if check {
		// if data founded
		getItem, err = s.ItemRepo.UpdateData(getId, item)

		if err != nil {
			log.Println(err)
			return getItem, err
		} else {
			return getItem, nil
		}
	} else {
		errs := errors.New("error when updating data using date")
		return getItem, errs
	}
}

func (s *ItemService) DeleteItemById(id string) (model.Item, error) {
	// make variable to keep file
	var getItem model.Item
	var err error

	// delete item using id
	getItem, err = s.ItemRepo.DeleteData(id)

	if err != nil {
		log.Println(err)
		return getItem, err
	} else {
		return getItem, nil
	}
}

func (s *ItemService) DeleteItemByName(name string) (model.Item, error) {
	// make variable
	var allItem []model.Item
	var getItem model.Item
	var err error

	// make algorithm to get id using name
	var getId string

	allItem, err = s.ItemRepo.GetAllData()

	if err != nil {
		log.Println(err)
		return getItem, err
	}
	check := false
	for _, item := range allItem {
		if name == item.GetItemName() {
			getId = item.GetItemId()
			check = true
			break
		}
	}

	if check {
		// if item founded
		getItem, err = s.ItemRepo.DeleteData(getId)
		if err != nil {
			log.Println(err)
			return getItem, err
		} else {
			return getItem, nil
		}
	} else {
		errs := errors.New("getting error when deleting data using name")
		return getItem, errs
	}
}

func (s *ItemService) DeleteItemByPrice(price string) (model.Item, error) {
	// make variable
	var allItem []model.Item
	var getItem model.Item
	var err error

	// make algorithm to get user id by price
	var getId string

	allItem, err = s.ItemRepo.GetAllData()

	if err != nil {
		log.Println(err)
		return getItem, err
	}

	check := false

	for _, item := range allItem {
		if price == item.GetItemPrice() {
			check = true
			getId = item.GetItemId()
			break
		}
	}

	if check {
		// if item with certain price founded
		getItem, err = s.ItemRepo.DeleteData(getId)
		if err != nil {
			log.Println(err)
			return getItem, err
		} else {
			return getItem, nil
		}
	} else {
		// item no founded
		errs := errors.New("errro when deleting data using price")
		return getItem, errs
	}
}

func (s *ItemService) DeleteItemUsingDate(date string) (model.Item, error) {
	// make variable
	var allItem []model.Item
	var getItem model.Item
	var err error

	// make logic to get id based on date
	var getId string

	allItem, err = s.ItemRepo.GetAllData()

	if err != nil {
		log.Println(err)
		return getItem, err
	}
	check := false
	for _, item := range allItem {
		if date == item.GetItemDate() {
			check = true
			getId = item.GetItemId()
			break
		}
	}

	if check {
		// if item exist
		getItem, err = s.ItemRepo.DeleteData(getId)
		if err != nil {
			log.Println(err)
			return getItem, err
		} else {
			return getItem, nil
		}
	} else {
		// not found
		errs := errors.New("getting error when deleting item using date")
		return getItem, errs
	}
}
