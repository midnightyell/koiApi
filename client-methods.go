package koiApi

type Client interface {
	GetResponse() string
	PrintError()
}
