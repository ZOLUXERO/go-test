package main

import (
	"testing"

	"github.com/ZOLUXERO/go-test/internal/config"
	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
	default:
		t.Error("Type is not *chi.Mux", v)
	}
}
