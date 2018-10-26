package client

import (
	"github.com/codingchipmunk/JolokiaGo/responses"
	"net/http"
)

type Client struct{
	url string
	client *http.Client
	sseID string
}

func (jc *Client) MakeRequest(request *interface{}) (resp responses.Root, err error){
	
}