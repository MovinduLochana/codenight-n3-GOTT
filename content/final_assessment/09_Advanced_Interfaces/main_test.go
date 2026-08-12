package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestItemInterfaceAndStructs(t *testing.T) {
	// 1. Test BaseItem implements Item
	var _ Item = BaseItem{}
	bi := BaseItem{Name: "Book", BasePrice: 10.0}
	if bi.Price() != 10.0 {
		t.Errorf("BaseItem Price() expected 10.0, got %f", bi.Price())
	}

	// 2. Test TaxableItem embeds BaseItem and overrides Price()
	var _ Item = TaxableItem{}
	ti := TaxableItem{
		BaseItem: BaseItem{Name: "Gadget", BasePrice: 100.0},
	}
	// Expected Price: 100.0 * 1.15 = 115.0
	if ti.Price() != 115.0 {
		t.Errorf("TaxableItem Price() expected 115.0, got %f", ti.Price())
	}

	// 3. Test CalculateTotal
	items := []Item{
		BaseItem{Name: "Normal Item", BasePrice: 50.0},      // 50.0
		TaxableItem{BaseItem: BaseItem{BasePrice: 200.0}},    // 230.0
		BaseItem{Name: "Another Normal", BasePrice: 20.0},    // 20.0
	}
	total := CalculateTotal(items)
	expectedTotal := 300.0
	if total != expectedTotal {
		t.Errorf("CalculateTotal() expected %f, got %f", expectedTotal, total)
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

	expected := "Total Cart Price: $1180.00"
	if output != expected {
		t.Errorf("Expected main output %q, got %q", expected, output)
	}
}
