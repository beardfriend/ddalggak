package user

import (
    "github.com/beardfriend/ddalggak/ent"
    _ "github.com/beardfriend/ddalggak/pkg/pagination"
    "github.com/beardfriend/ddalggak/pkg/validatorx"
     "github.com/beardfriend/ddalggak/internal/common"
    "net/http"
	"github.com/gofiber/fiber/v2"
	"strconv"
	
)

type API struct {
    usecase Usecase
    validator validatorx.Validator  
}

func NewAPI(usecase Usecase, validator validatorx.Validator) *API {
	return &API{usecase, validator}
}

// Create godoc
// @Summary      User Create
// @Description  User Create
// @Tags         user
// @Accept       json
// @Produce      json
// @Param request body ent.User true "body"
// @Param	id  path  int  true  "id"
// @Success      201  {object}	common.Response{}
// @Failure      400  {object}	common.Response{}
// @Failure      500  {object}	common.Response{}
// @Router       /users [post]
func (a *API) Create(c *fiber.Ctx) error {
	body := new(ent.User)
	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(common.Response{Message: "bad request"})
	}
	if err := a.validator.ValidateStruct(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(common.Response{Message: "bad request"})
	}
	err := a.usecase.Create(c.Context(), body)
	if err != nil {
		return common.ParseError(c, err)
	}
	return c.Status(http.StatusCreated).JSON(common.Response{
		Message: "created",
	})
}

// Get godoc
// @Summary      User Get
// @Description  User Get
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {object}	common.ResponseWithData{result=ent.User}
// @Failure      400  {object}	common.Response{}
// @Failure      500  {object}	common.Response{}
// @Router       /users/{id} [get]
func (a *API) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(common.Response{Message: "bad request"})
	}

	user, err := a.usecase.Get(c.Context(), idParsed)
	if err != nil {
		return common.ParseError(c, err)
	}

	return c.Status(http.StatusOK).JSON(common.ResponseWithData{
		Message: "ok",
		Result:  user,
	})
}

// List godoc
// @Summary      User List
// @Description  User List
// @Tags         user
// @Accept       json
// @Produce      json
// @Param request query common.ListRequest true "queries"
// @Success      200  {object}	common.ResponseWithPagination{result=ent.User, pagination=pagination.PaginationInfo}
// @Failure      400  {object}	common.Response{}
// @Failure      404  {object}	common.Response{}
// @Failure      500  {object}	common.Response{}
// @Router       /users [get]
func (a *API) List(c *fiber.Ctx) error {
    req := new(common.ListRequest)
    if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(common.Response{Message: "bad request"})
	}
	if err := a.validator.ValidateStruct(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(common.Response{Message: "bad request"})
	}

	list, pgInfo, err := a.usecase.List(c.Context(),req)
	if err != nil {
		return common.ParseError(c, err)
	}
	return c.Status(http.StatusOK).JSON(common.ResponseWithPagination{
		Message: "ok",
		Result:  list,
        Pagination: pgInfo,
	})
}

// Delete godoc
// @Summary      User Delete
// @Description  User Delete
// @Tags         user
// @Accept       json
// @Produce      json
// @Param	id  path  int  true  "id"
// @Success      200  {object}	common.Response{}
// @Failure      400  {object}	common.Response{}
// @Failure      500  {object}	common.Response{}
// @Router       /users/{id} [delete]
func (a *API) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(common.Response{Message: "bad request"})
	}
	err = a.usecase.Delete(c.Context(), idParsed)
	if err != nil {
		return common.ParseError(c, err)
	}

	return c.Status(http.StatusOK).JSON(common.Response{
		Message: "ok",
	})
}


// Update godoc
// @Summary      User Update
// @Description  User Update
// @Tags         user
// @Accept       json
// @Produce      json
// @Param request body ent.User true "body"
// @Param	id  path  int  true  "id"
// @Success      200  {object}	common.ResponseWithData{result=ent.User}
// @Failure      400  {object}	common.Response{}
// @Failure      500  {object}	common.Response{}
// @Router       /users/{id} [put]
func (a *API) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(common.Response{Message: "bad request"})
	}
	body := new(ent.User)
	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(common.Response{Message: "bad request"})
	}
	if err := a.validator.ValidateStruct(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(common.Response{Message: "bad request"})
	}
	body.ID = idParsed
	err = a.usecase.Update(c.Context(), body)
	if err != nil {
		return common.ParseError(c, err)
	}

	return c.Status(http.StatusOK).JSON(common.Response{
		Message: "ok",
	})
}