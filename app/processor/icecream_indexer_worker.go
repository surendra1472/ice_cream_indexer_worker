package processor

import (
	"encoding/json"
	"ic-indexer-service/models"
	"ic-indexer-worker/app/processor/handler"
)

//go:generate sh -c "$GOPATH/bin/mockery -case=underscore -dir=. -name=IcecreamIndexerWorker"
type IcecreamIndexerWorker interface {
	ProcessIcecreamStreamData([]byte) (map[string]interface{}, error)
}

type icecreamIndexerWorker struct {
	icecreamCreateHandler handler.IcecreamCreateOrUpdateHandler
	icecreamDeleteHandler handler.IcecreamDeleteHandler
}

func GetIcecreamIndexerWorker() *icecreamIndexerWorker {
	return &icecreamIndexerWorker{icecreamDeleteHandler: handler.NewIcecreamDeleteHandler(),icecreamCreateHandler: handler.NewIcecreamCreateOrUpdateHandler()}
}

func (ls icecreamIndexerWorker) ProcessIcecreamStreamData(data []byte) error {

	indexRequest := models.IcecreamClientRequest{}
	err := json.Unmarshal(data, &indexRequest)

	if err == nil {
		if indexRequest.IsDeleted {
			return ls.deleteIcecreamToIndexer(indexRequest)
		}else {
			return ls.sendIcecreamToIndexer(indexRequest)
		}
	}

	return err
}



func (ls icecreamIndexerWorker) sendIcecreamToIndexer(request models.IcecreamClientRequest) error {
	return ls.icecreamCreateHandler.IcecreamCreateOrUpdate(request)
}


func (ls icecreamIndexerWorker) deleteIcecreamToIndexer(request models.IcecreamClientRequest) error {

	return ls.icecreamDeleteHandler.IcecreamDelete(request)
}
