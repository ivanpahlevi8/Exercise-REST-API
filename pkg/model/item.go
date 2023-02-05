package model

type Item struct {
	ItemId    string `db:"item_id"`
	ItemName  string `db:"item_name"`
	ItemPrice string `db:"item_price"`
	ItemDate  string `db:"item_date"`
}

// Get Set Method
func (i *Item) GetItemId() string {
	return i.ItemId
}

func (i *Item) GetItemName() string {
	return i.ItemName
}

func (i *Item) GetItemPrice() string {
	return i.ItemPrice
}

func (i *Item) GetItemDate() string {
	return i.ItemDate
}

func (i *Item) SetItemId(id string) {
	i.ItemId = id
}

func (i *Item) SetItemName(name string) {
	i.ItemName = name
}

func (i *Item) SetItemPrice(price string) {
	i.ItemPrice = price
}

func (i *Item) SetItemDate(date string) {
	i.ItemDate = date
}
