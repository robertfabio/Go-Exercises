package main

import (
	"LearnGo/model"
	"fmt"
	"time"
)

func main() {

	var itemNames []string

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
}
