package lang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../types"
)

func get(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("unable to get" + url)
		return nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("unable to read request body:" + url)
		return nil
	}

	return body
}

func logLangs() {
	resp := get("https://github-trending-api.now.sh/languages")
	if resp == nil {
		panic("empty response")
	}

	var data types.Languages
	if err := json.Unmarshal(resp, &data); err != nil {
		panic(err)
	}

	for _, item := range data.All {
		fmt.Println(item.URLParam)
	}
}
