package spider

//go:generate mockgen -destination mock_spider.go -package spider tree/spider Spider
type Spider interface {
	GetBody() string
}
