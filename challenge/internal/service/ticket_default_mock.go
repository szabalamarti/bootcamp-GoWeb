package service

import "github.com/stretchr/testify/mock"

type ServiceTicketDefaultMock struct {
	mock.Mock
}

func (m *ServiceTicketDefaultMock) GetTotalAmountTickets() (total int, err error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *ServiceTicketDefaultMock) GetTicketsAmountByDestinationCountry(country string) (total int, err error) {
	args := m.Called(country)
	return args.Int(0), args.Error(1)
}

func (m *ServiceTicketDefaultMock) GetPercentageTicketsByDestinationCountry(country string) (percentage float64, err error) {
	args := m.Called(country)
	if arg, ok := args.Get(0).(float64); ok {
		percentage = arg
	}
	return percentage, args.Error(1)
}
