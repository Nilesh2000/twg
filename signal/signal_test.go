package signal

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Fatalf("http.NewRequest(GET, \"\", nil) err = %s", err)
	}
	Handler(w, r)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fatalf("Handler() status = %d; want %d", resp.StatusCode, 200)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Handler() Content-Type = %q; want %q", contentType, "application/json")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll(resp.Body) err=%s", err)
	}

	var p Person
	err = json.Unmarshal(data, &p)
	if err != nil {
		t.Fatalf("json.Unmarshal(resp.Body) err = %s", err)
	}

	wantAge := 21
	if p.Age != wantAge {
		t.Errorf("person.Age = %d; want %d", p.Age, wantAge)
	}

	wantName := "Jon Calhoun"
	if p.Name != wantName {
		t.Errorf("person.Name = %s; want %s", p.Name, wantName)
	}

	if err != nil {
		t.Fatalf("SomeFunc(name=%s, age=%d)", p.Name, p.Age)
	}
}
