package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fifo, err := os.Open("/tmp/shairport-sync-metadata")
	if err != nil {
		log.Fatal(err)
	}
	defer fifo.Close()

	decoder := xml.NewDecoder(fifo)

	var curItem item

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "item" {
				err = decoder.DecodeElement(&curItem, &se)
				if err != nil {
					log.Fatal(err)
				}
				if curItem.Code == itemCode("PICT") {
					fmt.Println("Saved picture.jpg")
					err := ioutil.WriteFile("picture.jpg", curItem.Data, 0644)
					if err != nil {
						log.Fatal(err)
					}
				} else {
					fmt.Println(curItem)
				}
			}
		default:
		}

	}

}
