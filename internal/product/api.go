package product

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
// @Summary      Product Create
// @Description  Product Create
// @Tags         product
// @Accept       json
// @Produce      json
// @Param request body ent.Product true "body"
// @Param	id  path  int  true  "id"
// @Success      201  {object}	common.Response{}
// @Failure      400  {object}	common.Response{}
// @Failure      500  {object}	common.Response{}
// @Router       /products [post]
func (a *API) Create(c *fiber.Ctx) error {
	body := new(ent.Product)
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
// @Summary      Product Get
// @Description  Product Get
// @Tags         product
// @Accept       json
// @Produce      json
// @Success      200  {object}	common.ResponseWithData{result=ent.Product}
// @Failure      400  {object}	common.Response{}
// @Failure      500  {object}	common.Response{}
// @Router       /products/{id} [get]
func (a *API) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(common.Response{Message: "bad request"})
	}

	product, err := a.usecase.Get(c.Context(), idParsed)
	if err != nil {
		return common.ParseError(c, err)
	}

	return c.Status(http.StatusOK).JSON(common.ResponseWithData{
		Message: "ok",
		Result:  product,
	})
}

// List godoc
// @Summary      Product List
// @Description  Product List
// @Tags         product
// @Accept       json
// @Produce      json
// @Param request query common.ListRequest true "queries"
// @Success      200  {object}	common.ResponseWithPagination{result=ent.Product, pagination=pagination.PaginationInfo}
// @Failure      400  {object}	common.Response{}
// @Failure      404  {object}	common.Response{}
// @Failure      500  {object}	common.Response{}
// @Router       /products [get]
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
// @Summary      Product Delete
// @Description  Product Delete
// @Tags         product
// @Accept       json
// @Produce      json
// @Param	id  path  int  true  "id"
// @Success      200  {object}	common.Response{}
// @Failure      400  {object}	common.Response{}
// @Failure      500  {object}	common.Response{}
// @Router       /products/{id} [delete]
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
// @Summary      Product Update
// @Description  Product Update
// @Tags         product
// @Accept       json
// @Produce      json
// @Param request body ent.Product true "body"
// @Param	id  path  int  true  "id"
// @Success      200  {object}	common.ResponseWithData{result=ent.Product}
// @Failure      400  {object}	common.Response{}
// @Failure      500  {object}	common.Response{}
// @Router       /products/{id} [put]
func (a *API) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(common.Response{Message: "bad request"})
	}
	body := new(ent.Product)
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