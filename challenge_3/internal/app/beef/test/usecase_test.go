package beef_test

import (
	"testing"

	"github.com/nithit-p/7-solutions-backend-challenge/challenge_3/internal/app/beef"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBeefUsecase(t *testing.T) {
	_, err := beef.NewBeefUsecase()
	if err != nil {
		t.Error("Error while creating new BeefUsecase")
	}
}

func TestGetSummary(t *testing.T) {
	tests := []struct {
		give        string
		wantSammary *beef.BeefSummary
	}{
		{
			give: `test
test
testtest		.,.,,	.testtest`,
			wantSammary: &beef.BeefSummary{
				Beef: map[string]int{
					"test":     2,
					"testtest": 2,
				},
			},
		},
		{
			give: "testtesttesttesttesttesttest.testtest",
			wantSammary: &beef.BeefSummary{
				Beef: map[string]int{
					"testtesttesttesttesttesttest": 1,
					"testtest":                     1,
				},
			},
		},
		{
			give: "Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone.",
			wantSammary: &beef.BeefSummary{
				Beef: map[string]int{
					"bresaola": 1,
					"enim":     1,
					"fatback":  1,
					"jowl":     1,
					"meatloaf": 1,
					"pastrami": 1,
					"pork":     1,
					"t-bone":   4,
				},
			},
		},
		{
			give: "Fatback t-bone chicken, pastrami  .chicken.,.   t-bone.  pork, meatloaf jowl ., ,.,.. chicken",
			wantSammary: &beef.BeefSummary{
				Beef: map[string]int{
					"chicken":  3,
					"fatback":  1,
					"jowl":     1,
					"meatloaf": 1,
					"pastrami": 1,
					"pork":     1,
					"t-bone":   2,
				},
			},
		},
		{
			give: " ..,. .  ,   ., ,.,.. ",
			wantSammary: &beef.BeefSummary{
				Beef: map[string]int{},
			},
		},
	}
	beefUsecase, _ := beef.NewBeefUsecase()

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			beefSummary, err := beefUsecase.GetSummary(tt.give)
			require.NoError(t, err)
			assert.Equal(t, tt.wantSammary, beefSummary)
		})
	}
}
