package main

import (
	"fmt"
	sp "github.com/voicurobert/golang-microservices/shipping_box_app/shipping_box"
)

func getBestBox(availableBoxes []sp.Box, products []sp.Product) sp.Box {
	totalVolume := 0
	for i := range products {
		totalVolume += sp.Volume(products[i].Parameter)
	}
	output := make(chan sp.Box)
	getBestBoxForVolume(totalVolume, availableBoxes, output)
	box := <-output
	close(output)
	return box
}

func getBestBoxForVolume(totalVolume int, boxes []sp.Box, output chan sp.Box) {
	var box sp.Box
	//var wg *sync.WaitGroup
	for i := range boxes {
		currentBox := boxes[i]
		go func() {
			if sp.Volume(currentBox.Parameter) > totalVolume {
				box = currentBox
				output <- box
			}
		}()
	}
	//return box
}

func main() {
	boxes := sp.GetBoxes()
	products := sp.GetProducts()
	box := getBestBox(boxes, products)
	fmt.Printf("best box is: %+v\n", box)
}
