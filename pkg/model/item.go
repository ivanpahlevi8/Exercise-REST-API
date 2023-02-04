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
