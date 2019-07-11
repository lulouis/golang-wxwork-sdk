package utils

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"

	"mime/multipart"

	"github.com/tealeg/xlsx"
)

// 网络请求GET
func HttpGet(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		fmt.Printf("解析url错误:\r\n%v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlstr := Url.String()
	resp, err := http.Get(urlstr)
	fmt.Println(urlstr)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// 网络请求POST body
func HttpPostBody(apiURL string, params url.Values, body string) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		fmt.Printf("解析url错误:\r\n%v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlstr := Url.String()
	resp, err := http.Post(urlstr, "application/json", strings.NewReader(body))
	fmt.Println(urlstr)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	rs, err = ioutil.ReadAll(resp.Body)
	return

}

// 网络请求POST multipart（二进制上传）
func HttpPostFile(filename string, apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		fmt.Printf("解析url错误:\r\n%v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlstr := Url.String()

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("pdf", filepath.Base(filename))
	if err != nil {
		return
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", urlstr, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := http.Client{}
	resp, err := client.Do(req)
	//resp, err := http.Post(urlstr, "multipart/form-data", body)
	fmt.Println(urlstr)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	rs, err = ioutil.ReadAll(resp.Body)
	return

}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//网络文件下载
func DownloadFile(fileName string, url string) (err error) {
	// Create the file
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

//移除重复数据
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	sort.Strings(arr)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// 获取文件大小的接口
type Size interface {
	Size() int64
}

func XlsxFileReader(mimeFile multipart.File) (*xlsx.File, error) {

	defer mimeFile.Close()
	var size int64
	if sizeInterface, ok := mimeFile.(Size); ok {
		size = sizeInterface.Size()
	}

	xlFile, err := xlsx.OpenReaderAt(mimeFile, size)
	return xlFile, err
}

// get 网络请求
func Get(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		fmt.Printf("解析url错误:\r\n%v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlstr := Url.String()
	resp, err := http.Get(urlstr)
	fmt.Println(urlstr)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
