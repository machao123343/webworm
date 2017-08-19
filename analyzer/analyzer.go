package analyzer

import (
	"logging"
	"errors"
	"net/url"
	"fmt"

	base "webworm/base"
	mdw "webworm/middleware"

)

//日志记录器
var logger logging.Logger = base.NewLogger()

//ID生成器
var analyzerIdGenerator mdw.IDGenerator = mdw.NewIdGenerator()

//生成并返回ID
func genAnalyzerId() uint32 {
	return analyzerIdGenerator.GetUint32()
}


type myAnalyzer struct {
	id uint32
}

//分析器的接口类型
type Analyzer interface {
	Id() uint32 //获得ID
	Analyze(
		respParsers  []ParseResponse,
		resp base.Response) ([]base.Data, []error)//根据规则分析响应并返回请求和条目
}

//创建分析器
func NewAnalyzer() Analyzer {
	return &myAnalyzer{id: genAnalyzerId()}
}

func (analyzer *myAnalyzer) Id() uint32 {//有返回时一定要声明返回的类型
	return analyzer.id
}

//添加请求值或条目值到列表
func appendDataList(dataList []base.Data, data base.Data, respDepth uint32) []base.Data {
	if data == nil {
		return dataList
	}
	req, ok := data.(*base.Request)//类型断言语句，判断数据是否可以解析为一个有效的HTTP请求
	if !ok {
		return append(dataList, data)
	}
	newDepth := respDepth + 1
	if req.Depth() != newDepth {
		req = base.NewRequest(req.HttpReq(), newDepth)
	}
	return append(dataList, req)
}

//添加错误值到列表
func appendErrorList(errorList []error, err error) []error {
	if err == nil {
		return errorList
	}
	return append(errorList, err)
}

func (analyzer *myAnalyzer) Analyze(
	respParsers []ParseResponse,
	resp base.Response) (datalist []base.Data, errorList []error) {
	if respParsers == nil {
		err := errors.New("The response parser list is invalid!")
		return  nil, []error{err}
	}

	httpResp := resp.HttpResp()
	if httpResp == nil {
		err := errors.New("The http reponse is invalid!")
		return nil, []error{err}
	}

	var reqUrl *url.URL = httpResp.Request.URL//复制请求的URL路径
	logger.Infof("Parse the response (reqUrl=%s)... \n", reqUrl)//Infof弄成一个接口型，什么都可以传进来。
	respDepth := resp.Depth()

	//解析HTTP响应
	datalist = make([]base.Data, 0)
	errorList = make([]error, 0)
	for i, respParser := range respParsers {
		if respParser == nil {
			err := errors.New(fmt.Sprintf("The document paser [%d] is invalid!", i))
			errorList = append(errorList, err)
			continue//执行下次循环内容
		}
		pDataList, pErrorList := respParser(httpResp, respDepth)
		if pDataList != nil {
			for _, pData := range pDataList {
				datalist = appendDataList(datalist, pData, respDepth)
			}
		}

		if pErrorList != nil {
			for _, pError := range pErrorList {
				errorList = appendErrorList(errorList, pError)
			}
		}
	}
	return datalist, errorList
}


