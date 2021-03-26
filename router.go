package main

import (
	"net/http"
)

func (h *HandlerStrust) Router() *http.ServeMux {

	router := http.NewServeMux()
	router.HandleFunc("/", h.HandleListCustomers)
	router.HandleFunc("/ee", h.HandleListCustomersEstonian)
	router.HandleFunc("/ru", h.HandleListCustomersRussian)
	router.HandleFunc("/add", h.HandleViewCustomer)
	router.HandleFunc("/add_ee", h.HandleViewCustomerEstonian)
	router.HandleFunc("/add_ru", h.HandleViewCustomerRussian)
	router.HandleFunc("/save", h.HandleSaveCustomer)
	router.HandleFunc("/save_ee", h.HandleSaveCustomerEstonian)
	router.HandleFunc("/save_ru", h.HandleSaveCustomerRussian)
	router.HandleFunc("/delete", h.HandleDeleteCustomer)
	router.HandleFunc("/delete_ee", h.HandleDeleteCustomerEstonian)
	router.HandleFunc("/delete_ru", h.HandleDeleteCustomerRussian)
	router.HandleFunc("/search", h.HandleSearchCustomer)
	router.HandleFunc("/search_ee", h.HandleSearchCustomerEstonian)
	router.HandleFunc("/search_ru", h.HandleSearchCustomerRussian)

	return router
}
