package splitdownload

import (
	"fmt"
	"net/http"
)

func Check(t *Target) error {
	fmt.Println("checking url...")
	res, err := http.Head(t.Url)

	ChkErr(err)
	if res.Header.Get("Accept-Ranges") != "bytes" {
		return fmt.Errorf("not supported range access.")
	}
	if res.ContentLength <= 0 {
		return fmt.Errorf("invalid content length.")
	}
	t.FileSize = int64(res.ContentLength)
	t.FileName = GetFileName(t.Url)
	fmt.Println("check ok.")

	return nil
}
