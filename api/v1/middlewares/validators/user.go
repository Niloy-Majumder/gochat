package validators

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gochat/db/mongoDB/models"
	"strings"
)

func CreateUserValidator(c *fiber.Ctx) error {
	myValidator := &XValidator{
		validator: validate,
	}
	userData := models.User{}
	_ = json.Unmarshal(c.Body(), &models.User{})

	if errs := myValidator.Validate(userData); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errMsgs, " and "),
		}
	}

	// Logic, validated with success
	return c.Next()
}
