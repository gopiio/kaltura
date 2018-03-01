package kaltura

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Kaltura - Base type
type Kaltura struct {
	Secret     string
	PartnerID  string
	ServiceURL string
	Format     int
	Session    *Session
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//DoRequest - Handle the kaltura HTTP Request
func (k *Kaltura) DoRequest(method, resource, action string, payload map[string]interface{}) string {

	httpClient := &http.Client{}
	requestURL := k.ServiceURL + "/" + resource + "/action/" + action

	if payload == nil {
		payload = make(map[string]interface{})
	}
	payload["ks"] = k.Session.Value
	payload["format"] = k.Format
	payloadData, _ := json.Marshal(payload)

	request, err := http.NewRequest(method, requestURL, bytes.NewBuffer(payloadData))
	handleError(err)
	request.Header.Add("Content-Type", "application/json")

	response, err := httpClient.Do(request)
	handleError(err)

	body, err := ioutil.ReadAll(response.Body)
	handleError(err)

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, body, "", " ")
	return string(prettyJSON.Bytes())
}
