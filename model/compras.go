package model

import (
	"errors"
	"time"
)

type Compra struct {
	Market   string
	Data     time.Time
	ItemBuys []Items
}

type Items struct {
	ItemName string
}

func NewCompra(market string, data time.Time, itemNames []string) (*Compra, error) {

	if market == "" {
		return nil, errors.New("market name cannot be empty")
	}

	if len(itemNames) == 0 {
		return nil, errors.New("item names cannot be empty")
	}

	var itens []Items

	for _, item := range itemNames {
		itens = append(itens, Items{ItemName: item})
	}

	return &Compra{
		Market:   market,
		Data:     time.Now(),
		ItemBuys: itens,
	}, nil
}


{/*

	itemNames = append(itemNames, "Arroz", "Feijão", "Macarrão")
	
	compra, err := model.NewCompra("Mercadinho", time.Now(), itemNames)
	
	if err != nil {
		fmt.Println("Error creating compra:", err)
		return
	} else {

	fmt.Println("Compra Details:")
	fmt.Println("Market:", compra.Market)
	fmt.Println("Date:", compra.Data.Format("2006-01-02"))
	fmt.Println("Items Bought:")
	
	for _, item := range compra.ItemBuys {
		fmt.Println("-", item.ItemName)
		}
	}
*/}