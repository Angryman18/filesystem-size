package core

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Operation struct {
	Args *OsArgs
}

type FileData struct {
	FileName string
	FileSize float64
	IsDir    bool
}

type FileInfo struct {
	TreeChan   chan FileData
	DetailTree bool
}

type Result struct {
	FileSize   float64
	TotalFiles int
	TotalDirs  int
}

func (o *Operation) Operate() {

	if o.Args.Version {
		fmt.Println("gtds version 1.0.0")
		return
	}

	fileInfo := FileInfo{TreeChan: make(chan FileData)}
	totalSize := Result{FileSize: 0.00, TotalFiles: 0}

	path := ArgsSetup(o)
	start := time.Now().UnixNano() / int64(time.Millisecond)

	go func() {
		defer close(fileInfo.TreeChan)
		fileInfo.traverseFileSystem(path)
	}()

	for ele := range fileInfo.TreeChan {
		if !ele.IsDir {
			totalSize.TotalFiles += 1
			totalSize.FileSize += ele.FileSize
			if o.Args.Short {
				fmt.Printf("\r%d Files %d Folders Took %.2f s Total Size = %s", totalSize.TotalFiles, totalSize.TotalDirs, Duration(start), getSizeStr(totalSize.FileSize))
			} else {
				buildString := getSizeStr(ele.FileSize) + " ---> " + ele.FileName
				fmt.Println(buildString)
			}
		} else {
			totalSize.TotalDirs += 1
		}
	}

	// start.
	if !o.Args.Short {
		PrintResult(Duration(start), totalSize)
	} else {
		fmt.Println("")
	}
}

func PrintResult(duration float32, totalSize Result) {
	fmt.Println("\nTotal Size = ", getSizeStr(totalSize.FileSize))
	fmt.Printf("%d Files %d Folders ", totalSize.TotalFiles, totalSize.TotalDirs)
	fmt.Println("Took ", duration, "s")
}

func ArgsSetup(o *Operation) string {
	var path string = "."
	if o.Args.Help {
		fmt.Println(`
Usage: size [OPTIONS] [OPTIONS]
-s			- will show results in short
-f=[FOLDER_NAME]	- size will run on this specefic folder
-s -f=[FOLDER_NAME]	- size will run short on this specefic folder.
		`)
		os.Exit(0)
	}
	if len(o.Args.Folder) > 0 {
		path = o.Args.Folder
	}
	if len(o.Args.Delete) > 0 {
		var sure string
		fmt.Printf("Are you sure want to delete? [yes/no] ")
		fmt.Scan(&sure)
		if strings.ToLower(sure) != "yes" {
			fmt.Println("Exited")
			os.Exit(0)
		}
		err := os.RemoveAll(o.Args.Delete)
		if err != nil {
			fmt.Println("Could not remove file")
			os.Exit(1)
		}
		fmt.Println("Deleted successfully")
		os.Exit(0)
	}
	return path
}

func (f *FileInfo) traverseFileSystem(path string) {
	d, pathError := os.ReadDir(path)
	if pathError != nil {
		fmt.Println(" -> ", pathError.Error())
	}
	var root string = path
	for _, ele := range d {
		if ele.IsDir() {
			f.TreeChan <- FileData{IsDir: true}
			f.traverseFileSystem(root + "/" + ele.Name())
		} else {
			fileInfo, err := ele.Info()
			if err != nil {
				panic("Error Reading File Info")
			}
			d := FileData{FileName: fileInfo.Name(), FileSize: float64(fileInfo.Size()), IsDir: false}
			f.TreeChan <- d
		}
	}
}

func Duration(start int64) float32 {
	remain := (time.Now().UnixNano() / int64(time.Millisecond)) - start
	return float32(remain) / float32(1000)
}

func getSizeStr(filesize float64) string {
	size := filesize
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

func NewOperation(osArgs *OsArgs) *Operation {
	return &Operation{Args: osArgs}
}
