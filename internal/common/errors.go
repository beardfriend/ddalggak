package common

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
)

type UsecaseErr interface {
	Error() string
}

type UsecaseError struct {
	HttpStatusCode int
	UsecaseError   error
	OriginError    error
	isPushMessage  bool
}

func (e UsecaseError) Error() string {
	return fmt.Sprintf("origin:%s  || usecase: %s", e.OriginError, e.UsecaseError)
}

func NewUsecaseError(httpStatusCode int, before error, new error, isPushMessage ...bool) UsecaseErr {
	var pushMessage bool
	if len(isPushMessage) > 0 {
		pushMessage = isPushMessage[0]
	}
	return UsecaseError{HttpStatusCode: httpStatusCode, OriginError: before, UsecaseError: new, isPushMessage: pushMessage}
}

func ParseError(c *fiber.Ctx, err error) error {
	e, ok := err.(UsecaseError)
	if ok {
		if e.isPushMessage {
			// send error meesage logic
		}

		if e.HttpStatusCode == 0 || e.HttpStatusCode == 500 {
			ErrorHttpLog(err, string(debug.Stack()), e.HttpStatusCode)
			return c.Status(http.StatusInternalServerError).JSON(Response{Message: "Internal Server Error"})
		}

		return c.Status(e.HttpStatusCode).JSON(Response{Message: e.UsecaseError.Error()})
	}

	return c.Status(http.StatusInternalServerError).JSON(Response{Message: "Internal Server Error"})
}
