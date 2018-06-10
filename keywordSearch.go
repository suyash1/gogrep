/*

Go program to find a keyword recursively in all files present 
in directories and subdirectories for a given path

Usage:
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
go build keywordSearch.go
./keywordSearch -path <some_file_path> -search <search_keyword>

-or-

go run keywordSearch.go -path <some_file_path> -search <search_keyword>
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

*/
package main
 
import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)
 
var path = flag.String("path", "", "file path to search in")
var search = flag.String("search", "", "search string to look for")
 
func main() {
	flag.Parse()
	dir, err := os.Stat(*path)
	if err != nil {
		log.Fatal(err)
	}
	
	if dir.Mode().IsDir() {
		*path = strings.TrimRight(*path, "/") + "/"
		dirSearch(*path, *search)
	} else {
		log.Fatal("path must be a directory, but file was provided: ", *path)
	}
}
 
func dirSearch(path, search string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fullpath := path + file.Name()
		if file.Mode().IsDir() {
			dirSearch(fullpath+"/", search)
		} else if file.Mode().IsRegular() {
			fileSearch(fullpath, search)
		}
	}
}
 
func fileSearch(fullpath, search string) {
	data, err := ioutil.ReadFile(fullpath)
	if err != nil {
		log.Fatal(err)
	}
	//need to check for file type to detect filter off non-text files
	fileType := http.DetectContentType(data)
	if strings.Index(fileType, "text") == -1 {
		//skip all non text files
		return
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.Index(line, search) > -1 {
			fmt.Printf("%s: %s\n", fullpath, line)
		}
	}
}
