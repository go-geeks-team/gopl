package movie

import (
	"encoding/json"
	"testing"
)

var data []byte

func TestMarshal(t *testing.T) {
	var err error
	data, err = json.Marshal(movies)
	if err != nil {
		t.Errorf("Сбой маршалинга JSON: %s", err)
	}
	t.Logf("%s\n", data)
}
func TestUnmarshal(t *testing.T) {
	var movies []Movie
	err := json.Unmarshal(data, &movies)
	if err != nil {
		t.Errorf("Сбой демаршалинга JSON: %s", err)
	}
	for _, m := range movies {
		t.Logf("title: %s actors: %s color: %t year: %d", m.Title, m.Actors, m.Color, m.Year)
	}
}
