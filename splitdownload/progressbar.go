package splitdownload

import (
	"os"
	"time"

	"github.com/cheggaaa/pb/v3"
)

func ProgressBar(t *Target, file_path *string, ch chan int) {
	file_size := t.FileSize
	now_size := int64(0)
	bar := pb.Start64(file_size)
	for {
		fi, err := os.Stat(*file_path)
		ChkErr(err)
		now_size = fi.Size()
		if now_size < file_size {
			bar.SetCurrent(now_size)
		} else {
			bar.SetCurrent(file_size)
			bar.Finish()
			ch <- 1
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}
