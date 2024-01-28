package scenario

import (
	"path/filepath"
	"runtime"
)

func GetBaseTestAsset(elem ...string) string {
	_, filename, _, _ := runtime.Caller(0)
	basedir := filepath.Dir(filename)

	listPath := []string{basedir, "../test"}

	listPath = append(listPath, elem...)

	return filepath.Join(listPath...)
}
