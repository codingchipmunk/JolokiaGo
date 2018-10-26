package client

import "net/http"

type Client struct{
	url string
	client *http.Client
	sseID string
}

func (jc *Client) MakeRequest(request *interface{}) (ResponseRoot, error){

}