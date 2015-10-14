package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	oss "github.com/kyf/oss_sdk/lib"
	"net/http"
	"strings"
)

func init() {
	var group string = "sight"
	myHandlers[fmt.Sprintf("/%s", group)] = func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		data := r.Form
		file_content := data.Get("file")
		file_content = strings.Replace(file_content, "data:image/jpg;base64,", "", -1)

		oss.Init(OSS_ACCESS_ID, OSS_ACCESS_KEY, logger)
		path := generationPath(group, "jpg")
		img_content, err := base64.StdEncoding.DecodeString(file_content)
		if err != nil {
			logger(err)
		}
		oss.MkDir(BUCKET, fmt.Sprintf("/%s", group))
		oss.Create(BUCKET, path, string(img_content))

		logger(fmt.Sprintf("file [%s] upload success!", path))
		result := map[string]string{
			"status": "ok",
			"path":   path,
		}

		jsonResult, _ := json.Marshal(result)
		w.Write(jsonResult)
	}
}
