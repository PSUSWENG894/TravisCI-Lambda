package main

import "testing"

func TestSum(t *testing.T) {
	testSlice := []string{"Success","Failed","Error","Failed","Success","Success","Success"}
	total := Sum(testSlice)
	if total != 3 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSumTable(t *testing.T) {
	tables := []struct {
		test []string
		total int
	}{
		{[]string{"Success","Failed"}, 1},
		{[]string{"Success","Failed","Error"}, 2},
		{[]string{"Success","Failed","Error","Failed","Success","Success","Success"}, 3},
		{[]string{"Error","Failed","Failed","Failed","Failed","Error","Error","Error",
			"Success","Skipped","Ignored","Error","Failed","Success"}, 10},
	}

	for _, table := range tables {
		total := Sum(table.test)
		if total != table.total {
			t.Errorf("Sum of (%v) was incorrect, got: %d, want: %d.", table.test, total, table.total)
		}
	}
}