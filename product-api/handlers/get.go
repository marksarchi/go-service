package handlers

import (
	"net/http"
	"github.com/sarchimark/go-service/product-api/data"
	"context"
	protos "github.com/marksarchi/go-service/currency/protos/currency"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")
	rw.Header().Add("Content-Type","application/json")

	prods := data.GetProducts()

	err := data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Return a single product from the database
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	 
	//get exchange rate
    rr:= &protos.RateRequest{
		Base : protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value["GBP"]),
	
	}
	

	resp, err := p.cc.GetRate(context.Background(),rr)
	p.l.Println("[DEBUG] transaction rate", resp.Rate) 

	rate := resp.Rate
	p.l.Println("[DEBUG] Rate=", rate)

	_rate := resp.GetRate()
	p.l.Println("[DEBUG] GetRate()", _rate)

	if err!= nil {
		p.l.Println("[Error] getting new rate" )
		data.ToJSON(&GenericError{Message:err.Error()},rw)
		return
	}
	prod.Price = prod.Price * resp.GetRate()
	p.l.Println("[DEBUG] get record price", prod.Price)

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}