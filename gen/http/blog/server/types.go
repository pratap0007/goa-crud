// Code generated by goa v3.1.1, DO NOT EDIT.
//
// blog HTTP server types
//
// Command:
// $ goa gen crud/design

package server

import (
	blog "crud/gen/blog"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "blog" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// Adding a new blog
	Blog *BlogRequestBody `form:"blog,omitempty" json:"blog,omitempty" xml:"blog,omitempty"`
}

// UpdateRequestBody is the type of the "blog" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// Details of blog to be updated
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Comments to be updated
	Comments []*CommentsRequestBody `form:"comments,omitempty" json:"comments,omitempty" xml:"comments,omitempty"`
}

// AddRequestBody is the type of the "blog" service "add" endpoint HTTP request
// body.
type AddRequestBody struct {
	// Comment added to an existing blog
	Comments *CommentsRequestBody `form:"comments,omitempty" json:"comments,omitempty" xml:"comments,omitempty"`
}

// ShowRequestBody is the type of the "blog" service "show" endpoint HTTP
// request body.
type ShowRequestBody struct {
	// Name of person
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Comments
	Comments []*CommentsRequestBody `form:"comments,omitempty" json:"comments,omitempty" xml:"comments,omitempty"`
}

// OauthRequestBody is the type of the "blog" service "oauth" endpoint HTTP
// request body.
type OauthRequestBody struct {
	// Access github token
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

// CreateResponseBody is the type of the "blog" service "create" endpoint HTTP
// response body.
type CreateResponseBody struct {
	// ID of a person
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of person
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Comments
	Comments []*CommentsResponseBody `form:"comments,omitempty" json:"comments,omitempty" xml:"comments,omitempty"`
}

// ListResponseBody is the type of the "blog" service "list" endpoint HTTP
// response body.
type ListResponseBody []*StoredblogResponse

// AddResponseBody is the type of the "blog" service "add" endpoint HTTP
// response body.
type AddResponseBody struct {
	// Id of blog
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Comment added to an existing blog
	Comments *CommentsResponseBody `form:"comments,omitempty" json:"comments,omitempty" xml:"comments,omitempty"`
}

// ShowResponseBody is the type of the "blog" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID of a person
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of person
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Comments
	Comments []*CommentsResponseBody `form:"comments,omitempty" json:"comments,omitempty" xml:"comments,omitempty"`
}

// CommentsResponseBody is used to define fields on response body types.
type CommentsResponseBody struct {
	// ID of a comment
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Comment for the blog
	Comments *string `form:"comments,omitempty" json:"comments,omitempty" xml:"comments,omitempty"`
}

// StoredblogResponse is used to define fields on response body types.
type StoredblogResponse struct {
	// ID is the unique id of the blog.
	ID uint32 `form:"id" json:"id" xml:"id"`
	// Name of person
	Name string `form:"name" json:"name" xml:"name"`
	// Comments
	Comments []*CommentsResponse `form:"comments,omitempty" json:"comments,omitempty" xml:"comments,omitempty"`
}

// CommentsResponse is used to define fields on response body types.
type CommentsResponse struct {
	// ID of a comment
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Comment for the blog
	Comments *string `form:"comments,omitempty" json:"comments,omitempty" xml:"comments,omitempty"`
}

// BlogRequestBody is used to define fields on request body types.
type BlogRequestBody struct {
	// ID of a person
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of person
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Comments
	Comments []*CommentsRequestBody `form:"comments,omitempty" json:"comments,omitempty" xml:"comments,omitempty"`
}

// CommentsRequestBody is used to define fields on request body types.
type CommentsRequestBody struct {
	// ID of a comment
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Comment for the blog
	Comments *string `form:"comments,omitempty" json:"comments,omitempty" xml:"comments,omitempty"`
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "blog" service.
func NewCreateResponseBody(res *blog.Blog) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:   res.ID,
		Name: res.Name,
	}
	if res.Comments != nil {
		body.Comments = make([]*CommentsResponseBody, len(res.Comments))
		for i, val := range res.Comments {
			body.Comments[i] = marshalBlogCommentsToCommentsResponseBody(val)
		}
	}
	return body
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "blog" service.
func NewListResponseBody(res []*blog.Storedblog) ListResponseBody {
	body := make([]*StoredblogResponse, len(res))
	for i, val := range res {
		body[i] = marshalBlogStoredblogToStoredblogResponse(val)
	}
	return body
}

