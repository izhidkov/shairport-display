package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"fmt"
)

func _decodeBase64String(in string) (string, error) {
	out, err := hex.DecodeString(in)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

type itemData []byte

func (item *itemData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tmp string
	err := d.DecodeElement(&tmp, &start)
	if err != nil {
		return err
	}

	var encoding string
	for _, attr := range start.Attr {
		if attr.Name.Local == "encoding" {
			encoding = attr.Value
		}
	}

	if encoding == "base64" {
		*item, err = base64.StdEncoding.DecodeString(tmp)
	} else {
		*item = itemData(tmp)
	}
	return err
}

type item struct {
	Code itemCode `xml:"code"`
	Type itemType `xml:"type"`
	Data itemData `xml:"data"`
}

type b64String string

func (el *b64String) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tmp string
	err := d.DecodeElement(&tmp, &start)
	if err != nil {
		return err
	}
	decodedString, err := _decodeBase64String(tmp)
	*el = b64String(decodedString)
	return err
}

func (it item) String() string {
	return fmt.Sprintf("%s: %s", it.Code, it.Data)
}

type itemType b64String
type itemCode b64String
