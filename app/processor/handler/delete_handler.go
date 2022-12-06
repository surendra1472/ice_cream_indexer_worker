package handler

import (
	"fmt"
	"ic-indexer-service/models"
	"ic-indexer-service/client"
	"ic-indexer-service/client/operations"
)

//go:generate sh -c "$GOPATH/bin/mockery -case=underscore -dir=. -name=IcecreamDeleteHandler"
type IcecreamDeleteHandler interface {
	IcecreamDelete(models.IcecreamClientRequest) error
}


type icecreamDeleteHandler struct {
}

func NewIcecreamDeleteHandler() *icecreamDeleteHandler {
	return &icecreamDeleteHandler{}
}


func(icdh icecreamDeleteHandler) IcecreamDelete(request models.IcecreamClientRequest) error {
	param := operations.NewDeleteIcecreamParams()
	param.ProductID = request.ProductID


	client := client.NewHTTPClientWithConfig(nil, client.DefaultTransportConfig().WithHost("localhost:5030"))

	client.Operations.SetTransport(client.Transport)
	client.SetTransport(client.Transport)

	icecream, err := client.Operations.DeleteIcecream(param)
	if err == nil {
		fmt.Print(icecream.Payload.Data)
	} else {
		data, ok := err.(*operations.DeleteIcecreamOK)
		if ok {
			fmt.Print(data.Payload.Data)
		}
	}

	return err

}