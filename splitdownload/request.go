package splitdownload

import (
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
)

type Target struct {
	Url       string
	DirPath  string
	FileName string
	FileSize int64
	SplitTimes int64
}

func BuildRequest(i int, req *http.Request, size_array []int64, t *Target)  *http.Request{

	if i == 0 {
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", 0, size_array[i]))
		return req
	}

	if i == int(t.SplitTimes) - 1 {
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", size_array[i-1]+1, t.FileSize))
		return req
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", size_array[i-1]+1, size_array[i]))
	return req
}

func GetFileName(u string) string {
	result, err := url.Parse(u)
	ChkErr(err)
	return filepath.Base(result.Path)
}

func GetFilePath(d, f string) string {
	return filepath.Join(d, f)
}
