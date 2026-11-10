package books

import (
	"maps"
	"testing"
)

func TestNewBook(t *testing.T) {
	got := NewBook("Dune", "Herbert", 412)
	want := Book{Title: "Dune", Author: "Herbert", Pages: 412, Read: false}
	if got != want {
		t.Errorf("NewBook(%q, %q, 412) = %+v, want %+v", "Dune", "Herbert", got, want)
	}
}

func TestTotalPages(t *testing.T) {
	books := []Book{
		{Title: "Dune", Author: "Herbert", Pages: 412},
		{Title: "Emma", Author: "Austen", Pages: 474},
	}
	if got := TotalPages(books); got != 886 {
		t.Errorf("TotalPages(Dune, Emma) = %d, want 886", got)
	}
	if got := TotalPages([]Book{}); got != 0 {
		t.Errorf("TotalPages([]) = %d, want 0", got)
	}
}

func TestCountByAuthor(t *testing.T) {
	books := []Book{
		{Title: "Dune", Author: "Herbert", Pages: 412},
		{Title: "Emma", Author: "Austen", Pages: 474},
		{Title: "Dune Messiah", Author: "Herbert", Pages: 256},
	}
	want := map[string]int{"Herbert": 2, "Austen": 1}
	if got := CountByAuthor(books); !maps.Equal(got, want) {
		t.Errorf("CountByAuthor(Dune, Emma, Dune Messiah) = %v, want %v", got, want)
	}
	if got := CountByAuthor([]Book{}); len(got) != 0 {
		t.Errorf("CountByAuthor([]) = %v, want map[]", got)
	}
}

func TestMarkRead(t *testing.T) {
	dune := Book{Title: "Dune", Author: "Herbert", Pages: 412}
	got := MarkRead(dune)
	want := Book{Title: "Dune", Author: "Herbert", Pages: 412, Read: true}
	if got != want {
		t.Errorf("MarkRead(dune) = %+v, want %+v", got, want)
	}
	if dune.Read != false {
		t.Errorf("after MarkRead, the original dune changed: %+v (it must stay unread)", dune)
	}
}
