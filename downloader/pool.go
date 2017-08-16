package downloader

import (
	mdw "webworm/middleware"
	"reflect"
)
//生成网页下载器的函数类型
type GenPageDownloader func() PageDownLoader

//网页下载器池的接口类型
type PageDownloaderPool interface {
	Take() (PageDownLoader, error)
	Return(dl PageDownLoader) error
	Total() uint32
	Used() uint32
}

//创建网页下载器池
func NewPageDownloaderPool(
	total uint32,
	gen GenPageDownloader) (PageDownloaderPool, error) {
	etype := reflect.TypeOf(gen())
}
