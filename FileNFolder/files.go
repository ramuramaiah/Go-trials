package FileNFolder

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func FileOpenTest() {
	file, err := os.Open("C:/Work/test.txt")
	if err != nil {
		// handle the error here
		panic(err)
	}

	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		panic(err)
	}

	str := string(bs)
	fmt.Println(str)
}

func FileOpenUsingIOUtil() {
	bs, err := ioutil.ReadFile("C:/Work/test.txt")
	if err != nil {
		panic(err)
	}

	str := string(bs)
	fmt.Println(str)
}

func ListFilesInDir() {
	dir, err := os.Open("C:/Work/")
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() &&
		(info.Name() == "xx" || info.Name() == "yy") {
		return filepath.SkipDir
	}

	fmt.Println(path, info.Size())
	return err
}

func ListFilesInDirRecurse() {
	err := filepath.Walk("C:/Work/", walkFunc)

	if err != nil {
		log.Println(err)
	}
}
