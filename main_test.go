package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_displa(t *testing.T) {
	n := randomnubers()

	assert.GreaterOrEqual(t, n, 0)
	assert.LessOrEqual(t, n, 1000)
}
