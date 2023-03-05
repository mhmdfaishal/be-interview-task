package virustotalservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/IguteChung/go-errors"
	"github.com/spf13/viper"
)

var (
	uploadFilePath = "/v3/files"
	defaultTimeout = 20
	baseUrl        = viper.GetString("VIRUSTOTAL_BASE_URL")
)

func UploadFileService(xApiKey string, file *multipart.FileHeader) (result map[string]interface{}, err error) {

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	part, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		return nil, err
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, err
	}

	_, err = part.Write(fileBytes)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s", baseUrl, uploadFilePath)

	var client = &http.Client{Timeout: time.Duration(defaultTimeout) * time.Second}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-apikey", xApiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&result)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("VirusTotal API Error")
	}

	return result, nil
}
