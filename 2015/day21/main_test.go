package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanWin(t *testing.T) {
	boss := &sprite{hp: 12, damage: 7, armor: 2}
	player := &sprite{hp: 8, damage: 5, armor: 5}

	assert.True(t, canWin(boss, player))

	boss = &sprite{hp: 12, damage: 7, armor: 2}
	player = &sprite{hp: 8, damage: 5, armor: 5}

	won, lost := player.turn(boss)
	assert.False(t, won)
	assert.False(t, lost)
	assert.Equal(t, 12-3, boss.hp)
	assert.Equal(t, 8-2, player.hp)
	won, lost = player.turn(boss)
	assert.False(t, won)
	assert.False(t, lost)
	assert.Equal(t, 9-3, boss.hp)
	assert.Equal(t, 6-2, player.hp)
}
