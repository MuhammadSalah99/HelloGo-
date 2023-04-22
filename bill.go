package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//make new bill\

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

func (b *bill) formatBill() string {
	fs := "Bil breakdown \n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	//add tip

	fs += fmt.Sprintf("%-25v ... $%v \n", "tip:", b.tip)

	//total

	fs += fmt.Sprintf("%-25v ... $%0.2f \n	", "total:", total+b.tip)

	return fs

}

// update tip

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

//add item to the bill

func (b *bill) addItems(name string, price float64) {
	b.items[name] = price
}

//save bill

func (b *bill) save() {
	data := []byte(b.formatBill())

	err := os.WriteFile("Bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file")
}
