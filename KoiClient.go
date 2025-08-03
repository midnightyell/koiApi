package koiApi

var (
	defaultClient *koiClient
)

func GetClient(*c ...koiClient) *koiClient {
	if len(c) > 0 {
		if defaultClient == nil {
			defaultClient = c[0]
		}
		return c[0]
	}
	if defaultClient == nil {
		defaultClient = NewKoiClient( "", 30 * time.Second )
		res, err := defaultClient.CheckLogin()


}