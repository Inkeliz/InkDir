package Directory

import (
	"os"
	"strings"
)

const (
	pathSeparator string = string(os.PathSeparator)

	files Filter = iota
	folders
	both
)

type (
	FileList []os.FileInfo
	Filter byte
)

func contentFrom(dir string) (*FileList, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}

	names, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}

	n := FileList(names)
	return &n, nil
}

func (filelist FileList) filterBy(filter Filter) ([]string, error){
	var names []string

	for _, fileInfo := range filelist {
		if filter == both || (filter & folders == folders && fileInfo.IsDir()) || (filter & files == files && !fileInfo.IsDir()) {
			names = append(names, fileInfo.Name())
		}
	}

	return names, nil
}

func Files(dirPath string) ([]string, error) {
	content, err := contentFrom(dirPath)
	if err != nil {
		return nil, err
	}
	return content.filterBy(files)
}

func Folders(dirPath string) ([]string, error) {
	content, err := contentFrom(dirPath)
	if err != nil {
		return nil, err
	}
	return content.filterBy(folders)
}

func All(dirPath string) ([]string, error) {
	content, err := contentFrom(dirPath)
	if err != nil {
		return nil, err
	}
	return content.filterBy(folders | files)
}

func Exists(dirPath string) error {
	_, exist := os.Stat(dirPath)
	return exist
}

func ClearDir(path string) string {
	return path[:strings.LastIndex(path, pathSeparator)]
}

func Local() (string, error) {
	return os.Executable()
}

// ----------------------------------------
// HELPER
// ----------------------------------------

func CreatePath(prefix, dir string) string {
	if strings.HasPrefix(dir, prefix) {
		return dir
	}
	return strings.TrimRight(prefix, pathSeparator) + pathSeparator + dir
}