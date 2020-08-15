package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fileUrl := "https://www.federalreserve.gov/datadownload/Output.aspx?rel=sloos&filetype=zip"
	err := DownloadFile("sloos.zip", fileUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + fileUrl)
}

func DownloadFile(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}