package handlers

import (
	"github.com/labstack/echo"
	"ncrypt-api/helpers"
	"ncrypt-api/models"
	"ncrypt-api/processors"
	"net/http"
)

//PostStoreSecureNoteV1 handles POST /api/v1/note
func (di *DI) PostStoreSecureNoteV1(c echo.Context) error {
	payload := models.SecureMessageRequest{}

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			helpers.BuildResponse(
				http.StatusUnprocessableEntity,
				"request data not accepted",
				nil,
				nil,
				nil,
			),
		)
	}

	err = c.Validate(payload)
	if err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			helpers.BuildResponse(
				http.StatusUnprocessableEntity,
				"validation failed",
				nil,
				helpers.FormatValidationErrorMessage(err),
				nil,
			),
		)
	}

	messageUuid, err := processors.StoreMessage(di.StorageDriver, payload)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helpers.BuildResponse(
				http.StatusInternalServerError,
				"internal error occurred",
				nil,
				nil,
				nil,
			),
		)
	}

	response := models.SecureMessageResponse{
		Id:  messageUuid.String(),
		URL: di.ApplicationConfig.ApplicationUrlConfig.AppBaseUrl + "/note/" + messageUuid.String(),
	}

	return c.JSON(
		http.StatusCreated,
		helpers.BuildResponse(
			http.StatusCreated,
			"Note stored.",
			&response,
			nil,
			nil,
		),
	)
}
