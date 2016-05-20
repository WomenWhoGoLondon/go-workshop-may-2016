package main

import "testing"

func TestLoadPage(t *testing.T) {
	page, err := loadPage("test")
	if err != nil {
		t.Fatal(err)
	}
	if page.Title != "test" {
		t.Fatal("Expected title page 'test'")
	}
}
