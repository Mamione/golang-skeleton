package account

import "testing"

func TestCreateAccount(t *testing.T) {
	got := create("mamionga", "123")
	want := true

	if got != want {
			t.Errorf("got %t, wanted %t", got, want)
	}
}