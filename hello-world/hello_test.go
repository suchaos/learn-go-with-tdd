package hello_world

import "testing"

func TestHello(t *testing.T) {
	result := Hello()
	excepted := "hello, world"

	if result != excepted {
		t.Errorf("excepted %s, but got %s", excepted, result)
	}
}
