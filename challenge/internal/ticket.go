package internal

// TicketAttributes is an struct that represents a ticket
type TicketAttributes struct {
	// Name represents the name of the owner of the ticket
	Name string
	// Email represents the email of the owner of the ticket
	Email string
	// Country represents the destination country of the ticket
	Country string
	// Hour represents the hour of the ticket
	Hour string
	// Price represents the price of the ticket
	Price float64
}

// Ticket represents a ticket
type Ticket struct {
	// Id represents the id of the ticket
	Id int `json:"id"`
	// Attributes represents the attributes of the ticket
	Attributes TicketAttributes `json:"attributes"`
}
