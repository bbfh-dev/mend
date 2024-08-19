package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Error(err error, during string) {
	var internal string
	if err != nil {
		internal = fmt.Sprintf(": %s", err.Error())
	}
	fmt.Printf("ERROR: %s%s\n", during, internal)
	os.Exit(1)
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadDir(path ...string) []os.DirEntry {
	entries, err := os.ReadDir(filepath.Join(path...))
	if err != nil {
		Error(err, "Reading dir")
	}

	return entries
}

func FileReader(path ...string) *os.File {
	file, err := os.OpenFile(filepath.Join(path...), os.O_RDONLY, 0755)
	if err != nil {
		Error(err, "Reading file")
	}

	return file
}

func ReadFile(path ...string) string {
	data, err := os.ReadFile(filepath.Join(path...))
	if err != nil {
		Error(err, "Reading file")
	}

	return string(data)
}

func WriteFile(path string, content string) {
	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			Error(err, "Making dirs")
		}
	}

	err := os.WriteFile(path, []byte(content), os.ModePerm)
	if err != nil {
		Error(err, "Writing file")
	}
}

func Base(entry os.DirEntry) string {
	name := entry.Name()
	return strings.TrimSuffix(name, filepath.Ext(name))
}
