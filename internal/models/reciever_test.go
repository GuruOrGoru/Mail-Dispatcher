package models

import (
	"encoding/json"
	"testing"
)

func TestReciever_JSON(t *testing.T) {
	r := Reciever{Name: "John Doe", Email: "john@example.com"}
	data, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{"name":"John Doe","email":"john@example.com"}`
	if string(data) != expected {
		t.Errorf("expected %s, got %s", expected, string(data))
	}

	var r2 Reciever
	err = json.Unmarshal(data, &r2)
	if err != nil {
		t.Fatal(err)
	}

	if r != r2 {
		t.Errorf("expected %+v, got %+v", r, r2)
	}
}
