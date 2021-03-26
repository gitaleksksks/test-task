package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (h *HandlerStrust) ErrorOfServer(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	h.printError.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}

func (h *HandlerStrust) ErrorOfClient(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (h *HandlerStrust) NotFound(w http.ResponseWriter) {
	h.ErrorOfClient(w, http.StatusNotFound)
}
