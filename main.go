package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const artworkFilepath = "/tmp/shairport-sync-picture.jpg"

type metadataRecord struct {
	Title  string
	Artist string
	Album  string

	HasArtwork bool
}

func main() {
	fifo, err := os.Open("/tmp/shairport-sync-metadata")
	if err != nil {
		log.Fatal(err)
	}
	defer fifo.Close()

	decoder := xml.NewDecoder(fifo)

	var curItem item
	var rec metadataRecord

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
				switch curItem.Code {

				// Metadata stream start
				case itemCode("mdst"):
					rec = metadataRecord{}

				// Metadata stream end
				case itemCode("mden"):
					DisplayRecord(rec)
					rec = metadataRecord{}

				case itemCode("asar"):
					rec.Artist = string(curItem.Data)

				case itemCode("minm"):
					rec.Title = string(curItem.Data)

				case itemCode("asal"):
					rec.Album = string(curItem.Data)

				// Album artwork
				case itemCode("PICT"):
					fmt.Printf("Saved %s\n", artworkFilepath)
					err := ioutil.WriteFile(artworkFilepath, curItem.Data, 0644)
					if err != nil {
						log.Fatal(err)
					}
					rec.HasArtwork = true

				default:
					fmt.Println(curItem)
				}
			}
		}
	}
}

func DisplayRecord(rec metadataRecord) {
	fmt.Println(rec)
}
