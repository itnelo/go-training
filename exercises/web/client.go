package web

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	Url string = "http://ip-api.com/csv"
)

func CallClient() {
	fmt.Println("Starting a webserver instance...")

	var client = &http.Client{}

	response, clientErr := client.Get(Url)
	if nil != clientErr {
		fmt.Println("Failed to get contents by URL.", clientErr)
		os.Exit(1)
	}

	var stringBuilder = &strings.Builder{}

	// ioutil.ReadAll is potential unsafe for big data
	_, copyErr := io.Copy(stringBuilder, response.Body)
	if nil != copyErr {
		fmt.Println("Failed to read response body.", copyErr)
		os.Exit(2)
	}

	var responseAsString = stringBuilder.String()

	fmt.Println(responseAsString)
}
