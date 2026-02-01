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
