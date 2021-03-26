package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func (h *HandlerStrust) HandleSaveCustomer(w http.ResponseWriter, r *http.Request) {

	var id = 0
	var err error
	r.ParseForm()
	params := r.PostForm
	idStr := params.Get("id")

	if len(idStr) > 0 {
		id, err = strconv.Atoi(idStr)
		if err != nil {
			h.ErrorPage(w, err)
			return
		}
	}

	firstname := strings.Title(strings.ToLower(params.Get("firstname")))
	lastname := strings.Title(strings.ToLower(params.Get("lastname")))
	gender := params.Get("gender")
	email := params.Get("email")
	homeaddress := params.Get("homeaddress")

	BirthDateStr := params.Get("birthdate")
	var birthdate time.Time

	if len(BirthDateStr) > 0 {
		birthdate, err = time.Parse("02 January 2006", BirthDateStr)
		if err != nil {
			h.ErrorPage(w, err)
		}
	}

	if id == 0 {
		_, err = h.views.CreateCustomer(firstname, lastname, birthdate, gender, email, homeaddress)
	} else {
		_, err = h.views.EditCustomer(id, firstname, lastname, birthdate, gender, email, homeaddress)
	}

	if err != nil {
		h.ErrorPage(w, err)
	}

	http.Redirect(w, r, "/", 302)

}

func (h *HandlerStrust) HandleSaveCustomerEstonian(w http.ResponseWriter, r *http.Request) {

	var id = 0
	var err error
	r.ParseForm()
	params := r.PostForm
	idStr := params.Get("id")

	if len(idStr) > 0 {
		id, err = strconv.Atoi(idStr)
		if err != nil {
			h.ErrorPage(w, err)
			return
		}
	}

	firstname := strings.Title(strings.ToLower(params.Get("firstname")))
	lastname := strings.Title(strings.ToLower(params.Get("lastname")))
	gender := params.Get("gender")
	email := params.Get("email")
	homeaddress := params.Get("homeaddress")

	BirthDateStr := params.Get("birthdate")
	var birthdate time.Time

	if len(BirthDateStr) > 0 {
		birthdate, err = time.Parse("02 January 2006", BirthDateStr)
		if err != nil {
			h.ErrorPage(w, err)
		}
	}

	if id == 0 {
		_, err = h.views.CreateCustomer(firstname, lastname, birthdate, gender, email, homeaddress)
	} else {
		_, err = h.views.EditCustomer(id, firstname, lastname, birthdate, gender, email, homeaddress)
	}

	if err != nil {
		h.ErrorPage(w, err)
	}

	http.Redirect(w, r, "/ee", 302)

}

func (h *HandlerStrust) HandleSaveCustomerRussian(w http.ResponseWriter, r *http.Request) {

	var id = 0
	var err error
	r.ParseForm()
	params := r.PostForm
	idStr := params.Get("id")

	if len(idStr) > 0 {
		id, err = strconv.Atoi(idStr)
		if err != nil {
			h.ErrorPage(w, err)
			return
		}
	}

	firstname := strings.Title(strings.ToLower(params.Get("firstname")))
	lastname := strings.Title(strings.ToLower(params.Get("lastname")))
	gender := params.Get("gender")
	email := params.Get("email")
	homeaddress := params.Get("homeaddress")

	BirthDateStr := params.Get("birthdate")
	var birthdate time.Time

	if len(BirthDateStr) > 0 {
		birthdate, err = time.Parse("02 January 2006", BirthDateStr)
		if err != nil {
			h.ErrorPage(w, err)
		}
	}

	if id == 0 {
		_, err = h.views.CreateCustomer(firstname, lastname, birthdate, gender, email, homeaddress)
	} else {
		_, err = h.views.EditCustomer(id, firstname, lastname, birthdate, gender, email, homeaddress)
	}

	if err != nil {
		h.ErrorPage(w, err)
	}

	http.Redirect(w, r, "/ru", 302)

}

