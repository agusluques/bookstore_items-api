package items

import (
	"github.com/agusluques/bookstore_items-api/clients/elasticsearch"
	"github.com/agusluques/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
	typeItem   = "item"
)

func (i *Item) Save() *rest_errors.RestError {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", err)
	}
	i.ID = result.Id
	return nil
}
