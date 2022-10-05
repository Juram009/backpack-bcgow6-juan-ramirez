package service

import "fmt"

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
	Reader() []Ticket
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}

}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	if t.Names == "" || t.Email == "" || t.Destination == "" || t.Date == "" || t.Price == 0 {
		return Ticket{}, fmt.Errorf("Todos los datos son necesarios")
	}
	t.Id = len(b.Tickets) + 1
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	if len(b.Tickets) >= id {
		return b.Tickets[id-1], nil
	} else {
		return Ticket{}, fmt.Errorf("Este id no existe en el booking")
	}
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	if t.Names == "" || t.Email == "" || t.Destination == "" || t.Date == "" || t.Price == 0 {
		return Ticket{}, fmt.Errorf("Todos los datos son necesarios")
	}
	if len(b.Tickets) >= id {
		if id != t.Id {
			return Ticket{}, fmt.Errorf("El id y el id del ticket no es el mismo")
		}
		b.Tickets[id-1] = t
		return t, nil
	} else {
		return Ticket{}, fmt.Errorf("Este id no existe en el booking")
	}
}

func (b *bookings) Delete(id int) (int, error) {
	if len(b.Tickets) < id {
		return 0, fmt.Errorf("Este id no existe en el booking")
	}
	b.Tickets = append(b.Tickets[:id-1], b.Tickets[id:]...)
	b.updateIds()
	return id, nil
}

func (b *bookings) updateIds() {
	newB := b.get()
	b.Tickets = append([]Ticket{})
	for i, value := range newB.Tickets {
		value.Id = i + 1
		b.Tickets = append(b.Tickets, value)
	}
}

func (b *bookings) Reader() []Ticket {
	return b.Tickets
}

func (b bookings) get() bookings {
	bNew := b
	return bNew
}
