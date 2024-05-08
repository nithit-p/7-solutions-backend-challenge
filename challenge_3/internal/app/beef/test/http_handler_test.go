package beef_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nithit-p/7-solutions-backend-challenge/challenge_3/internal/app/beef"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHttpHandlerGetBeefSummary(t *testing.T) {
	tests := []struct {
		giveBody        string
		mockBeefSummary *beef.BeefSummary
		wantCode        int
		wantBody        string
	}{
		{
			giveBody: `mock`,
			mockBeefSummary: &beef.BeefSummary{
				Beef: map[string]int{
					"test": 2,
				},
			},
			wantCode: http.StatusOK,
			wantBody: `{"beef":{"test":2}}
`,
		},
		{
			giveBody: `mock2`,
			mockBeefSummary: &beef.BeefSummary{
				Beef: map[string]int{
					"test":  2,
					"test2": 3,
					"test3": 4,
				},
			},
			wantCode: http.StatusOK,
			wantBody: `{"beef":{"test":2,"test2":3,"test3":4}}
`,
		},
	}

	e := echo.New()

	for _, test := range tests {
		req := httptest.NewRequest(http.MethodPost, "/beef/summary", bytes.NewBufferString(test.giveBody))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetRequest(req)

		h, err := beef.NewBeefHttpHandler(&beefUsercaseMock{
			GetSummaryReturn: test.mockBeefSummary,
		})
		require.NoError(t, err)
		if assert.NoError(t, h.GetBeefSummary(c)) {
			assert.Equal(t, test.wantCode, rec.Code)
			assert.Equal(t, test.wantBody, rec.Body.String())
		}
	}
}

func TestHttpHandlerSummaryError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/beef/summary", bytes.NewBufferString(""))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetRequest(req)
	h, err := beef.NewBeefHttpHandler(&beefUsercaseMock{
		GetSummaryError: assert.AnError,
	})
	require.NoError(t, err)
	if assert.NoError(t, h.GetBeefSummary(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
}
