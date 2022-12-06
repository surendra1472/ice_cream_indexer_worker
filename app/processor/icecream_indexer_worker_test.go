package processor

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ic-indexer-worker/app/processor/handler/mocks"
	"testing"
)

type IcecreamClientRequest struct {

	// allergy info
	AllergyInfo string `json:"allergy_info,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// dietary certifications
	DietaryCertifications string `json:"dietary_certifications,omitempty"`

	// an icecream request
	// in: body
	ID int64 `json:"id,omitempty"`

	// image closed
	ImageClosed string `json:"image_closed,omitempty"`

	// image open
	ImageOpen string `json:"image_open,omitempty"`

	// ingredients
	Ingredients []string `json:"ingredients"`

	// is deleted
	IsDeleted bool `json:"is_Deleted,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// product Id
	ProductID string `json:"product_id,omitempty"`

	// sourcing values
	SourcingValues []string `json:"sourcing_values"`

	// story
	Story string `json:"story,omitempty"`
}



func getIcecreamClientRequest() IcecreamClientRequest {

	icecreamIndexRequest := IcecreamClientRequest{}
	icecreamIndexRequest.SourcingValues = []string{"1234"}
	icecreamIndexRequest.Story = "atlas"
	icecreamIndexRequest.ProductID = "2"
	icecreamIndexRequest.Name = "test"
	icecreamIndexRequest.Ingredients = []string{"I1"}

	return icecreamIndexRequest
}

func getIcecreamDeleteClientRequest() IcecreamClientRequest {

	icecreamClientRequest := IcecreamClientRequest{}
	icecreamClientRequest.SourcingValues = []string{"1234"}
	icecreamClientRequest.Story = "atlas"
	icecreamClientRequest.ProductID = "2"
	icecreamClientRequest.Name = "test"
	icecreamClientRequest.Ingredients = []string{"I1"}
	icecreamClientRequest.IsDeleted = true

	return icecreamClientRequest
}



func TestCreateIcecream(t *testing.T) {

	createMock := mocks.IcecreamCreateOrUpdateHandler{}

	b, _ := json.Marshal(getIcecreamClientRequest())

	worker := GetIcecreamIndexerWorker()

	worker.icecreamCreateHandler = &createMock

	createMock.On("IcecreamCreateOrUpdate", mock.Anything).Return(nil)

	err := worker.ProcessIcecreamStreamData(b)


	assert.Nil(t, err)

}

func TestDeleteIcecream(t *testing.T) {

	deleteMock := mocks.IcecreamDeleteHandler{}

	b, _ := json.Marshal(getIcecreamDeleteClientRequest())

	worker := GetIcecreamIndexerWorker()

	worker.icecreamDeleteHandler = &deleteMock

	deleteMock.On("IcecreamDelete", mock.Anything).Return(nil)

	err := worker.ProcessIcecreamStreamData(b)


	assert.Nil(t, err)

}


func TestJsonUnMarshallError(t *testing.T) {


	worker := GetIcecreamIndexerWorker()

	err := worker.ProcessIcecreamStreamData([]byte("yolo"))


	assert.NotNil(t, err)

}
