package requests

type POSTRequest interface {
	POSTBody() ([]byte,error)
}