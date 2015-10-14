package main

import (
	"fmt"
	oss "github.com/kyf/oss_sdk/lib"
	"net/http"
)

func init() {
	var group string = "demo"
	myHandlers[fmt.Sprintf("/%s", group)] = func(w http.ResponseWriter, r *http.Request) {
		oss.Init(OSS_ACCESS_ID, OSS_ACCESS_KEY, logger)
		//oss.RmDir(BUCKET, group)
		oss.Remove(BUCKET, "/sight/2015_10_14_14447953997024873811298498081.jpg")

		w.Write([]byte("done"))
	}
}
