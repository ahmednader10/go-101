package messages

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("gopher")
	expect := "Hello gopher\n"

	if got != expect {
		t.Errorf("didn't get what's expected, got %v, expected %v", got, expect)
	}

}

func TestGreetTableDriven(t *testing.T) {
	scenarios := []struct {
		input string
		expect string
	} {
		{input: "gopher", expect:"Hello gopher\n"},
		{input: "", expect:"Hello \n"},
	}
	for _,s := range scenarios {
		got := Greet(s.input)
		if got != s.expect {
			t.Errorf("didn't get what's expected, got %v, expected %v", got, s.expect)
		}
	}

}

func TestDepart(t *testing.T) {
	got := Depart("gopher")
	expect := "Goodbye gopher\n"

	if got != expect {
		t.Errorf("didn't get what's expected, got %v, expected %v", got, expect)
	}
}
