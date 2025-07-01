package controller

import "pet_shelter_and_store/internal/models"

func newErrorResponse(error string) models.ErrorResponse {
	return models.ErrorResponse{
		Error: error,
	}
}