// NewAddResponseBody builds the HTTP response body from the result of the
// "add" endpoint of the "blog" service.
func NewAddResponseBody(res *blog.NewComment) *AddResponseBody {
	body := &AddResponseBody{
		ID: res.ID,
	}
	if res.Comments != nil {
		body.Comments = marshalBlogCommentsToCommentsResponseBody(res.Comments)
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "blog" service.
func NewShowResponseBody(res *blog.Blog) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:   res.ID,
		Name: res.Name,
	}
	if res.Comments != nil {
		body.Comments = make([]*CommentsResponseBody, len(res.Comments))
		for i, val := range res.Comments {
			body.Comments[i] = marshalBlogCommentsToCommentsResponseBody(val)
		}
	}
	return body
}

// NewCreatePayload builds a blog service create endpoint payload.
func NewCreatePayload(body *CreateRequestBody, auth string) *blog.CreatePayload {
	v := &blog.CreatePayload{}
	if body.Blog != nil {
		v.Blog = unmarshalBlogRequestBodyToBlogBlog(body.Blog)
	}
	v.Auth = auth

	return v
}

// NewRemovePayload builds a blog service remove endpoint payload.
func NewRemovePayload(id uint32) *blog.RemovePayload {
	v := &blog.RemovePayload{}
	v.ID = id

	return v
}

// NewUpdatePayload builds a blog service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, id uint32) *blog.UpdatePayload {
	v := &blog.UpdatePayload{
		Name: *body.Name,
	}
	v.Comments = make([]*blog.Comments, len(body.Comments))
	for i, val := range body.Comments {
		v.Comments[i] = unmarshalCommentsRequestBodyToBlogComments(val)
	}
	v.ID = &id

	return v
}

// NewAddNewComment builds a blog service add endpoint payload.
func NewAddNewComment(body *AddRequestBody, id uint32) *blog.NewComment {
	v := &blog.NewComment{}
	if body.Comments != nil {
		v.Comments = unmarshalCommentsRequestBodyToBlogComments(body.Comments)
	}
	v.ID = &id

	return v
}

// NewShowBlog builds a blog service show endpoint payload.
func NewShowBlog(body *ShowRequestBody, id uint32) *blog.Blog {
	v := &blog.Blog{
		Name: body.Name,
	}
	if body.Comments != nil {
		v.Comments = make([]*blog.Comments, len(body.Comments))
		for i, val := range body.Comments {
			v.Comments[i] = unmarshalCommentsRequestBodyToBlogComments(val)
		}
	}
	v.ID = &id

	return v
}

// NewOauthPayload builds a blog service oauth endpoint payload.
func NewOauthPayload(body *OauthRequestBody) *blog.OauthPayload {
	v := &blog.OauthPayload{
		Token: body.Token,
	}

	return v
}

// NewJWTPayload builds a blog service jwt endpoint payload.
func NewJWTPayload(auth *string) *blog.JWTPayload {
	v := &blog.JWTPayload{}
	v.Auth = auth

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Blog != nil {
		if err2 := ValidateBlogRequestBody(body.Blog); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Comments == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("comments", "body"))
	}
	return
}

// ValidateShowRequestBody runs the validations defined on ShowRequestBody
func ValidateShowRequestBody(body *ShowRequestBody) (err error) {
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if len(body.Comments) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.comments", body.Comments, len(body.Comments), 100, false))
	}
	return
}

// ValidateBlogRequestBody runs the validations defined on BlogRequestBody
func ValidateBlogRequestBody(body *BlogRequestBody) (err error) {
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if len(body.Comments) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.comments", body.Comments, len(body.Comments), 100, false))
	}
	return
}
