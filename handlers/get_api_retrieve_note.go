package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"ncrypt-api/helpers"
	"ncrypt-api/models"
	"ncrypt-api/processors"
	"net/http"
)

//GetSecureNoteV1 handle GET /api/v1/note/{id}
func (di *DI) GetSecureNoteV1(c echo.Context) error {
	id := c.Param("id")

	noteId, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			helpers.BuildResponse(
				http.StatusUnprocessableEntity,
				"input failure",
				nil,
				[]string{"invalid id provided"},
				nil,
			),
		)
	}

	payload := models.RetrieveNoteRequest{
		Id: noteId,
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

	note, err := processors.RetrieveSecureNote(di.StorageDriver, payload)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helpers.BuildResponse(
				http.StatusBadRequest,
				"unable to retrieve note",
				nil,
				[]string{err.Error()},
				nil,
			),
		)
	}

	response := models.RetrieveNoteResponse{Note: string(note)}

	return c.JSON(
		http.StatusOK,
		helpers.BuildResponse(
			http.StatusOK,
			"Note retrieved.",
			&response,
			nil,
			nil,
		),
	)
}
