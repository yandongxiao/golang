package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/yandongxiao/go/lib/gomock/tree/spider"
)

// 在单元测试中再也不用先去实现一个Spider接口了
func TestGetGoVersion(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockSpider := spider.NewMockSpider(mockCtl)
	mockSpider.EXPECT().GetBody().Return("go1.8.3")
	goVer := GetGoVersion(mockSpider)

	if goVer != "go1.8.3" {
		t.Errorf("Get wrong version %v", goVer)
	}
}
