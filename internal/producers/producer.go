package producers

import (
	"encoding/csv"
	"os"

	"github.com/guruorgoru/email-dispatcher/internal/models"
)

func LoadRecievers(filepath string, recieverChannel chan models.Reciever) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		recieverChannel <- models.Reciever{
			Name:  record[0],
			Email: record[1],
		}
	}

	close(recieverChannel)

	return nil
}
