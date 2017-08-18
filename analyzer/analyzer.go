package analyzer

import (
	"logging"
	"errors"
	base "webworm/base"
	mdw "webworm/middleware"
	"net/url"
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
	


}


