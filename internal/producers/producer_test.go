package producers

import (
	"os"
	"testing"

	"github.com/guruorgoru/email-dispatcher/internal/models"
)

func TestLoadRecievers(t *testing.T) {
	csvContent := `name,email
John Doe,john@example.com
Jane Smith,jane@example.com`

	tmpFile, err := os.CreateTemp("", "test_*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(csvContent); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	ch := make(chan models.Reciever, 10)
	err = LoadRecievers(tmpFile.Name(), ch)
	if err != nil {
		t.Fatal(err)
	}

	var recievers []models.Reciever
	for r := range ch {
		recievers = append(recievers, r)
	}

	expected := []models.Reciever{
		{Name: "John Doe", Email: "john@example.com"},
		{Name: "Jane Smith", Email: "jane@example.com"},
	}

	if len(recievers) != len(expected) {
		t.Fatalf("expected %d recievers, got %d", len(expected), len(recievers))
	}

	for i, r := range recievers {
		if r != expected[i] {
			t.Errorf("expected %+v, got %+v", expected[i], r)
		}
	}
}

func TestLoadRecievers_InvalidFile(t *testing.T) {
	ch := make(chan models.Reciever, 10)
	err := LoadRecievers("nonexistent.csv", ch)
	if err == nil {
		t.Error("expected error for nonexistent file")
	}
}
