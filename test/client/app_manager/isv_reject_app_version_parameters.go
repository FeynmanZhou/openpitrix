// Code generated by go-swagger; DO NOT EDIT.

package app_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// NewIsvRejectAppVersionParams creates a new IsvRejectAppVersionParams object
// with the default values initialized.
func NewIsvRejectAppVersionParams() *IsvRejectAppVersionParams {
	var ()
	return &IsvRejectAppVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIsvRejectAppVersionParamsWithTimeout creates a new IsvRejectAppVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIsvRejectAppVersionParamsWithTimeout(timeout time.Duration) *IsvRejectAppVersionParams {
	var ()
	return &IsvRejectAppVersionParams{

		timeout: timeout,
	}
}

// NewIsvRejectAppVersionParamsWithContext creates a new IsvRejectAppVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewIsvRejectAppVersionParamsWithContext(ctx context.Context) *IsvRejectAppVersionParams {
	var ()
	return &IsvRejectAppVersionParams{

		Context: ctx,
	}
}

// NewIsvRejectAppVersionParamsWithHTTPClient creates a new IsvRejectAppVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIsvRejectAppVersionParamsWithHTTPClient(client *http.Client) *IsvRejectAppVersionParams {
	var ()
	return &IsvRejectAppVersionParams{
		HTTPClient: client,
	}
}

/*IsvRejectAppVersionParams contains all the parameters to send to the API endpoint
for the isv reject app version operation typically these are written to a http.Request
*/
type IsvRejectAppVersionParams struct {

	/*Body*/
	Body *models.OpenpitrixRejectAppVersionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the isv reject app version params
func (o *IsvRejectAppVersionParams) WithTimeout(timeout time.Duration) *IsvRejectAppVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the isv reject app version params
func (o *IsvRejectAppVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the isv reject app version params
func (o *IsvRejectAppVersionParams) WithContext(ctx context.Context) *IsvRejectAppVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the isv reject app version params
func (o *IsvRejectAppVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the isv reject app version params
func (o *IsvRejectAppVersionParams) WithHTTPClient(client *http.Client) *IsvRejectAppVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the isv reject app version params
func (o *IsvRejectAppVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the isv reject app version params
func (o *IsvRejectAppVersionParams) WithBody(body *models.OpenpitrixRejectAppVersionRequest) *IsvRejectAppVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the isv reject app version params
func (o *IsvRejectAppVersionParams) SetBody(body *models.OpenpitrixRejectAppVersionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IsvRejectAppVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
