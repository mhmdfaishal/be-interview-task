package virustotalservice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IguteChung/go-errors"
)

var (
	resultAnalyzePath = "/v3/analyses"
)

func GetResultService(xApiKey string, id string) (result map[string]interface{}, err error) {

	url := fmt.Sprintf("%s%s/%s", baseUrl, resultAnalyzePath, id)
	log.Println(url)
	var client = &http.Client{Timeout: time.Duration(defaultTimeout) * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-apikey", xApiKey)

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
