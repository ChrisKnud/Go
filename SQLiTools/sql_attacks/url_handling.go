package sql_attacks

import (
	"fmt"
	"io"
	"net/http"
)

// Get url request and body as string
func Get_url_req(url string) (*http.Response, string, error) {

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error: ", err)
		return nil, "Request failed: no body", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return resp, string(body), nil
}
