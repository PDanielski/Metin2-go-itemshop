package item

import (
	"mt2is/pkg/wallet"
	"time"
)

//Provider is used to retrieve items
type Provider interface {
	Provide() ([]*Item, error)
}

//Price is used to store multiple prices for one single item
type Price struct {
	prices map[string]int
}

//UsingCurrency returns the price using the given currency
func (p *Price) UsingCurrency(c *wallet.Currency) (int, bool) {
	price, ok := p.prices[c.Code]
	return price, ok
}

//Set sets the price using the given currency
func (p *Price) Set(c *wallet.Currency, amount int) {
	p.prices[c.Code] = amount
}

//Info holds the in game information of a Item
type Info struct {
	Vnum       uint
	Count      uint
	Socket0    uint
	Socket1    uint
	Socket2    uint
	Socket3    uint
	Socket4    uint
	Socket5    uint
	AttrType0  uint
	AttrValue0 uint
	AttrType1  uint
	AttrValue1 uint
	AttrType2  uint
	AttrValue2 uint
	AttrType3  uint
	AttrValue3 uint
	AttrType4  uint
	AttrValue4 uint
	AttrType5  uint
	AttrValue5 uint
	AttrType6  uint
	AttrValue6 uint
}

type Item struct {
	ID         int
	Info       Info
	InsertTime time.Time
	Price      Price
	Discount   Discount
}

//IsDiscounted checks if the item is discounted
func (i *Item) IsDiscounted() bool {
	return i.Discount.ID > 0
}

//ApplyDiscount applies the given discount to the item
func (i *Item) ApplyDiscount(d Discount) bool {
	if balance, ok := i.Price.UsingCurrency(d.Currency); ok {
		if d.AbsoluteModifier != 0 {
			balance = d.AbsoluteModifier
		}
		if d.RelativeModifier != 0 {
			balance = balance + d.RelativeModifier
		}
		if d.PercentageModifier != 0 {
			balance = int(float32(balance) * d.PercentageModifier)
		}
		i.Price.Set(d.Currency, balance)
		i.Discount = d
		return true
	}
	return false
}

//New creates a new Item object
func New(id int, info Info, insertTime time.Time, price Price) *Item {
	return &Item{ID: id, Info: info, InsertTime: insertTime, Price: price}
}
