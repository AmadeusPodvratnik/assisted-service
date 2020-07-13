// Code generated by go-swagger; DO NOT EDIT.

package installer

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

	"github.com/filanov/bm-inventory/models"
)

// NewCompleteInstallationParams creates a new CompleteInstallationParams object
// with the default values initialized.
func NewCompleteInstallationParams() *CompleteInstallationParams {
	var ()
	return &CompleteInstallationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompleteInstallationParamsWithTimeout creates a new CompleteInstallationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompleteInstallationParamsWithTimeout(timeout time.Duration) *CompleteInstallationParams {
	var ()
	return &CompleteInstallationParams{

		timeout: timeout,
	}
}

// NewCompleteInstallationParamsWithContext creates a new CompleteInstallationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompleteInstallationParamsWithContext(ctx context.Context) *CompleteInstallationParams {
	var ()
	return &CompleteInstallationParams{

		Context: ctx,
	}
}

// NewCompleteInstallationParamsWithHTTPClient creates a new CompleteInstallationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompleteInstallationParamsWithHTTPClient(client *http.Client) *CompleteInstallationParams {
	var ()
	return &CompleteInstallationParams{
		HTTPClient: client,
	}
}

/*CompleteInstallationParams contains all the parameters to send to the API endpoint
for the complete installation operation typically these are written to a http.Request
*/
type CompleteInstallationParams struct {

	/*ClusterID*/
	ClusterID strfmt.UUID
	/*CompletionParams*/
	CompletionParams *models.CompletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the complete installation params
func (o *CompleteInstallationParams) WithTimeout(timeout time.Duration) *CompleteInstallationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the complete installation params
func (o *CompleteInstallationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the complete installation params
func (o *CompleteInstallationParams) WithContext(ctx context.Context) *CompleteInstallationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the complete installation params
func (o *CompleteInstallationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the complete installation params
func (o *CompleteInstallationParams) WithHTTPClient(client *http.Client) *CompleteInstallationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the complete installation params
func (o *CompleteInstallationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the complete installation params
func (o *CompleteInstallationParams) WithClusterID(clusterID strfmt.UUID) *CompleteInstallationParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the complete installation params
func (o *CompleteInstallationParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithCompletionParams adds the completionParams to the complete installation params
func (o *CompleteInstallationParams) WithCompletionParams(completionParams *models.CompletionParams) *CompleteInstallationParams {
	o.SetCompletionParams(completionParams)
	return o
}

// SetCompletionParams adds the completionParams to the complete installation params
func (o *CompleteInstallationParams) SetCompletionParams(completionParams *models.CompletionParams) {
	o.CompletionParams = completionParams
}

// WriteToRequest writes these params to a swagger request
func (o *CompleteInstallationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if o.CompletionParams != nil {
		if err := r.SetBodyParam(o.CompletionParams); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
