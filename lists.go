package aweber

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// GetLists returns the list data from the Aweber lists API
func (a *Aweber) GetLists() *Lists {
	client := a.Client
	lists := fmt.Sprintf("%slists", a.baseURL)

	resp, err := client.Get(lists)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	allLists := new(Lists)
	if resp.StatusCode == http.StatusOK {
		json.NewDecoder(resp.Body).Decode(allLists)
	}

	return allLists
}

// GetListID returns the list ID of the requested list by name
func (a *Aweber) GetListID(listName string) (int, error) {
	lists := a.GetLists()
	for _, list := range lists.Entries {
		if list.Name == listName {
			return list.ID, nil
		}
	}

	return 0, errors.New("requested list not found")
}
