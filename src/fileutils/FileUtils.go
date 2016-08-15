package fileutils

import (
	."code"
	"io/ioutil"
	"fmt"
	"path/filepath"
)

func LoadPage(file string) (Page, error){
	var page Page
	path, err := GetAbsPath(file)
	if err != nil {
		fmt.Print(err)
		return page, err
	}
	
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
		return page, err
	}
	page.Body = bytes
	return page, err
}

func GetAbsPath(path string) (string,error) {
	return filepath.Abs(path)
}