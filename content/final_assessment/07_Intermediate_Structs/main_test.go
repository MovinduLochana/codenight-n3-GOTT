package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestBankAccountStructAndMethods(t *testing.T) {
	// 1. Test struct creation
	acc := BankAccount{
		Owner:   "Charlie",
		Balance: 200.0,
	}

	if acc.Owner != "Charlie" {
		t.Errorf("Expected Owner to be 'Charlie', got %q", acc.Owner)
	}
	if acc.Balance != 200.0 {
		t.Errorf("Expected Balance to be 200.0, got %f", acc.Balance)
	}

	// 2. Test Deposit (verifies pointer receiver mutation)
	acc.Deposit(100.25)
	if acc.Balance != 300.25 {
		t.Errorf("Expected Balance to be 300.25 after Deposit, got %f. Ensure Deposit is implemented with a pointer receiver.", acc.Balance)
	}

	// 3. Test GetDetails (value receiver)
	details := acc.GetDetails()
	expectedDetails := "Account: Charlie, Balance: $300.25"
	if details != expectedDetails {
		t.Errorf("GetDetails() = %q; want %q", details, expectedDetails)
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

	expected := "Account: Alice, Balance: $150.50"
	if output != expected {
		t.Errorf("Expected main output %q, got %q", expected, output)
	}
}
