package koiApi

var (
	defaultClient *Client
)

func GetClient(*c ...koiClient) *Client {
	if len(c) > 0 {
		if defaultClient == nil {
			defaultClient = c[0]
		}
		return c[0]
	}
	if defaultClient == nil {
		defaultClient = NewKoiClient( "", 30 * time.Second )


}