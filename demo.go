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
		/*
			oss.RmDir(BUCKET, "/sight")
			oss.Remove(BUCKET, "/sight/2015_10_14_14447985047668171981298498081.jpg")
			oss.RemoveBucket(BUCKET)
		*/
		//oss.CreateBucket(BUCKET)
		w.Write([]byte("done"))
	}
}
