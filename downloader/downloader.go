package downloader

import (
	"net/http"
	"logging"

	base "webworm/base"
	mdw "webworm/middleware"
)

//func Download(req base.Request) (*base.Response, error)

//网页下载器的接口类型
type PageDownloader interface {
	Id() uint32 //获得ID
	Download(req base.Request) (*base.Response, error) //根据请求下载网页并回应
}

//网页下载器的实现类型
type myPageDownloader struct {
	id          uint32      //ID
	httpClient  http.Client //HTTP客户端
}
//创建日志生成器
var logger logging.Logger = base.NewLogger()

//ID生成器
var downloaderIdGenerator mdw.IDGenerator = mdw.NewIdGenerator()//方法实现类型

//生成并返回ID
func genDownloaderId() uint32 {
	return downloaderIdGenerator.GetUint32()
}

//创建网页下载器 todo:golang实现请求远程网页使用client提供的方法
func NewPageDownloader(client *http.Client) PageDownloader {
	id := genDownloaderId()
	if client == nil {
		client = &http.Client{}
	}
	return &myPageDownloader{
		id:         id,
		httpClient: *client,
	}
}
//Client是一个方法集合接口，能实现DownLoad的类型

//实现网络下载器接口里的方法
func (dl *myPageDownloader) Id() uint32 {
	return dl.id
}

func (dl *myPageDownloader) Download(req base.Request) (*base.Response, error) {
	httpReq := req.HttpReq()
	logger.Infof("Do the request (url=%s)...\n", httpReq.URL)
	httpResp, err := dl.httpClient.Do(httpReq)
    if err != nil {
		return nil, err
    }
	return base.NewResponse(httpResp, req.Depth()), nil
}





