package koiApi

import (
	"fmt"
	"time"
)

var (
	defaultClient *koiClient
)

func GetClient() *koiClient {
	// if len(c) > 0 {
	// 	if defaultClient == nil {
	// 		defaultClient =
	// 	}
	// 	return c[0]
	// }

	if defaultClient == nil {
		defaultClient = NewKoiClient("", 30*time.Second)
		_, err := defaultClient.CheckLogin()
		if err != nil {
			panic(fmt.Sprintf("Login failed: %v", err))
		}
	}
	return defaultClient
}

type foo struct {
	item  string
	value string
}

func buildDatumIndex(d []*Datum) bool {
	var index map[string]string

	for _, obj := range d {
		index[obj.Label] = foo{value: obj.Value, item: obj.Item}
	}

	if defaultClient == nil {
		defaultClient = NewKoiClient("", 30*time.Second)
	}
	return true
}

// DatumExists checks if a datum with a specific label and value exists.
func (c *koiClient) DatumExists(label, value string) (bool, error) {
	var data []Datum
	err := c.listResources("/api/data", &data, fmt.Sprintf("label=%s", label), fmt.Sprintf("value=%s", value))
	if err != nil {
		return false, err
	}

	return len(data) > 0, nil
}
