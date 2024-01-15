package handler

import (
	"app/internal"
	"app/platform/web/response"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewHandlerTicketDefault creates a new default handler of the tickets
func NewHandlerTicketDefault(sv internal.ServiceTicket) *HandlerTicketDefault {
	return &HandlerTicketDefault{
		sv: sv,
	}
}

// HandlerTicketDefault represents the default handler of the tickets
type HandlerTicketDefault struct {
	sv internal.ServiceTicket
}

// TotalTicketsResponse represents the response of the total number of tickets
type TotalTicketsResponse struct {
	Total int `json:"total"`
}

// Get returns the total amount of tickets
func (h *HandlerTicketDefault) GetTotalAmountTickets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the total amount of tickets
		total, err := h.sv.GetTotalAmountTickets()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		// write the response
		response.JSON(w, http.StatusOK, "succesfully fetched total", &TotalTicketsResponse{
			Total: total,
		})
	}
}

// GetTotalAmountTicketsByDestinationCountry returns the total amount of tickets by destination country
func (h *HandlerTicketDefault) GetTotalAmountTicketsByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the destination country from the query params
		country := chi.URLParam(r, "dest")
		// get the total amount of tickets by destination country
		total, err := h.sv.GetTicketsAmountByDestinationCountry(country)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		// write the response
		response.JSON(w, http.StatusOK, "succesfully fetched total", &TotalTicketsResponse{
			Total: total,
		})
	}
}

// PercentageTicketsResponse represents the response of the percentage of tickets
type PercentageTicketsResponse struct {
	Percentage float64 `json:"percentage"`
}

// GetPercentageTicketsByDestinationCountry returns the percentage of tickets by destination country
func (h *HandlerTicketDefault) GetPercentageTicketsByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the destination country from the query params
		country := chi.URLParam(r, "dest")
		// get the percentage of tickets by destination country
		percentage, err := h.sv.GetPercentageTicketsByDestinationCountry(country)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		// write the response
		response.JSON(w, http.StatusOK, "succesfully fetched total", &PercentageTicketsResponse{
			Percentage: percentage,
		})
	}
}
