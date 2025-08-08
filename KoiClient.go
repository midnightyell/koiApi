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
