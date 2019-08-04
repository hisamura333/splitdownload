package main

import (
	"flag"
	"os"
	"fmt"

	"github.com/hisamura333/splitdownload/splitdownload"
)

func main() {
	var (
		times = flag.Int64("t", 5, "並列に処理する数")
		dir = flag.String("d", ".", "ダウンロードされるディレクトリの指定")
	)

	flag.Parse()
	var targetUrl string

	if len(os.Args) < 2 {
		fmt.Println("ダウンロードするファイルを指定してください")
		os.Exit(1)
	}

	targetUrl = os.Args[1]
	target := splitdownload.Target{
		Url:        targetUrl,
		DirPath:   *dir,
		SplitTimes: *times,
	}

	err := splitdownload.Check(&target)
	splitdownload.ChkErr(err)

	splitdownload.Download(&target)
}

