package ent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHangman_NewHangman(t *testing.T) {
	// arrange
	expectedWord := Word("test1")
	expectedState := Word("_____")
	// act
	sut := NewHangman(expectedWord)
	// assert
	actualWord := Word(sut.Word)
	assert.Equal(t, expectedWord, actualWord)
	actualState := Word(sut.State)
	assert.Equal(t, expectedState, actualState)
	assert.Equal(t, 8, sut.TriesCnt)
	assert.False(t, sut.Won)
}
