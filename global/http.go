package global

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(url string, headers map[string]string) ([]byte, error) {
	client := http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	if headers != nil {
		for key, val := range headers {
			reqest.Header.Set(key, val)
		}
	}

	response, err := client.Do(reqest)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	respByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return respByte, nil
}
