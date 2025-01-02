package middleware

import (
	"bytes"
	Logger "gin-quick-start/internal/components/logger"
	contextManager "gin-quick-start/internal/context"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	ingorePaths = []string{"/swagger", "/docs"}
)

type copyWriter struct {
	gin.ResponseWriter
	response *bytes.Buffer
}

func (w copyWriter) Write(data []byte) (int, error) {
	_, _ = w.response.Write(data)
	return w.ResponseWriter.Write(data)
}

func hasIngorePath(path string) bool {
	for _, ingorepath := range ingorePaths {
		if strings.HasPrefix(path, ingorepath) {
			return true
		}
	}
	return false
}

func LogPrinter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if hasIngorePath(c.Request.RequestURI) {
			c.Next()
			return
		}
		start := time.Now()
		requestId := contextManager.GetRequestId(c)
		entry := Logger.Logger.WithFields(logrus.Fields{
			REQUEST_ID_KEY: requestId,
		})
		entry.Info(resolveRequest(c))
		writer := copyWriter{c.Writer, bytes.NewBuffer([]byte{})}
		c.Writer = writer
		contextManager.SetLogger(c, entry)
		c.Next()
		end := time.Now()
		entry.Info(resolveResponse(c, writer.response.Bytes(), int(end.UnixMilli()-start.UnixMilli())))
	}
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

// 打印请求体
func resolveRequest(c *gin.Context) string {
	requestBody, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
	var builder strings.Builder
	builder.WriteString("SystemLog: [ \n")
	builder.WriteString(formateHeaders(c.Request.Header))
	builder.WriteString("Request Method: " + c.Request.Method + "\n")
	builder.WriteString("Request URI: " + c.Request.RequestURI + "\n")
	builder.WriteString("Request Body: " + string(requestBody) + "\n")
	builder.WriteString("] \n")
	return builder.String()
}

// 打印响应体
func resolveResponse(c *gin.Context, response []byte, executionTime int) string {
	var builder strings.Builder
	builder.WriteString("SystemLog: [ \n")
	builder.WriteString("Response Method: " + c.Request.Method + "\n")
	builder.WriteString("Response URI: " + c.Request.RequestURI + "\n")
	builder.WriteString("Response Body: " + string(response) + "\n")
	builder.WriteString("Response Status: " + strconv.Itoa(c.Writer.Status()) + "\n")
	builder.WriteString("Execution Time: " + strconv.Itoa(executionTime) + "ms \n")
	builder.WriteString("] \n")
	return builder.String()
}
