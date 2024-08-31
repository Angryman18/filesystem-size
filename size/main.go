package main

import (
	"fmt"
	"size/core"
)

func main() {
	coreArgs := core.OsArgs{Folder: "", Short: false, Delete: ""}
	cmdArgs := coreArgs.GetArgs()
	operation := core.NewOperation(cmdArgs)
	operation.Operate()
	// wg := &sync.WaitGroup{}
	// dChan := make(chan string)

	// go createFold(dChan)
	// for i := 0; i <= 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		for dir := range dChan {
	// 			fmt.Println(dir, i)
	// 			os.Mkdir(dir, 0777)
	// 		}
	// 	}()
	// }

	// wg.Wait()

}

func createFold(dChan chan<- string) {
	defer close(dChan)
	for i := range Xrange(1000) {
		parentDir := fmt.Sprintf("test1/test %d", i)
		dChan <- parentDir
		for l := range Xrange(1000) {
			childDir := parentDir + "/" + fmt.Sprintf("test %d", l)
			dChan <- childDir

		}
	}
}

func Xrange(length int) chan int {
	dChan := make(chan int)
	go func() {
		defer close(dChan)
		loop := func() {
			for i := 0; i <= length; i++ {
				dChan <- i
			}
		}
		loop()
	}()
	return dChan
}
