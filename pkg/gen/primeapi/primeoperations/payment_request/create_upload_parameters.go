// Code generated by go-swagger; DO NOT EDIT.

package payment_request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCreateUploadParams creates a new CreateUploadParams object
// no default values defined in spec.
func NewCreateUploadParams() CreateUploadParams {

	return CreateUploadParams{}
}

// CreateUploadParams contains all the bound params for the create upload operation
// typically these are obtained from a http.Request
//
// swagger:parameters createUpload
type CreateUploadParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The file to upload.
	  Required: true
	  In: formData
	*/
	File io.ReadCloser
	/*UUID of payment request to use.
	  Required: true
	  In: path
	*/
	PaymentRequestID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateUploadParams() beforehand.
func (o *CreateUploadParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "file", err))
	} else if err := o.bindFile(file, fileHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.File = &runtime.File{Data: file, Header: fileHeader}
	}

	rPaymentRequestID, rhkPaymentRequestID, _ := route.Params.GetOK("paymentRequestID")
	if err := o.bindPaymentRequestID(rPaymentRequestID, rhkPaymentRequestID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFile binds file parameter File.
//
// The only supported validations on files are MinLength and MaxLength
func (o *CreateUploadParams) bindFile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}

// bindPaymentRequestID binds and validates parameter PaymentRequestID from path.
func (o *CreateUploadParams) bindPaymentRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PaymentRequestID = raw

	return nil
}