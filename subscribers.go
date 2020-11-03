package aweber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// NewSubscriberData returns a NewSubscriberData struct without assigned data
// it is up to the user to assign the valuse they want included in the POST request.
// Only "Email" is required.
func NewSubscriberData(email, name string, tags []string) *SubscriberData {
	return new(SubscriberData)
}

// AddSubscriber sends an authorized POST request to add an the SubscriberData as a new subscriber.
func (a *Aweber) AddSubscriber(listName string, sub *SubscriberData) error {
	listID, err := a.GetListID(listName)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/lists/%d/subscribers", a.baseURL, listID)
	buf := new(bytes.Buffer)

	token, err := tokenFromFile("token.json")
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	subJSON, err := json.Marshal(sub)
	if err != nil {
		return err
	}

	req.Body = ioutil.NopCloser(bytes.NewReader(subJSON))

	resp, err := a.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
