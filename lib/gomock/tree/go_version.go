package main

import (
	"github.com/yandongxiao/go/lib/gomock/tree/spider"
)

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
