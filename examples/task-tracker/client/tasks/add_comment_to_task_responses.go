package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/validate"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// AddCommentToTaskReader is a Reader for the AddCommentToTask structure.
type AddCommentToTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *AddCommentToTaskReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddCommentToTaskCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddCommentToTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewAddCommentToTaskCreated creates a AddCommentToTaskCreated with default headers values
func NewAddCommentToTaskCreated() *AddCommentToTaskCreated {
	return &AddCommentToTaskCreated{}
}

/*AddCommentToTaskCreated handles this case with default header values.

Comment added
*/
type AddCommentToTaskCreated struct {
}

func (o *AddCommentToTaskCreated) Error() string {
	return fmt.Sprintf("[POST /tasks/{id}/comments][%d] addCommentToTaskCreated ", 201)
}

func (o *AddCommentToTaskCreated) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddCommentToTaskDefault creates a AddCommentToTaskDefault with default headers values
func NewAddCommentToTaskDefault(code int) *AddCommentToTaskDefault {
	return &AddCommentToTaskDefault{
		_statusCode: code,
	}
}

/*AddCommentToTaskDefault handles this case with default header values.

AddCommentToTaskDefault add comment to task default
*/
type AddCommentToTaskDefault struct {
	_statusCode int
}

// Code gets the status code for the add comment to task default response
func (o *AddCommentToTaskDefault) Code() int {
	return o._statusCode
}

func (o *AddCommentToTaskDefault) Error() string {
	return fmt.Sprintf("[POST /tasks/{id}/comments][%d] addCommentToTask default ", o._statusCode)
}

func (o *AddCommentToTaskDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*AddCommentToTaskBody A comment to create

These values can have github flavored markdown.


swagger:model AddCommentToTaskBody
*/
type AddCommentToTaskBody struct {

	/* Content content

	Required: true
	*/
	Content string `json:"content,omitempty"`

	/* UserID user id

	Required: true
	*/
	UserID int64 `json:"userId,omitempty"`
}

// Validate validates this add comment to task body
func (o *AddCommentToTaskBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateUserID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddCommentToTaskBody) validateContent(formats strfmt.Registry) error {

	if err := validate.RequiredString("body"+"."+"content", "body", string(o.Content)); err != nil {
		return err
	}

	return nil
}

func (o *AddCommentToTaskBody) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"userId", "body", int64(o.UserID)); err != nil {
		return err
	}

	return nil
}
