package discord

import (
	"main/requests"
	"encoding/json"
	"fmt"
)

func (dc *Discord) GetUser(id string) error {
	if dc.token == "" {
		return fmt.Errorf("no token provided")
	}

	url := BaseURL + "/" + ApiVersions["9"] + "/users/" + id
	resp, err := requests.GetWithHeaders(url, dc.GetHeaders())
	if err != nil {
		return err
	}

	var result User
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return err
	}

	dc.User = result
	return nil
}
