package beef

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BeefHttpServer interface {
	GetBeefSummary(c echo.Context) error
}

type BeefHttpHandler struct {
	beefService BeefService
}

var _ BeefHttpServer = (*BeefHttpHandler)(nil)

func NewBeefHttpHandler(beefService BeefService) (*BeefHttpHandler, error) {
	return &BeefHttpHandler{
		beefService: beefService,
	}, nil
}

func (h *BeefHttpHandler) GetBeefSummary(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	summary, err := h.beefService.GetSummary(string(body))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, summary)
}
