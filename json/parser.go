package json

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseJsonFromUrl(url string) map[string]any {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var jsonObject map[string]any
	if err := json.Unmarshal(body, &jsonObject); err != nil {
		panic(err)
	}
	return jsonObject
}
