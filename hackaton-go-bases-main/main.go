package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	var tickets []service.Ticket
	var ticket service.Ticket
	f := &file.File{Path: "./tickets.csv"}
	tickets, err := f.Read()
	if err != nil {
		panic("Error leyendo el archivo")
	}
	bkng := service.NewBookings(tickets)
	ticket, err = bkng.Read(1)
	ticket, err = bkng.Update(5, service.Ticket{
		Id:          5,
		Names:       "a",
		Email:       "a@a.a",
		Destination: "awa",
		Date:        "13:45",
		Price:       3,
	})
	ticket, err = bkng.Read(5)
	ticket, err = bkng.Create(service.Ticket{
		Names:       "a",
		Email:       "a@a.a",
		Destination: "awa",
		Date:        "13:45",
		Price:       3,
	})
	id, err := bkng.Delete(5)
	ticket, err = bkng.Read(id)
	fmt.Println(ticket, err)
	f.Write(bkng.Reader())
}
