package lib

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	ENDPOINT   string = "oss-cn-beijing.aliyuncs.com"
	GMT_LAYOUT string = "Mon, 2 Jan 2006 15:04:05 GMT"

	METHOD_GET    string = "GET"
	METHOD_PUT    string = "PUT"
	METHOD_DELETE string = "DELETE"
	METHOD_POST   string = "POST"
)

type oss struct {
	access_id   string
	access_key  string
	urlString   string
	method      string
	bucket      string
	contentType string
	resource    string
	content     io.Reader
	req         *http.Request
	ossHeaders  map[string]string
}

func New(access_id, access_key string) *oss {
	return &oss{access_id: access_id, access_key: access_key, ossHeaders: make(map[string]string)}
}

func (o *oss) SetMethod(method string) {
	o.method = method
}

func (o *oss) SetContentType(contentType string) {
	o.contentType = contentType
}

func (o *oss) SetResource(resource string) {
	o.resource = resource
}

func (o *oss) SetContent(content string) {
	o.content = strings.NewReader(content)
}

func (o *oss) SetOSSHeader(key, value string) {
	o.ossHeaders[key] = value
}

func (o *oss) PrepReq() {
	req, err := http.NewRequest(o.method, o.urlString, o.content)
	location, err := time.LoadLocation("GMT")
	now := time.Now().In(location).Format(GMT_LAYOUT)
	signature := generationSign(o.method, now, o.bucket, o.resource)
	req.Header.Set("Date", now)
	req.Header.Set("Host", ENDPOINT)
	req.Header.Set("Authorization", fmt.Sprintf("OSS %s:%s", OSS_ACCESS_ID, signature))
	req.Header.Set("Content-Type", content_type)
	for k, v := range o.ossHeaders {
		req.Header.Set(strings.ToUpper(k), v)
	}
	o.req = req
}

func (o *oss) Do() (int, http.Header, []byte) {
	client := &http.Client{}
	res, err := client.Do(o.req)
	if err != nil {
		//log
	}
	return res.StatusCode, res.Header, res.Body
}

func generationSign(method, date, bucket, resource, access_key, content_type string, headers map[string]string) string {
	content_md5 := ""
	content_type := ""
	CanonicalizedOSSHeaders := ""
	for k, v := range headers {
		CanonicalizedOSSHeaders = fmt.Sprintf("%s%s:%s\n", CanonicalizedOSSHeaders, v, k)
	}
	CanonicalizedResource := fmt.Sprintf("%s%s", bucket, resource)

	sign := fmt.Sprintf("%s\n%s\n%s\n%s\n%s%s", method, content_md5, content_type, date, CanonicalizedOSSHeaders, CanonicalizedResource)

	hmacencoder := hmac.New(sha1.New, []byte(access_key))
	hmacencoder.Write([]byte(sign))

	signature := base64.StdEncoding.EncodeToString(hmacencoder.Sum(nil))
	return signature
}

var stdoss = New()

func MkDir(bucket, resource string) {
	stdoss.SetOSSHeader("x-oss-acl", "primate")
	stdoss.SetMethod(METHOD_PUT)
	stdoss.SetBucket(bucket)
	stdoss.SetContentType("")
	stdoss.SetContent("")
	stdoss.SetResource(fmt.Sprintf("%s/", resource))
	stdoss.PrepReq()
	status, header, body := stdoss.Do()
	fmt.Println(status, body, header)
}

func RmDir() {

}

func Create() {

}

func Remove() {

}
