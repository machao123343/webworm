package base

import (
	"net/http"
	"golang.org/x/crypto/openpgp/errors"
)

type Request struct {
	httpReq *http.Request   //HTTP请求的指针值
	depth uint32 //请求的深度
}

//创建新的请求
func NewRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{httpReq: httpReq, depth: depth}
}

//获取HTTP请求----函数名之前是返回类型的接收者
func (req *Request)  HttpReq() *http.Request {
	return req.httpReq
}

//获取深度值
func(req *Request) Depth() uint32 {
	return req.depth
}

type Response struct {
	httpResp *http.Response
	depth     uint32
}

//创建新的请求
func NewResponse(httpResp *http.Response, depth uint32) *Response {
	return  &Response{httpResp: httpResp, depth: depth}
}

//获取HTTP响应
func (resp *Response) HttpResp() *http.Response {
	return resp.httpResp
}

//获取深度值
func (resp *Response) Depth() uint32 {
	return  resp.depth
}

//条目
type Item map[string]interface{}

//数据接口
type Data interface{
	Valid() bool //数据是否有效
}

//数据是否有效
func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

//数据是否有效
func (resp *Response) Vaild() bool {
	return resp.httpResp != nil && resp.httpResp.Body != nil
}

//条目数据是否有效
func (item Item) Valid() bool {
	return item != nil
}

//func Error() string








