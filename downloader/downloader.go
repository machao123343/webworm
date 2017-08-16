package downloader

import "webworm/base"

//func Download(req base.Request) (*base.Response, error)

//网页下载器的接口类型
type PageDownLoader interface {
	Id() uint32 //获得ID
	Download(req base.Request) (*base.Response, error) //根据请求下载网页并回应
}



