package components

import (
	Logger "gin-quick-start/internal/components/logger"
	"net/http"
	"strings"
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

const (
	SIMPLE ClinetType = "SIMPLE"
)

var (
	instanceCache sync.Map = sync.Map{}
)

type ClinetType string

type httpClientWrapper struct {
	Client   *resty.Client
	LogEntry *logrus.Entry
}

func SetUp() {
	instanceCache.Store(SIMPLE, createHttpClientWrapper(SIMPLE))
}

func GetInstance(clientType ClinetType) *resty.Client {
	instace, exist := instanceCache.Load(clientType)
	if exist {
		return instace.(*httpClientWrapper).Client
	}
	return nil
}

func createHttpClientWrapper(clientType ClinetType) *httpClientWrapper {
	entry := Logger.Logger.WithFields(logrus.Fields{
		"ClinetType": clientType,
	})
	return &httpClientWrapper{
		createRestyInstance(entry),
		entry,
	}
}

func createRestyInstance(logEntry *logrus.Entry) *resty.Client {
	client := resty.New()
	client.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		return printRequestLog(logEntry, c, r)
	})
	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		return printResponseLog(logEntry, r)
	})
	return client
}

// 打印客户端请求日志
func printRequestLog(logEntry *logrus.Entry, c *resty.Client, r *resty.Request) error {
	var builder strings.Builder
	builder.WriteString("[RestyHttpClient] Request: [ \n")
	builder.WriteString(formateHeaders(r.Header))
	builder.WriteString("Request Method: " + r.Method + "\n")
	builder.WriteString("Request URI: " + r.URL + "\n")
	pathParams, _ := c.JSONMarshal(r.PathParams)
	builder.WriteString("Request PathParams: " + string(pathParams) + "\n")
	builder.WriteString("Request QueryParam: " + r.QueryParam.Encode() + "\n")
	json, _ := c.JSONMarshal(r.Body)
	builder.WriteString("Request Body: " + string(json) + "\n")
	builder.WriteString("] \n")
	logEntry.Info(builder.String())
	return nil
}

// 打印客户端响应日志
func printResponseLog(logEntry *logrus.Entry, r *resty.Response) error {
	var builder strings.Builder
	builder.WriteString("[RestyHttpClient] Response: [ \n")
	builder.WriteString(formateHeaders(r.Header()))
	builder.WriteString("Response Method: " + r.Request.Method + "\n")
	builder.WriteString("Response URI: " + r.Request.URL + "\n")
	builder.WriteString("Response Status: " + r.Status() + "\n")
	builder.WriteString("Response Body: " + r.String() + "\n")
	builder.WriteString("] \n")
	logEntry.Info(builder.String())
	return nil
}

// 格式化请求头
func formateHeaders(headers http.Header) string {
	var headerBuilder strings.Builder
	headerBuilder.WriteString("Headers:[ \n")
	for key, values := range headers {
		headerBuilder.WriteString(key + ": " + strings.Join(values, ",") + "\n")
	}
	headerBuilder.WriteString("] \n")
	return headerBuilder.String()
}
