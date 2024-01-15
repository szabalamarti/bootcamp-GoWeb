package service

import (
	"app/internal"
	"errors"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp internal.RepositoryTicket
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp internal.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	// get all the tickets
	t, err := s.rp.Get()
	if err != nil {
		return 0, errors.New("error getting the tickets")
	}

	// return the total number of tickets
	return len(t), nil
}

// GetTotalAmountTicketsByDestinationCountry returns the total amount of tickets by destination country
func (s *ServiceTicketDefault) GetTicketsAmountByDestinationCountry(country string) (total int, err error) {
	// get all the tickets
	t, err := s.rp.GetTicketsByDestinationCountry(country)
	if err != nil {
		return 0, errors.New("error getting the tickets")
	}

	// return the total number of tickets
	return len(t), nil
}

func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(country string) (percentage float64, err error) {
	// get all the tickets
	t, err := s.rp.Get()
	if err != nil {
		return 0, errors.New("error getting the tickets")
	}

	// get tickets by destination country
	tDest, err := s.rp.GetTicketsByDestinationCountry(country)
	if err != nil {
		return 0, errors.New("error getting the tickets")
	}

	// return the percentage of tickets by destination country
	if len(t) == 0 {
		return 0, errors.New("no tickets found")
	}

	return float64(len(tDest)) / float64(len(t)), nil
}
