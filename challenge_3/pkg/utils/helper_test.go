package utils_test

import (
	"fmt"
	"testing"

	"github.com/nithit-p/7-solutions-backend-challenge/challenge_3/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func TestContainsRune(t *testing.T) {
	tests := []struct {
		giveString string
		giveRune   rune
		wantBool   bool
	}{
		{
			giveString: "test",
			giveRune:   't',
			wantBool:   true,
		},
		{
			giveString: "test",
			giveRune:   'a',
			wantBool:   false,
		},
		{
			giveString: "\n\r\t ",
			giveRune:   rune("\n"[0]),
			wantBool:   true,
		},
		{
			giveString: ".,r",
			giveRune:   '.',
			wantBool:   true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("(%s,%s)", tt.giveString, string(tt.giveRune)), func(t *testing.T) {
			result := utils.ContainsRune(tt.giveString, tt.giveRune)
			assert.Equal(t, tt.wantBool, result)
		})
	}
}
