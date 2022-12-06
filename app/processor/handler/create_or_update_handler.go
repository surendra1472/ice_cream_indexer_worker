package handler

import (
	"fmt"
	"ic-indexer-service/client"
	"ic-indexer-service/client/operations"
	"ic-indexer-service/models"
)

//go:generate sh -c "$GOPATH/bin/mockery -case=underscore -dir=. -name=IcecreamCreateOrUpdateHandler"
type IcecreamCreateOrUpdateHandler interface {
	IcecreamCreateOrUpdate(models.IcecreamClientRequest) error
}

type icecreamCreateOrUpdateHandler struct {
}

func NewIcecreamCreateOrUpdateHandler() *icecreamCreateOrUpdateHandler {
	return &icecreamCreateOrUpdateHandler{}
}


func(icuh icecreamCreateOrUpdateHandler) IcecreamCreateOrUpdate(request models.IcecreamClientRequest) error {

	param := operations.NewSaveOrUpdateIcecreamParams()
	param.IcecreamClientRequest = &request


	client := client.NewHTTPClientWithConfig(nil, client.DefaultTransportConfig().WithHost("localhost:5030"))

	client.Operations.SetTransport(client.Transport)
	client.SetTransport(client.Transport)

	icecream, err := client.Operations.SaveOrUpdateIcecream(param)
	if err == nil {
		fmt.Print(icecream.Payload.Data)
	} else {
		data, ok := err.(*operations.SaveOrUpdateIcecreamDefault)
		if ok {
			fmt.Print(data.Payload.Data)
		}
	}

	return err
}