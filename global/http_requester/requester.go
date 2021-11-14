package httprequester

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func RequestCreator(method string, url string, body io.Reader, headers []map[string]string) (*http.Request, error) {
	//headers is nothing but a slice of map[string]string,
	//each element of this slice will contain a key value pair, which is namely, header
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	// we'll add headers to the request created.
	for _, header := range headers {
		if len(header) != 1 {
			fmt.Println("problem.")
			return nil, errors.New("header object you insert is empty")
		}
		req.Header.Add("Authorization", header["Authorization"])
	}

	return req, nil
}

func CreateReqestAndDo(method string, url string, body io.Reader, headers []map[string]string) (*http.Response, error) {
	client := http.Client{}
	httpRequest, err := RequestCreator(method, url, body, headers)
	if err != nil {
		return &http.Response{}, nil
	}
	resp, err2 := client.Do(httpRequest)
	if err2 != nil {
		return &http.Response{}, nil
	}
	return resp, nil

}
