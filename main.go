package main

import (
	"encoding/xml"
	"fmt"
	"github.com/ajstarks/openvg"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const artworkFilepath = "/tmp/shairport-sync-picture.jpg"

type metadataRecord struct {
	Title  string
	Artist string
	Album  string

	HasArtwork bool
	OnStop bool
}

var width, height int

func main() {
	fifo, err := os.Open("/tmp/shairport-sync-metadata")
	if err != nil {
		log.Fatal(err)
	}
	defer fifo.Close()

	decoder := xml.NewDecoder(fifo)

	items := make(chan item)
	go CollateItems(items)

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "item" {
				var curItem item
				err = decoder.DecodeElement(&curItem, &se)
				if err != nil {
					log.Fatal(err)
				}
				items <- curItem
			}
		}
	}
	openvg.Finish()
}

func CollateItems(items <-chan item) {
	width, height = openvg.Init()
	defer openvg.Finish()
	log.Printf("Initialized display size: %d x %d\n", width, height)
	// openvg.Start(width, height)
	// openvg.BackgroundColor("black")
	// openvg.End()

	timer := time.NewTimer(time.Second)
	var rec metadataRecord

	for {
		select {
		case <- timer.C:
			if (rec != metadataRecord{}) {
				DisplayRecord(rec)
				rec = metadataRecord{}
			}

		case curItem := <-items:
			timer.Reset(time.Second)

			switch curItem.Code {
			case hexString("asar"):
				fmt.Println(curItem)
				rec.Artist = string(curItem.Data)

			case hexString("minm"):
				fmt.Println(curItem)
				rec.Title = string(curItem.Data)

			case hexString("asal"):
				fmt.Println(curItem)
				rec.Album = string(curItem.Data)

			// Album artwork
			case hexString("PICT"):
				fmt.Printf("Saved %s\n", artworkFilepath)
				err := ioutil.WriteFile(artworkFilepath, curItem.Data, 0644)
				if err != nil {
					log.Fatal(err)
				}
				rec.HasArtwork = true
			case hexString("pbeg"):
				rec.OnStop = false
			case hexString("prsm"):
				rec.OnStop = false
			case hexString("pend"):
				log.Printf("clear window on stop")
				rec.OnStop = true

			default:
				fmt.Println(curItem)
			}
		}
	}
}

func DisplayRecord(rec metadataRecord) {
	fmt.Println(rec)
	if (!rec.OnStop) {
		openvg.Start(width, height)
		openvg.BackgroundColor("black")
		openvg.FillColor("rgb(255,255,255)")

		// FIXME: better scaling as per this example:
		// https://github.com/ajstarks/openvg/blob/master/go-client/picshow/picshow.go
		if (rec.HasArtwork) {
			openvg.Image(
				openvg.VGfloat(width-height/2)*0.51,
				openvg.VGfloat(height)*0.31,
				height/2,
				height/2,
				artworkFilepath,
			)
		}

		textX := openvg.VGfloat(width) * 0.5
		textY := openvg.VGfloat(height) * 0.2
		size := width / 40
		openvg.TextMid(textX, textY, rec.Title, "sans", size)
		openvg.TextMid(textX, textY-1.2*openvg.VGfloat(size), rec.Artist + " - " + rec.Album, "sans", size/2)
		// openvg.TextMid(textX, textY+4*openvg.VGfloat(size), rec.Album, "sans", size/2)

		openvg.End()
	} else {
		log.Printf("clear window on stop finish")
		openvg.Finish()
	}
}
