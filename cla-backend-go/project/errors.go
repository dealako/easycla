// Copyright The Linux Foundation and each contributor to CommunityBridge.
// SPDX-License-Identifier: MIT

package project

import (
	"github.com/communitybridge/easycla/cla-backend-go/gen/models"
)

// codedResponse interface
type codedResponse interface {
	Code() string
}

// errorResponse is a helper to wrap the specified error into an error response model
func errorResponse(err error) *models.ErrorResponse {
	code := ""
	if e, ok := err.(codedResponse); ok {
		code = e.Code()
	}

	e := models.ErrorResponse{
		Code:    code,
		Message: err.Error(),
	}

	return &e
}
