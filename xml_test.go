package main

import (
	"encoding/xml"
	"testing"
)

func TestUnmarshelHexString(t *testing.T) {
	d := []byte(`<type>73736e63</type>`)

	var tp hexString

	err := xml.Unmarshal(d, &tp)
	if err != nil {
		t.Error(err)
	}

	if tp != hexString("ssnc") {
		t.Errorf("Wrong hexString: %s", tp)
	}
}
func TestUnmarshelItem(t *testing.T) {
	d := []byte(`<item><type>73736e63</type><code>736e7561</code><length>39</length>
<data encoding="base64">
aVR1bmVzLzEyLjMuMyAoTWFjaW50b3NoOyBPUyBYIDEwLjEwLjUp</data></item>`)

	var curItem item

	err := xml.Unmarshal(d, &curItem)
	if err != nil {
		t.Error(err)
	}

	if curItem.Type != hexString("ssnc") {
		t.Errorf("Wrong type: %s", curItem.Type)
	}

	if curItem.Code != hexString("snua") {
		t.Errorf("Wrong code: %s", curItem.Code)
	}

	/*if bytes.Equals(curItem.Data, itemData("")) {
		t.Errorf("Wrong data: %s", curItem.Data)
	}*/

}
