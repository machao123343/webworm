package analyzer

import(
	base "webworm/base"
	"net/http"
)

//被用于解析HTTP响应的函数类型----让爬虫使用者自定义响应规则
type ParseResponse func(httpResp *http.Response, respDepth uint32) ([]base.Data, []error)

