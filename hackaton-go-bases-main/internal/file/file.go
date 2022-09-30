package file

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
	file, err := os.Open(f.Path)
	var tickets []service.Ticket
	if err != nil {
		return nil, fmt.Errorf("Archivo no encontrado")
	}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		splittedLine := strings.Split(line, ",")
		if len(splittedLine) != 6 {
			return nil, fmt.Errorf("Archivo con valores invalidos")
		}
		id, err := strconv.Atoi(splittedLine[0])
		if err != nil {
			return nil, fmt.Errorf("Archivo con valores de id invalidos")
		}
		price, err := strconv.Atoi(splittedLine[5])
		if err != nil {
			return nil, fmt.Errorf("Archivo con valores de precio invalidos")
		}
		tickets = append(tickets, service.Ticket{
			Id:          id,
			Names:       splittedLine[1],
			Email:       splittedLine[2],
			Destination: splittedLine[3],
			Date:        splittedLine[4],
			Price:       price,
		})
	}
	file.Close()
	return tickets, nil
}

func (f *File) Write(ts []service.Ticket) error {
	var txt string
	for i := 0; i < len(ts); i++ {
		txt += fmt.Sprintf("%d,%s,%s,%s,%s,%d\n", ts[i].Id, ts[i].Names, ts[i].Email, ts[i].Destination, ts[i].Date, ts[i].Price)
	}
	e := os.WriteFile(
		f.Path,
		[]byte(txt),
		0644,
	)
	if e != nil {
		return fmt.Errorf("Error escribiendo el archivo")
	}
	return nil
}
