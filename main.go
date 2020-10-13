package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	matches   []string
	waitGroup = sync.WaitGroup{}
	lock      = sync.Mutex{}
)

func fileSearch(root, filename string) {
	fmt.Println("Searching in root", root)
	files, _ := ioutil.ReadDir(root)
	//if err!=nil{
	//	panic(err)
	//}
	for _, file := range files {
		if strings.Contains(strings.ToUpper(file.Name()), strings.ToUpper(filename)) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, filename))
			lock.Unlock()
		}
		if file.IsDir() && !strings.Contains(strings.ToUpper(file.Name()), "NODE") {
			waitGroup.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), filename)
			//fileSearch(filepath.Join(root, file.Name()), filename)
		}
	}
	waitGroup.Done()
}

func main() {
	start := time.Now()
	waitGroup.Add(1)
	go fileSearch("/home/cassio.roos/Documentos", "README.md")
	//fileSearch("/home/cassio.roos/Documentos", "README.md")
	waitGroup.Wait()
	for _, file := range matches {
		fmt.Println("Matched", file)
	}

	fmt.Println("TOOK", time.Since(start))
	fmt.Println("FOUND", len(matches))
}
