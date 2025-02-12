package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
)

// OpenDefaultFileManager
// launches the default graphical text editor
// запускает графический текстовый редактор по-умолчанию
func OpenDefaultFileManager(desc string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("os.Getwd: %v", err)
	}
	tmp := "tmp_file_" + strconv.Itoa(rand.Int())
	tmpFile, err := os.Create(wd + string(os.PathSeparator) + tmp)
	if err != nil {
		return "", fmt.Errorf("os.Create: %v", err)
	}

	_, err = tmpFile.WriteString(desc + "\n")
	if err != nil {
		return "", fmt.Errorf("tmpFile.WriteString: %v", err)
	}

	path, err := exec.LookPath("xdg-open") // searching a preferred text editor (ищем предпочтительный текстовый редактор)

	cmd := exec.Command(path, tmpFile.Name())
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("cmd.Run: %v\n", err)
	}

	return tmpFile.Name(), nil
}
