package handler

import (
	"app/internal/service"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestGetTotalAmountTickets(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// ARRANGE
		// create the mock of the service
		service := new(service.ServiceTicketDefaultMock)
		service.On("GetTotalAmountTickets").Return(10, nil)
		expected := `{"message":"succesfully fetched total","data":{"total":10}}`
		// create the handler
		handler := NewHandlerTicketDefault(service)
		// create the request
		req := httptest.NewRequest("GET", "/tickets/total", nil)
		// create the response
		w := httptest.NewRecorder()
		// create the function
		handfunc := http.HandlerFunc(handler.GetTotalAmountTickets())

		// ACT
		handfunc.ServeHTTP(w, req)

		// ASSERT
		service.AssertCalled(t, "GetTotalAmountTickets")
		require.Equal(t, http.StatusOK, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})
	t.Run("error", func(t *testing.T) {
		// ARRANGE
		// create the mock of the service
		service := new(service.ServiceTicketDefaultMock)
		service.On("GetTotalAmountTickets").Return(0, errors.New("error getting the tickets"))
		expected := `{"message":"error getting the tickets", "status":"Internal Server Error"}`
		// create the handler
		handler := NewHandlerTicketDefault(service)
		// create the request
		req := httptest.NewRequest("GET", "/tickets/total", nil)
		// create the response
		w := httptest.NewRecorder()
		// create the function
		handfunc := http.HandlerFunc(handler.GetTotalAmountTickets())

		// ACT
		handfunc.ServeHTTP(w, req)

		// ASSERT
		service.AssertCalled(t, "GetTotalAmountTickets")
		require.Equal(t, http.StatusInternalServerError, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})
}

func TestGetTicketsAmountByDestinationCountry(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// ARRANGE
		// create the mock of the service
		service := new(service.ServiceTicketDefaultMock)
		service.On("GetTicketsAmountByDestinationCountry", "argentina").Return(10, nil)
		expected := `{"message":"succesfully fetched total","data":{"total":10}}`
		// create the handler
		handler := NewHandlerTicketDefault(service)
		// create a new RouteContext and set the "id" URL parameter
		routeContext := chi.NewRouteContext()
		routeContext.URLParams.Add("dest", "argentina")
		// create the request
		req := httptest.NewRequest("GET", "/tickets/getByCountry/", nil)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeContext))
		// create the response
		w := httptest.NewRecorder()
		// create the function
		handfunc := http.HandlerFunc(handler.GetTotalAmountTicketsByDestinationCountry())

		// ACT
		handfunc.ServeHTTP(w, req)

		// ASSERT
		service.AssertCalled(t, "GetTicketsAmountByDestinationCountry", "argentina")
		require.Equal(t, http.StatusOK, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})

	t.Run("error", func(t *testing.T) {
		// ARRANGE
		// create the mock of the service
		service := new(service.ServiceTicketDefaultMock)
		service.On("GetTicketsAmountByDestinationCountry", "argentina").Return(0, errors.New("error getting the tickets"))
		expected := `{"message":"error getting the tickets", "status":"Internal Server Error"}`
		// create the handler
		handler := NewHandlerTicketDefault(service)
		// create a new RouteContext and set the "id" URL parameter
		routeContext := chi.NewRouteContext()
		routeContext.URLParams.Add("dest", "argentina")
		// create the request
		req := httptest.NewRequest("GET", "/tickets/getByCountry/", nil)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeContext))
		// create the response
		w := httptest.NewRecorder()
		// create the function
		handfunc := http.HandlerFunc(handler.GetTotalAmountTicketsByDestinationCountry())

		// ACT
		handfunc.ServeHTTP(w, req)

		// ASSERT
		service.AssertCalled(t, "GetTicketsAmountByDestinationCountry", "argentina")
		require.Equal(t, http.StatusInternalServerError, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})
}

func TestGetPercentageTicketsByDestinationCountry(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// ARRANGE
		// create the mock of the service
		service := new(service.ServiceTicketDefaultMock)
		service.On("GetPercentageTicketsByDestinationCountry", "argentina").Return(0.5, nil)
		expected := `{"message":"succesfully fetched total","data":{"percentage":0.5}}`
		// create the handler
		handler := NewHandlerTicketDefault(service)
		// create a new RouteContext and set the "id" URL parameter
		routeContext := chi.NewRouteContext()
		routeContext.URLParams.Add("dest", "argentina")
		// create the request
		req := httptest.NewRequest("GET", "/tickets/getByCountry/", nil)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeContext))
		// create the response
		w := httptest.NewRecorder()
		// create the function
		handfunc := http.HandlerFunc(handler.GetPercentageTicketsByDestinationCountry())

		// ACT
		handfunc.ServeHTTP(w, req)

		// ASSERT
		service.AssertCalled(t, "GetPercentageTicketsByDestinationCountry", "argentina")
		require.Equal(t, http.StatusOK, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})
	t.Run("error", func(t *testing.T) {
		// ARRANGE
		// create the mock of the service
		service := new(service.ServiceTicketDefaultMock)
		service.On("GetPercentageTicketsByDestinationCountry", "argentina").Return(0, errors.New("error getting the tickets"))
		expected := `{"message":"error getting the tickets", "status":"Internal Server Error"}`
		// create the handler
		handler := NewHandlerTicketDefault(service)
		// create a new RouteContext and set the "id" URL parameter
		routeContext := chi.NewRouteContext()
		routeContext.URLParams.Add("dest", "argentina")
		// create the request
		req := httptest.NewRequest("GET", "/tickets/getByCountry/", nil)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeContext))
		// create the response
		w := httptest.NewRecorder()
		// create the function
		handfunc := http.HandlerFunc(handler.GetPercentageTicketsByDestinationCountry())

		// ACT
		handfunc.ServeHTTP(w, req)

		// ASSERT
		service.AssertCalled(t, "GetPercentageTicketsByDestinationCountry", "argentina")
		require.Equal(t, http.StatusInternalServerError, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})
}
