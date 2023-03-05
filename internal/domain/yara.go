package domain

import "be-interview-task/pkg/validation"

type GetResultPayload struct {
	Id string `json:"id" validate:"required"`
}

func (r *GetResultPayload) Validate() []*validation.ErrorResponse {
	return validation.Validate(r)
}