func (h *HandlerStrust) HandleListCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.views.AllCustomers()
	if err != nil {
		h.ErrorPage(w, err)
		return
	}

	buf, err := ioutil.ReadFile("web/index.html")
	if err != nil {
		h.ErrorPage(w, err)
		return
	}

	var page = IndexPage{AllCustomers: customers}
	IndexPage := string(buf)
	t := template.Must(template.New("IndexPage").Parse(IndexPage))
	t.Execute(w, page)
}

func (h *HandlerStrust) HandleListCustomersEstonian(w http.ResponseWriter, r *http.Request) {
	customers, err := h.views.AllCustomers()
	if err != nil {
		h.ErrorPage(w, err)
		return
	}

	buf, err := ioutil.ReadFile("web/index_ee.html")
	if err != nil {
		h.ErrorPage(w, err)
		return
	}

	var page = IndexPage{AllCustomers: customers}
	IndexPage := string(buf)
	t := template.Must(template.New("IndexPage").Parse(IndexPage))
	t.Execute(w, page)
}

func (h *HandlerStrust) HandleListCustomersRussian(w http.ResponseWriter, r *http.Request) {
	customers, err := h.views.AllCustomers()
	if err != nil {
		h.ErrorPage(w, err)
		return
	}

	buf, err := ioutil.ReadFile("web/index_ru.html")
	if err != nil {
		h.ErrorPage(w, err)
		return
	}

	var page = IndexPage{AllCustomers: customers}
	IndexPage := string(buf)
	t := template.Must(template.New("IndexPage").Parse(IndexPage))
	t.Execute(w, page)
}

func (h *HandlerStrust) HandleSearchCustomer(w http.ResponseWriter, r *http.Request) {
	var customers []Customer
	var err error
	r.ParseForm()
	params := r.PostForm
	name := params.Get("SearchName")
	name = strings.Title(strings.ToLower(name))

	if len(name) == 0 {
		customers, err = h.views.AllCustomers()
	} else {
		customers, err = h.views.SearchCustomers(name)
	}

	if err != nil {
		h.ErrorPage(w, err)
		return
	}

	buf, err := ioutil.ReadFile("web/index.html")
	if err != nil {
		h.ErrorPage(w, err)
		return
	}

	var page = IndexPage{AllCustomers: customers}
	IndexPage := string(buf)
	t := template.Must(template.New("IndexPage").Parse(IndexPage))
	err = t.Execute(w, page)
	if err != nil {
		h.ErrorOfServer(w, err)
	}

}

func (h *HandlerStrust) HandleSearchCustomerEstonian(w http.ResponseWriter, r *http.Request) {
	var customers []Customer
	var err error
	r.ParseForm()
	params := r.PostForm
	name := params.Get("SearchName")
	name = strings.Title(strings.ToLower(name))

	if len(name) == 0 {
		customers, err = h.views.AllCustomers()
	} else {
		customers, err = h.views.SearchCustomers(name)
	}

	if err != nil {
		h.ErrorPage(w, err)
		return
	}

	buf, err := ioutil.ReadFile("web/index_ee.html")
	if err != nil {
		h.ErrorPage(w, err)
		return
	}

	var page = IndexPage{AllCustomers: customers}
	IndexPage := string(buf)
	t := template.Must(template.New("IndexPage").Parse(IndexPage))
	err = t.Execute(w, page)
	if err != nil {
		h.ErrorOfServer(w, err)
	}

}

func (h *HandlerStrust) HandleSearchCustomerRussian(w http.ResponseWriter, r *http.Request) {
	var customers []Customer
	var err error
	r.ParseForm()
	params := r.PostForm
	name := params.Get("SearchName")
	name = strings.Title(strings.ToLower(name))

	if len(name) == 0 {
		customers, err = h.views.AllCustomers()
	} else {
		customers, err = h.views.SearchCustomers(name)
	}

	if err != nil {
		h.ErrorPage(w, err)
		return
	}

	buf, err := ioutil.ReadFile("web/index_ru.html")
	if err != nil {
		h.ErrorPage(w, err)
		return
	}

	var page = IndexPage{AllCustomers: customers}
	IndexPage := string(buf)
	t := template.Must(template.New("IndexPage").Parse(IndexPage))
	err = t.Execute(w, page)
	if err != nil {
		h.ErrorOfServer(w, err)
	}

}

