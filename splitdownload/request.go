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

//ゴルーチン内でアクセスするためのリクエストを作成する
// rangeの値にはfromとtoの値を入力する必要がある
// 最初と最後とそれ以外でbytesに設定する値の設定を変えている

func BuildRequest(i int, req *http.Request, size_array []int64, t *Target)  *http.Request{

	// 1回目の時のパターン
	if i == 0 {
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", 0, size_array[i]))
		return req
	}

	// 最後の時のパターン
	if i == int(t.SplitTimes) - 1 {
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", size_array[i-1]+1, t.FileSize))
		return req
	}

	// それ以外の時のパターン
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", size_array[i-1]+1, size_array[i]))
	return req
}

func GetFileName(u string) string {
	result, err := url.Parse(u)
	ChkErr(err)
	return filepath.Base(result.Path)
}

func GetFilePath(t *Target) string {
	return filepath.Join(t.DirPath, t.FileName)
}
