package format

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/wecatch/devops-ui/utils/log"
)

var logger = log.MakeLogger()

func init() {
	logger = log.Logger("utils.format")
}

// JSONResponseWriter for json api response
type JSONResponseWriter struct {
	*httptest.ResponseRecorder
}

//NewJSONResponseWriter for JSONResponseWriter init
func NewJSONResponseWriter() *JSONResponseWriter {
	return &JSONResponseWriter{
		ResponseRecorder: httptest.NewRecorder(),
	}
}

func formatResp(code int, message string, data []byte) []byte {
	var codeByte = []byte(`{"code":` + fmt.Sprintf("%d", code) + `,`)
	var messageByte = []byte(`"message":"` + message + `",`)
	var respByte = []byte(`"resp":`)
	var jsonData []byte
	if code == 0 {
		jsonData = bytes.Join([][]byte{codeByte, messageByte, respByte, data, []byte(`}`)}, []byte(""))
	} else {
		jsonData = bytes.Join([][]byte{codeByte, messageByte, respByte, []byte(`{}`), []byte(`}`)}, []byte(""))
	}

	return jsonData
}

//FormatResponseMiddleware for format response json data
func FormatResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/app/deploy/poll/log" {
			jsonRw := NewJSONResponseWriter()
			var jsonData []byte
			defer func() {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				if err := recover(); err != nil {
					jsonData = formatResp(1, fmt.Sprintf("%s", err), nil)
				} else {
					data := jsonRw.Body.Bytes()
					if len(data) == 0 {
						data = []byte("{}")
					}
					jsonData = formatResp(0, "success", data)
				}
				w.Header().Set("Content-Length", fmt.Sprintf("%d", len(jsonData)))
				w.Write(jsonData)
			}()
			next.ServeHTTP(jsonRw, r)
			// copy the original headers
			for k, v := range jsonRw.Header() {
				w.Header()[k] = v
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