func (h *HandlerStrust) HandleViewCustomer(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")

	var currentCustomer = Customer{}
	currentCustomer.BirthDate = time.Now()

	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			h.ErrorPage(w, err)
			return
		}

		currentCustomer, err = h.views.GetCustomer(id)
		if err != nil {
			h.ErrorPage(w, err)
			return
		}
	}

	buf, err := ioutil.ReadFile("web/add.html")
	if err != nil {
		h.ErrorPage(w, err)
		return
	}

	var page = CustomerPage{TargetCustomer: currentCustomer}
	customerPage := string(buf)
	t := template.Must(template.New("customerPage").Parse(customerPage))
	t.Execute(w, page)

}

func (h *HandlerStrust) HandleViewCustomerEstonian(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")

	var currentCustomer = Customer{}
	currentCustomer.BirthDate = time.Now()

	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			h.ErrorPage(w, err)
			return
		}

		currentCustomer, err = h.views.GetCustomer(id)
		if err != nil {
			h.ErrorPage(w, err)
			return
		}
	}

	buf, err := ioutil.ReadFile("web/add_ee.html")
	if err != nil {
		h.ErrorPage(w, err)
		return
	}

	var page = CustomerPage{TargetCustomer: currentCustomer}
	customerPage := string(buf)
	t := template.Must(template.New("customerPage").Parse(customerPage))
	t.Execute(w, page)

}

func (h *HandlerStrust) HandleViewCustomerRussian(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")

	var currentCustomer = Customer{}
	currentCustomer.BirthDate = time.Now()

	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			h.ErrorPage(w, err)
			return
		}

		currentCustomer, err = h.views.GetCustomer(id)
		if err != nil {
			h.ErrorPage(w, err)
			return
		}
	}

	buf, err := ioutil.ReadFile("web/add_ru.html")
	if err != nil {
		h.ErrorPage(w, err)
		return
	}

	var page = CustomerPage{TargetCustomer: currentCustomer}
	customerPage := string(buf)
	t := template.Must(template.New("customerPage").Parse(customerPage))
	t.Execute(w, page)

}

func (h *HandlerStrust) HandleDeleteCustomer(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")

	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			h.ErrorPage(w, err)
			return
		}

		n, err := h.views.DeleteCustomer(id)
		if err != nil {
			h.ErrorPage(w, err)
			return
		}

		fmt.Printf("Rows deleted: %v\n", n)
	}
	http.Redirect(w, r, "/", 302)
}

func (h *HandlerStrust) HandleDeleteCustomerEstonian(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")

	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			h.ErrorPage(w, err)
			return
		}

		n, err := h.views.DeleteCustomer(id)
		if err != nil {
			h.ErrorPage(w, err)
			return
		}

		fmt.Printf("Rows deleted: %v\n", n)
	}
	http.Redirect(w, r, "/ee", 302)
}

func (h *HandlerStrust) HandleDeleteCustomerRussian(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")

	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			h.ErrorPage(w, err)
			return
		}

		n, err := h.views.DeleteCustomer(id)
		if err != nil {
			h.ErrorPage(w, err)
			return
		}

		fmt.Printf("Rows deleted: %v\n", n)
	}
	http.Redirect(w, r, "/ru", 302)
}

func (h *HandlerStrust) ErrorPage(w http.ResponseWriter, errorMsg error) {
	buf, err := ioutil.ReadFile("web/error.html")

	if err != nil {
		log.Printf("%v\n", err)
		fmt.Fprintf(w, "%v\n", err)
		return
	}

	var page = ErrorPage{ErrorMsg: errorMsg.Error()}
	errorPage := string(buf)
	t := template.Must(template.New("errorPage").Parse(errorPage))
	t.Execute(w, page)
}
