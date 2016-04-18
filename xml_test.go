package main

import (
	"encoding/xml"
	"testing"
)

func TestUnmarshelB64String(t *testing.T) {
	d := []byte(`<type>73736e63</type>`)

	var tp b64String

	err := xml.Unmarshal(d, &tp)
	if err != nil {
		t.Error(err)
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

}
