package beef

import (
	"strings"

	"github.com/nithit-p/7-solutions-backend-challenge/challenge_3/pkg/utils"
)

type BeefService interface {
	GetSummary(text string) (*BeefSummary, error)
}

type BeefUsecase struct{}

var _ BeefService = (*BeefUsecase)(nil)

func NewBeefUsecase() (BeefService, error) {
	return &BeefUsecase{}, nil
}

func (b *BeefUsecase) GetSummary(text string) (*BeefSummary, error) {
	if text == "" {
		return &BeefSummary{
			Beef: make(map[string]int),
		}, nil
	}
	summary := make(map[string]int)
	var beef string
	for _, char := range text {
		if utils.ContainsRune("., \r\n\t", char) {
			if beef == "" {
				continue
			}
			summary[strings.ToLower(beef)]++
			beef = ""
			continue
		}
		beef += string(char)
	}
	if beef != "" {
		summary[strings.ToLower(beef)]++
	}
	return &BeefSummary{
		Beef: summary,
	}, nil
}
