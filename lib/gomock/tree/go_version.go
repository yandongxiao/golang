package main

import (
	"tree/spider"
)

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
