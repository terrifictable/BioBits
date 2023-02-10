package discord

import "net/http"

type Discord struct {
	token string

	User Profile
}


func New(token string) *Discord {
	return &Discord{
		token: token,
		User: Profile{},
	}
}


func (dc Discord) GetHeaders() http.Header {
	return http.Header{
		"authorization": { dc.token },
	}
}
