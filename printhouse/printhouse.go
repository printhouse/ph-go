// Package printhouse:
// go-printhouse provides an easy-to-use API
// to access Printhouse's WEB API
package printhouse

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"encoding/base64"
	"fmt"
)

// Printhouse struct which we use
// to wrap our request operations.
type Printhouse struct {
	clientID           string
	clientSecret       string
	accessToken        string
}

const (
	BASE_URL     = "https://api.printhouse.io"
	API_VERSION  = "v1"
)

// Creates a New Printhouse API object with the
// clientID and clientSecret
// Usage:
// 	printhouse.New("XXX","YYY")
func New(clientID, clientSecret string) Printhouse {

	return initialize(clientID, clientSecret)
}


func initialize(clientID, clientSecret string) Printhouse {
	ph := Printhouse{clientID: clientID, clientSecret: clientSecret}
	return ph
}

// Creates a new GET Request to Printhouse and returns
// the response as a map[string]interface{}.
//
// format: target endpoint format like "products/%s" - string
//
// data: content to be sent with the request - map[string]interface{}
//
// args: Arguments to be used based on format
//
// Usage:
// 	printhouse.Get("order/%s",nil,"0001")
func (printhouse *Printhouse) Get(format string, data map[string]interface{}, args ...interface{}) ([]byte, []error) {
	return printhouse.Request("GET", format, data, args...)
}

// Creates a new POST Request to Printhouse and returns
// the response as a map[string]interface{}.
//
// format: target endpoint format like "order/%s/confirm" - string
//
// data: content to be sent with the request - map[string]interface{}
//
// args: Arguments to be used based on format
//
// Usage:
// 	printhouse.Post("order/%s/confirm",map[string]interface{},"0001")
func (printhouse *Printhouse) Post(format string, data map[string]interface{}, args ...interface{}) ([]byte, []error) {
	return printhouse.Request("POST", format, data, args...)
}

// Creates a new PUT Request to Printhouse and returns
// the response as a map[string]interface{}.
//
// format: target endpoint format like "orders/%s" - string
//
// data: content to be sent with the request - map[string]interface{}
//
// args: Arguments to be used based on format
//
// Usage:
// 	printhouse.Put("orders/%s",nil,"00001")
func (printhouse *Printhouse) Put(format string, data map[string]interface{}, args ...interface{}) ([]byte, []error) {
	return printhouse.Request("PUT", format, data, args...)
}

// Creates a new Request to Printhouse and returns
// the response as a map[string]interface{}.
//
// method: GET/POST/PUT - string
//
// format: target endpoint format like "products/%s" - string
//
// data: content to be sent with the request - map[string]interface{}
//
// args: Arguments to be used based on format
//
// Usage:
// 	printhouse.request("GET","products/%s",nil,"00001")
func (printhouse *Printhouse) Request(method, format string, data map[string]interface{}, args ...interface{}) ([]byte, []error) {

	// create endpoint based on passed format
	endpoint := fmt.Sprintf(format, args...)

	targetURL := printhouse.createTargetURL(endpoint)

	request := gorequest.New()

	// Check method type to call corresponding
	// go-request method
	if method == "GET" { request.Get(targetURL) }
	if method == "POST" { request.Post(targetURL) }
	if method == "PUT" { request.Put(targetURL) }
	if method == "DELETE" { request.Delete(targetURL) }

	// Add the data to the request if it
	// isn't null
	if data != nil {
		jsonData, _ := getJsonBytesFromMap((data))
		if jsonData != nil { request.Send(string(jsonData)) }
	}

	_, body, errs := request.End()

	result := []byte(body)

	return result, errs
}

// Creates target URL for making a Printhouse Request
// to a given endpoint
func (printhouse *Printhouse) createTargetURL(endpoint string) string {
	result := fmt.Sprintf("%s/%s/%s", BASE_URL, API_VERSION, endpoint)
	return result
}

// returns base64 encoded authorization
// Keys for Printhouse.
func (printhouse *Printhouse) getEncodedKeys() string {

	data := fmt.Sprintf("%v:%v", printhouse.clientID, printhouse.clientSecret)
	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	return encoded;
}

// Extracts Json Bytes from map[string]interface
func getJsonBytesFromMap(data map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Invalid data object, can't parse to json:")
		fmt.Println("Error:", err)
		fmt.Println("Data:", data)
		return nil, err
	}
	return jsonData, nil
}
