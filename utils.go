// Copyright (c) 2013, Nikolay Georgiev
// All rights reserved.

// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:

// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.

// * Redistributions in binary form must reproduce the above copyright notice, this
//   list of conditions and the following disclaimer in the documentation and/or
//   other materials provided with the distribution.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
// ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package sharky

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"unicode"
)

// ==================================== Util section ==================================

// Makes POST request to the API's method with params. SessionID should also
// be provided for some of the methods. You should also provide protocol (HTTP or HTTPS)
func makeCall(method string, params map[string]interface{}, sessionId, protocol, key, secret string) map[string]interface{} {
	response := getResponse(method, params, sessionId, protocol, key, secret)
	var resp Response
	json.Unmarshal(response, &resp)

	if resp.Errors != nil {
		error(resp.Errors, method)
	}

	return resp.Result
}

func makeSingleResultCall(method string, params map[string]interface{}, sessionId, protocol, key, secret string) interface{} {
	response := getResponse(method, params, sessionId, protocol, key, secret)
	var resp SingleResponse
	json.Unmarshal(response, &resp)

	if resp.Errors != nil {
		error(resp.Errors, method)
	}

	return resp.Result
}

func getResponse(method string, params map[string]interface{}, sessionId, protocol, key, secret string) []byte {
	reqData := buildRequestData(key, method, sessionId, params)
	buf, _ := json.Marshal(&reqData)
	signature := generateSignature(buf, []byte(secret))
	url := buildApiURL(signature, protocol)
	body := bytes.NewReader(buf)
	r, err := http.Post(url, CONTENT_TYPE, body)
	if err != nil {
		log.Panic(err)
	}
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}
	defer r.Body.Close()

	return response
}

func error(errors []map[string]interface{}, method string) {
	line := "======================="
	errMessage := fmt.Sprintf("\n%v\nError while executing %v()\n%v\n", line, method, line)
	for _, err := range errors {
		code := err["code"]
		msg := err["message"]
		data := err["data"]
		errMessage += fmt.Sprintf("Error Code: %v, %v [%v]\n", code, msg, data)
	}
	log.Panic(errMessage)
}

func buildRequestData(key, method, sessionID string, params map[string]interface{}) *RequestData {
	data := new(RequestData)
	data.Method = method
	data.Parameters = params

	header := make(map[string]string)
	header["wsKey"] = key
	header["sessionID"] = sessionID
	data.Header = header

	return data
}

// The signature is generated via HMAC using MD5 and the
// secret provided by Grooveshark team.
func generateSignature(postData, secret []byte) string {
	mac := hmac.New(md5.New, secret)
	mac.Write(postData)
	signature := fmt.Sprintf("%x", mac.Sum(nil))

	return signature
}

// Build the entire URL to the API. For some calls HTTPS
// protocol is not mandatory.
func buildApiURL(sig, protocol string) string {
	return protocol + API_HOST + API_ENDPOIT + SIG_GET_KEY + sig
}

// Util method to check empty values
func isEmpty(value string) bool {
	if len(strings.TrimSpace(value)) == 0 {
		return true
	} else {
		return false
	}
}

func extractNonEmptyStrings(arr []interface{}) []string {
	words := make([]string, 0)
	for _, val := range arr {
		if word, ok := val.(string); ok && !isEmpty(word) {
			words = append(words, word)
		}
	}
	return words
}

func md5sum(value string) string {
	h := md5.New()
	io.WriteString(h, value)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func logMsg(result map[string]interface{}, sucMsg, errMsg string) {
	if suc, ok := result["success"].(bool); ok {
		if suc {
			log.Println(sucMsg)
		} else {
			log.Panic(errMsg)
		}
	}
}

// ==================================== Reflection section ==================================

func mapToStruct(params map[string]interface{}, elem *reflect.Value) {
	for k, v := range params {
		if _, ok := v.([]interface{}); ok {
			// skip if the value v is array
			continue
		}
		setFieldOfElem(elem, k, v)
	}
}

func setFieldOfElem(elem *reflect.Value, key string, val interface{}) {
	field := elem.FieldByName(firstToUpper(key))

	if !field.CanSet() {
		return
	}

	switch field.Kind() {
	case reflect.String:
		iVal := getInt64(val)
		if iVal != -1 {
			field.SetString(strings.TrimSpace(fmt.Sprintf("%v", iVal)))
		} else {
			field.SetString(strings.TrimSpace(fmt.Sprintf("%v", val)))
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		log.Println("KEY: " + key)
		v := getInt64(val)
		field.SetInt(v)
	case reflect.Float32, reflect.Float64:
		v := getFloat64(val)
		field.SetFloat(v)
	case reflect.Bool:
		if v, ok := val.(bool); ok {
			field.SetBool(v)
		}
	}
}

func firstToUpper(value string) string {
	uni := []rune(value)
	uni[0] = unicode.ToUpper(uni[0])
	value = string(uni)
	return value
}

func getFloat64(value interface{}) float64 {
	if v, ok := value.(float32); ok {
		return float64(v)
	}
	if v, ok := value.(float64); ok {
		return v
	}
	return 0
}

func getInt64(value interface{}) int64 {
	if v, ok := value.(int); ok {
		return int64(v)
	}
	if v, ok := value.(int8); ok {
		return int64(v)
	}
	if v, ok := value.(int16); ok {
		return int64(v)
	}
	if v, ok := value.(int32); ok {
		return int64(v)
	}
	if v, ok := value.(rune); ok {
		return int64(v)
	}
	if v, ok := value.(int64); ok {
		return v
	}
	if v, ok := value.(float32); ok {
		return int64(v)
	}
	if v, ok := value.(float64); ok {
		return int64(v)
	}

	return -1
}
