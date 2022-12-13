package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_compare(t *testing.T) {
	testSignals := parseJsons(testInput)

	assert.Equal(t, -1, compare(testSignals[0], testSignals[1]))
	assert.Equal(t, 1, compare(testSignals[1], testSignals[0]))

	assert.Equal(t, -1, compare(testSignals[2], testSignals[3]))
	assert.Equal(t, 1, compare(testSignals[3], testSignals[2]))

	assert.Equal(t, 1, compare(testSignals[4], testSignals[5]))
	assert.Equal(t, -1, compare(testSignals[5], testSignals[4]))

	assert.Equal(t, -1, compare(testSignals[6], testSignals[7]))
	assert.Equal(t, 1, compare(testSignals[7], testSignals[6]))

	assert.Equal(t, 1, compare(testSignals[8], testSignals[9]))
	assert.Equal(t, -1, compare(testSignals[9], testSignals[8]))

	assert.Equal(t, -1, compare(testSignals[10], testSignals[11]))
	assert.Equal(t, 1, compare(testSignals[11], testSignals[10]))

	assert.Equal(t, 1, compare(testSignals[12], testSignals[13]))
	assert.Equal(t, -1, compare(testSignals[13], testSignals[12]))

	assert.Equal(t, 1, compare(testSignals[14], testSignals[15]))
	assert.Equal(t, -1, compare(testSignals[15], testSignals[14]))
}
