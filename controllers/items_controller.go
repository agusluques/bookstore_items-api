package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/agusluques/bookstore_items-api/domain/items"
	"github.com/agusluques/bookstore_items-api/services"
	"github.com/agusluques/bookstore_items-api/utils/http_utils"
	"github.com/agusluques/bookstore_oauth-go/oauth"
	"github.com/agusluques/bookstore_utils-go/rest_errors"
)

var ItemsController itemsControllerInterface = &itemsController{}

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//http_utils.RespondError(w, err)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(reqBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = oauth.GetCallerId(r)

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		//http_utils.RespondJson(w, createErr.Status, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)

}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
