package model

import (
	"net/http"
	"time"
)

type (
	//The Book Model
	Book struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	//The Route Model
	Route struct {
		Method      string
		Pattern     string
		HandlerFunc http.HandlerFunc
	}

	//The Pagination Model
	Pagination struct {
		Limit        int `json:"limit"`
		Offset       int `json:"offset"`
		TotalRecords int `json:"totalRecords"`
	}

	//The error model
	ErrorBody struct {
		Error  string `json:"error"`
		Status int    `json:"status"`
	}

	//The item model
	Item struct {
		Value interface{} `json:"value"`
	}

	//PaginatedResponse model
	PaginatedResponse struct {
		Data       interface{} `json:"data,omitempty"`
		Pagination *Pagination `json:"pagination,omitempty"`
	}

	//The responseBody model
	ResponseBody struct {
		Data *[]Item `json:"data, omitempty"`
	}

	//The responseSingleBody model
	ResponseSingleBody struct {
		Data *Item `json:"data,omitempty"`
	}
)
