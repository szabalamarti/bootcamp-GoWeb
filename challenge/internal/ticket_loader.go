package internal

type TicketLoader interface {
	Load() (t map[int]TicketAttributes, err error)
}
