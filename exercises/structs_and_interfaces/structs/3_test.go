package main

import (
	"reflect"
	"testing"
)

func TestAddItem(t *testing.T) {
	c := Cart{}
	c = AddItem(c, Product{Name: "Coffee", Price: 4.50})
	c = AddItem(c, Product{Name: "Tea", Price: 3.00})

	if len(c.Items) != 2 {
		t.Fatalf("cart has %d items; want 2", len(c.Items))
	}
	if c.Items[0].Name != "Coffee" || c.Items[0].Price != 4.50 {
		t.Errorf("Items[0] = %v; want Coffee/$4.50", c.Items[0])
	}
	if c.Items[1].Name != "Tea" || c.Items[1].Price != 3.00 {
		t.Errorf("Items[1] = %v; want Tea/$3.00", c.Items[1])
	}

	start := Cart{Items: []Product{{Name: "A", Price: 1}}}
	added := AddItem(start, Product{Name: "B", Price: 2})
	if !reflect.DeepEqual(start.Items, []Product{{Name: "A", Price: 1}}) {
		t.Errorf("AddItem must not mutate the input cart; got %v", start.Items)
	}
	if len(added.Items) != 2 {
		t.Errorf("added cart has %d items; want 2", len(added.Items))
	}
}