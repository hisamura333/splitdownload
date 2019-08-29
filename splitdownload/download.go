package splitdownload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func Download(t *Target) {
	fmt.Println("downloading...")

	out_path := GetFilePath(t)
	out, err := os.Create(out_path)
	ChkErr(err)
	defer func() {
		cerr := out.Close()
		if cerr == nil {
			return
		}
		err = fmt.Errorf("Failed to close: %v, the original error was %v", cerr, err)
		ChkErr(err)
	}()

	ch := make(chan int)
	go ProgressBar(t, &out_path, ch)

	// 1回のDLで行うファイルサイズ
	split_size := t.FileSize / t.SplitTimes

	// 並列で回すときに、何byteまでDLするかの値listに保持する
	size_array := []int64{}

	var times int64
	for times = 1; times < t.SplitTimes+1; times++ {
		size_array = append(size_array, split_size*times)
	}
	var downloadFiles int

	for i := range size_array {

		i := i
		out := out

		go func() {
			req, err := http.NewRequest("GET", t.Url, nil)
			ChkErr(err)

			buildReq := BuildRequest(i, req, size_array, t)

			res, err := http.DefaultClient.Do(buildReq)
			ChkErr(err)

			defer res.Body.Close()

			for {
				if downloadFiles == i {
					v, err := io.Copy(out, res.Body)
					ChkErr(err)

					// ゴルーチンで処理したファイルのサイズが分割した値の半分よりも小さい時はエラー
					// 数byteしか処理ができない時があり、分割DLが完了しない時があったので、
					// それを防ぐために一時的に追加
					if v < split_size/2 {
						fmt.Println("split_size error")
						os.Exit(1)
					}
					downloadFiles++
				}
				time.Sleep(100 * time.Millisecond)
			}
		}()
	}
	<-ch
}