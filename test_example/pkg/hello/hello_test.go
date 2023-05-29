package hello_test

import (
	"test_example/pkg/hello"
	"testing"
)

func TestHello(t *testing.T) {
	got := hello.SayHello("Dela")
	want := "Hello Dela!"
	if got != want {
		t.Errorf(`Hello("Dela") = "%s"; want "%s"`, got, want)
	}
}
