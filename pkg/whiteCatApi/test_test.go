package whiteCatApi

import (
	"testing"
)

func TestName(t *testing.T) {
	client := NewCatClient()
	client.SetToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3NDg1NjI0MDMsInJvbGUiOiJkeXAiLCJleHAiOjE3ODAwOTg0MDMsInN1YiI6NDk2MjgzfQ.UqVhDxEiapIyJNl3Y9fAN4SXLEpupM-Qw-Mn4UZrVnM")
	client.GetCityList()

}
