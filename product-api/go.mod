module github.com/sarchimark/go-service/product-api

go 1.14

require (
	github.com/go-openapi/runtime v0.19.22
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-hclog v0.14.1
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/marksarchi/go-service/currency v0.0.0-20201009133438-c26d01af1288
	github.com/nicholasjackson/env v0.6.0
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.32.0
)

replace github.com/marksarchi/go-service/currency => ../currency
