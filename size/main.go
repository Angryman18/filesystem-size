package main

import (
	"fmt"
	"os"
)

type FileData struct {
	FileName string
	FileSize float64
}

type FileInfo struct {
	TreeChan   chan FileData
	DetailTree bool
}

type FileSize struct {
	Size float64
}

func (f *FileData) getSizeStr() string {
	size := f.FileSize
	switch {
	case size > 1024*1024*1024:
		return fmt.Sprintf("%.2f %s", size/float64(1024*1024*1024), "GB")
	case size > 1024*1024:
		return fmt.Sprintf("%.2f %s", size/float64(1024*1024), "MB")
	case size > 1024:
		return fmt.Sprintf("%.2f %s", size/float64(1024), "KB")
	case size > 0:
		return fmt.Sprintf("%.0f %s", size, "B")
	default:
		return "N/A"
	}
}

func main() {

	fileInfo := FileInfo{TreeChan: make(chan FileData)}
	totalSize := FileData{FileSize: 0.00}

	go func() {
		defer close(fileInfo.TreeChan)
		fileInfo.traverseFileSystem(".")
	}()

	for ele := range fileInfo.TreeChan {
		buildString := ele.getSizeStr() + " ---> " + ele.FileName
		fmt.Println(buildString)
		totalSize.FileSize += ele.FileSize
	}

	fmt.Println("Total Size = ", totalSize.getSizeStr())
}

func (f *FileInfo) traverseFileSystem(path string) {
	d, _ := os.ReadDir(path)
	var root string = path
	for _, ele := range d {
		if ele.IsDir() {
			f.traverseFileSystem(root + "/" + ele.Name())
		} else {
			fileInfo, err := ele.Info()
			if err != nil {
				panic("Error Reading File Info")
			}
			d := FileData{FileName: fileInfo.Name(), FileSize: float64(fileInfo.Size())}
			f.TreeChan <- d
		}
	}
}
