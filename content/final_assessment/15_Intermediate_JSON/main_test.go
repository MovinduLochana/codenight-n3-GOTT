package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestProductJSONTagsAndSerialization(t *testing.T) {
	// 1. Verify fields and tags using reflection
	typ := reflect.TypeOf(Product{})
	fields := []struct {
		name    string
		jsonTag string
	}{
		{"ID", "id"},
		{"Name", "name"},
		{"Price", "price"},
		{"InStock", "in_stock"},
	}

	for _, f := range fields {
		field, ok := typ.FieldByName(f.name)
		if !ok {
			t.Fatalf("Product struct is missing field %q", f.name)
		}
		tag := field.Tag.Get("json")
		if tag != f.jsonTag {
			t.Errorf("Field %q expected JSON tag %q, got %q", f.name, f.jsonTag, tag)
		}
	}

	// 2. Test serialization
	p := Product{ID: 101, Name: "Wireless Mouse", Price: 29.99, InStock: true}
	jsonStr, err := SerializeProduct(p)
	if err != nil {
		t.Fatalf("SerializeProduct failed: %v", err)
	}

	// Unmarshal back to check correctness
	var got Product
	if err := json.Unmarshal([]byte(jsonStr), &got); err != nil {
		t.Fatalf("Failed to unmarshal serialized product: %v", err)
	}

	if got != p {
		t.Errorf("Serialization roundtrip failed. Got %+v, want %+v", got, p)
	}
}

func TestMainOutput(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := strings.TrimSpace(buf.String())

	expected := `{"id":101,"name":"Wireless Mouse","price":29.99,"in_stock":true}`
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
