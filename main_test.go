package main

import (
	"gsd/hello"
	"testing"
)

func TestHello(t *testing.T) {
	// Arrange
	myName := "mark"

	// Act
	actual := hello.Hello(myName)

	// Assert
	expect := "hello, world, Eric"

	if actual != expect {
		t.Errorf("[actual][%s][expect][%s]", actual, expect)
	}

}
