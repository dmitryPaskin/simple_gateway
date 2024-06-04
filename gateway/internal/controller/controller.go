package controller

import (
	"API/gateway/internal/service"
	"net/http"
)

type Controller struct {
	service.Service
}

func (c *Controller) GetMessageService1(w http.ResponseWriter, r *http.Request) {
	c.Service = service.NewService1()
	message, err := c.Service.GetResponse(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte(message)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Controller) GetMessageService2(w http.ResponseWriter, r *http.Request) {
	c.Service = service.NewService2()
	message, err := c.Service.GetResponse(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte(message)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}