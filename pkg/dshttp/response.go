package dshttp

import (
	"encoding/json"
	"net/http"
)

type responseStatus string

const (
	responseStatusSuccess responseStatus = "success" // all went well with (usually) payload returned
	responseStatusFail    responseStatus = "fail"    // there was a problem with the data submitted or some pre-condition of the API call wasn't satisfied
	responseStatusError   responseStatus = "error"   // an error occurred in processing the request, i.e. an exception was thrown
)

// Response is API response format
type Response struct {
	Status  responseStatus `json:"status"`            // response status
	Payload interface{}    `json:"payload,omitempty"` // response payload (any data)
	Meta    interface{}    `json:"meta,omitempty"`    // meta information (pagination, objects count, etc)
	Code    string         `json:"code,omitempty"`    // error code
	Message string         `json:"message,omitempty"` // error message
}

var (
	errorInternalServerJSON = []byte(`{"status":"error","code":"internal-error","message":"Internal server error"}`)
	errorNotFoundJSON       = []byte(`{"status":"fail","code":"not-found","message":"Not found"}`)
	emptyJSON               = []byte(`{"status":"success"}`)
)

// InternalServerErrorResponse writes a "500 Internal Server Error" response
func InternalServerErrorResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(errorInternalServerJSON)
}

// NotFoundErrorResponse writes a "404 Not Found" response
func NotFoundErrorResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusNotFound)
	w.Write(errorNotFoundJSON)
}

// JSONEmptyResponse writes an empty successful json response
func JSONEmptyResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(emptyJSON)
}

// FailResponse writes a fail to response stream
func FailResponse(w http.ResponseWriter, code, message string, payload interface{}) {
	resp := Response{Status: responseStatusFail}
	if code != "" {
		resp.Code = code
	}
	if message != "" {
		resp.Message = message
	}
	if payload != nil {
		resp.Payload = payload
	}

	e, err := json.Marshal(resp)
	if err != nil {
		InternalServerErrorResponse(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(e)
}

// ErrorResponse writes an error to response stream
func ErrorResponse(w http.ResponseWriter, code, message string, payload interface{}, httpCode int) {
	resp := Response{Status: responseStatusError}
	if code != "" {
		resp.Code = code
	}
	if message != "" {
		resp.Message = message
	}
	if payload != nil {
		resp.Payload = payload
	}

	e, err := json.Marshal(resp)
	if err != nil {
		InternalServerErrorResponse(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(httpCode)
	w.Write(e)
}

// JSONResponse writes the given data as json
func JSONResponse(w http.ResponseWriter, payload interface{}, meta interface{}) {
	resp := Response{Status: responseStatusSuccess}
	if payload != nil {
		resp.Payload = payload
	}
	if meta != nil {
		resp.Meta = meta
	}

	b, err := json.Marshal(resp)
	if err != nil {
		InternalServerErrorResponse(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
