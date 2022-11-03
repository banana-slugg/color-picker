package art_test

import (
	"fmt"
	"testing"

	"github.com/crims1n/color-picker/internal/util"
)

func TestFilesFromDir(t *testing.T) {
	files := util.FilesFromDir("./testdata")

	for _, v := range files {
		fmt.Println(v)
	}
}
