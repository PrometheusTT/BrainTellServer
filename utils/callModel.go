package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func CallPredictService(filePath string) (string, error) {
	savePath := "/home/BrainTellServer/tmppath/" + fmt.Sprintf("%s_%d.v3draw", "callModel", time.Now().UnixNano()) //转换后输出地址
	//savePath := "D://BrainTellServer//tmppath//" + fmt.Sprintf("%s_%d.v3draw", "callModel", time.Now().UnixNano()) //转换后输出地址
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 准备表单数据
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}
	writer.Close()

	// 发送POST请求
	req, err := http.NewRequest("POST", "http://114.117.165.134:26011/predict", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	// 创建输出文件
	outFile, err := os.Create(savePath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	// 将响应写入文件
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return "", err
	}

	return savePath, nil
}
