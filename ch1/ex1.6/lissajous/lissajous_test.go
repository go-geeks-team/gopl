package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

// Тест создаёт папку out(если она не создана) и генерирует файлы png
func TestLissajous(t *testing.T) {
	dir, err := os.Getwd() // get current working directory
	if err != nil && !os.IsExist(err) {
		t.Fatalf("os.Getwd: %s\n", err)
	}

	err = os.Mkdir(dir+"/out", 0750) // create output directory
	if err != nil && !os.IsExist(err) {
		t.Fatalf("os.Mkdir: %s\n", err)
	}

	err = os.Chdir("./out") // change directory
	if err != nil {
		t.Fatalf("os.Chdir: %s\n", err)
	}
	_ = os.RemoveAll("./") // delete all files in out directory

	created := createFiles() // create files: 1.png, 2.png, 3.png
	for {
		if created == false { // waiting until files is creating...
			return
		} else {

			s := os.DirFS("./")
			err = fs.WalkDir(s, ".", func(path string, d fs.DirEntry, err error) error {
				if err == nil {
					if d.Type().IsRegular() && filepath.Ext(d.Name()) == ".png" { // check is png file (not dir or other file)

						f, err := os.Open(d.Name())
						if err != nil {
							t.Fatalf("os.Open: %v\n", err)
						}
						fi, err := f.Stat()
						if err != nil {
							t.Fatalf("f.Stat: %v\n", err)
						}
						size := fi.Size() / 1000
						t.Logf("file %s is correct, size: %v kb\n", d.Name(), size)
					}
					return nil
				}
				return err
			})
			break
		}
	}
}
