package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func createReqBody(filePath string) (string, io.Reader, error) {
	var err error

	buf := new(bytes.Buffer)
	bw := multipart.NewWriter(buf)

	f, err := os.Open(filePath)
	if err != nil {
		return "", nil, err
	}
	defer f.Close()

	// file part
	_, fileName := filepath.Split(filePath)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", fileName))
	h.Set("Content-Type", "image/png")
	formWriter, _ := bw.CreatePart(h)
	io.Copy(formWriter, f)
	bw.Close()

	return bw.FormDataContentType(), buf, nil
}

func doUpload(url, filePath string) error {
	// create body
	contType, reader, err := createReqBody(filePath)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return err
	}

	fmt.Println("file content type: ", contType)
	req.Header.Add("Content-Type", contType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed, err: %v\n", err)
		return err
	}

	fmt.Println(string(b))

	return nil
}

func uploadFile(wg *sync.WaitGroup, filePath string) {
	defer wg.Done()

	url := "http://127.0.0.1:8888/upload"
	err := doUpload(url, filePath)
	if err != nil {
		fmt.Println("upload file failed, error: ", err)
	}
}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 1; i++ {
		wg.Add(1)
		go uploadFile(&wg, "test"+strconv.Itoa(i)+".png")
	}

	wg.Wait()
}
