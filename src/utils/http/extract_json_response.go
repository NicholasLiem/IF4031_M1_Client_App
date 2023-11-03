package http

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func GetJSONDataBytesFromResponse(response *http.Response) ([]byte, error) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	var jsonResponse map[string]interface{}
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&jsonResponse); err != nil {
		return nil, err
	}

	data, ok := jsonResponse["data"].(map[string]interface{})
	if !ok {
		return nil, errors.New("data not found in the JSON response")
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return dataBytes, nil
}
