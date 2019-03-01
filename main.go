package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"
)

var uploadTemplate = template.Must(template.ParseFiles("index.html"))

func indexHandle(w http.ResponseWriter, r *http.Request) {
	if err := uploadTemplate.Execute(w, nil); err != nil {
		log.Fatal("Execute: ", err.Error())
		return
	}
}

func uploadHandle(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Fatal("FormFile: ", err.Error())
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal("Close: ", err.Error())
			return
		}
	}()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("ReadAll: ", err.Error())
		return
	}
	fmt.Println("file:::", string(bytes))
	w.Write(bytes)
}

func main() {
	//upload.FileUpload("*/.html")
	//http.HandleFunc("/", indexHandle)
	//http.HandleFunc("/upload", uploadHandle)
	//http.ListenAndServe(":8099", nil)
	str := "恭喜 发财"
	str1 := strings.Trim(str, " ")
	fmt.Println("str:", str1)
	//str1 := "hello word"
	count := len(str)
	fmt.Println("count:", count)
	num := utf8.RuneCountInString(str1)
	fmt.Println("num:", num)
	var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")
	fmt.Println(hzRegexp.MatchString(str))

	fmt.Println("Hello, 世界", len("世界"), utf8.RuneCountInString("世界"))
}
