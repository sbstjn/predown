package lib

import (
	"io/ioutil"
	"path"
	"path/filepath"
)

func filesInDirectory(dir string, root string) ([]string, error) {
	var list []string

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		absPath := path.Join(dir, file.Name())
		relPath, err := filepath.Rel(root, absPath)

		if err != nil {
			return nil, err
		}

		if file.IsDir() {
			children, err := filesInDirectory(absPath, root)

			if err != nil {
				return nil, err
			}

			list = append(list, children...)
		} else {

			list = append(list, relPath)
		}
	}

	return list, nil
}

// Files returns all files in template prefix
func Files(path string) ([]string, error) {
	var files []string
	var err error

	if files, err = filesInDirectory(path, path); err != nil {
		return nil, err
	}

	return files, nil
}
