package splitdownload

import (
	"fmt"
	"os"
)

func ChkErr(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}