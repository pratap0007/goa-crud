// Code generated by goa v3.1.1, DO NOT EDIT.
//
// blog service
//
// Command:
// $ goa gen crud/design

package blog

import (
	"context"
)

// The blog service gives blog details.
type Service interface {
	// Add new blog and return its ID.
	Create(context.Context, *CreatePayload) (res *Blog, err error)
	// List all entries
	List(context.Context) (res []*Storedblog, err error)
	// Remove blog from storage
	Remove(context.Context, *RemovePayload) (err error)
	// Updating the existing blog
	Update(context.Context, *UpdatePayload) (err error)
	// Add new blog and return its ID.
	Add(context.Context, *NewComment) (res *NewComment, err error)
	// Show blog based on the id given
	Show(context.Context, *Blog) (res *Blog, err error)
	// Github authentication to post a new blog
	Oauth(context.Context, *OauthPayload) (res string, err error)
	// Getting auth
	JWT(context.Context, *JWTPayload) (res string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "blog"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [8]string{"create", "list", "remove", "update", "add", "show", "oauth", "jwt"}

// CreatePayload is the payload type of the blog service create method.
type CreatePayload struct {
	// Adding a new blog
	Blog *Blog
	// Access github token
	Auth string
}

// Blog is the result type of the blog service create method.
type Blog struct {
	// ID of a person
	ID *uint32
	// Name of person
	Name *string
	// Comments
	Comments []*Comments
}

// RemovePayload is the payload type of the blog service remove method.
type RemovePayload struct {
	// ID of blog to remove
	ID uint32
}

// UpdatePayload is the payload type of the blog service update method.
type UpdatePayload struct {
	// ID of blog to be updated
	ID *uint32
	// Details of blog to be updated
	Name string
	// Comments to be updated
	Comments []*Comments
}

// NewComment is the payload type of the blog service add method.
type NewComment struct {
	// Id of blog
	ID *uint32
	// Comment added to an existing blog
	Comments *Comments
}

// OauthPayload is the payload type of the blog service oauth method.
type OauthPayload struct {
	// Access github token
	Token *string
}

// JWTPayload is the payload type of the blog service jwt method.
type JWTPayload struct {
	// Access JWT token
	Auth *string
}

// Id and comments
type Comments struct {
	// ID of a comment
	ID *uint32
	// Comment for the blog
	Comments *string
}

// A Storedblog describes a blog retrieved by the storage service.
type Storedblog struct {
	// ID is the unique id of the blog.
	ID uint32
	// Name of person
	Name string
	// Comments
	Comments []*Comments
}

// NotFound is the type returned when attempting to show or delete a blog that
// does not exist.
type NotFound struct {
	// ID of missing blog
	ID uint32
}

// Error returns an error description.
func (e *NotFound) Error() string {
	return "NotFound is the type returned when attempting to show or delete a blog that does not exist."
}

// ErrorName returns "NotFound".
func (e *NotFound) ErrorName() string {
	return "not_found"
}
