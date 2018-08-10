package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
 *
 */

type GraqhQL string

/**
 *
 */
func GraphQLPost(graqhql GraqhQL) {
	url := "http://hk.prpr.io:8080/v1alpha1/graphql"

	var jsonStr = []byte(graqhql)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Hasura-Access-Key", "shenchaozj")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
