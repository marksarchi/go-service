// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/sarchimark/go-microservice/client/models"
)

// NewCreateProductParams creates a new CreateProductParams object
// with the default values initialized.
func NewCreateProductParams() *CreateProductParams {
	var ()
	return &CreateProductParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProductParamsWithTimeout creates a new CreateProductParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateProductParamsWithTimeout(timeout time.Duration) *CreateProductParams {
	var ()
	return &CreateProductParams{

		timeout: timeout,
	}
}

// NewCreateProductParamsWithContext creates a new CreateProductParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateProductParamsWithContext(ctx context.Context) *CreateProductParams {
	var ()
	return &CreateProductParams{

		Context: ctx,
	}
}

// NewCreateProductParamsWithHTTPClient creates a new CreateProductParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateProductParamsWithHTTPClient(client *http.Client) *CreateProductParams {
	var ()
	return &CreateProductParams{
		HTTPClient: client,
	}
}

/*CreateProductParams contains all the parameters to send to the API endpoint
for the create product operation typically these are written to a http.Request
*/
type CreateProductParams struct {

	/*Body
	  Product data structure to Update or Create.
	Note: the id field is ignored by update and create operations

	*/
	Body *models.Product

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create product params
func (o *CreateProductParams) WithTimeout(timeout time.Duration) *CreateProductParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create product params
func (o *CreateProductParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create product params
func (o *CreateProductParams) WithContext(ctx context.Context) *CreateProductParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create product params
func (o *CreateProductParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create product params
func (o *CreateProductParams) WithHTTPClient(client *http.Client) *CreateProductParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create product params
func (o *CreateProductParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create product params
func (o *CreateProductParams) WithBody(body *models.Product) *CreateProductParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create product params
func (o *CreateProductParams) SetBody(body *models.Product) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProductParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}