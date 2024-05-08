package beef_test

import (
	"github.com/nithit-p/7-solutions-backend-challenge/challenge_3/internal/app/beef"
	"github.com/nithit-p/7-solutions-backend-challenge/challenge_3/pb"

	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type beefUsercaseMock struct {
	GetSummaryReturn *beef.BeefSummary
	GetSummaryError  error
}

var _ beef.BeefService = (*beefUsercaseMock)(nil)

func (m *beefUsercaseMock) GetSummary(body string) (*beef.BeefSummary, error) {
	return m.GetSummaryReturn, m.GetSummaryError
}

func TestGetBeefSummary(t *testing.T) {
	beefGrpcHandler, err := beef.NewBeefGrpcHandler(&beefUsercaseMock{
		GetSummaryReturn: &beef.BeefSummary{
			Beef: map[string]int{
				"test": 2,
			},
		},
	})
	require.NoError(t, err)
	beefSummary, err := beefGrpcHandler.GetBeefSummary(context.Background(), &pb.BeefSummaryRequest{})
	require.NoError(t, err)
	assert.Equal(t, map[string]int32{
		"test": 2,
	}, beefSummary.Beef)
}

func TestGetBeefSummary2(t *testing.T) {
	beefGrpcHandler, err := beef.NewBeefGrpcHandler(&beefUsercaseMock{
		GetSummaryReturn: &beef.BeefSummary{
			Beef: map[string]int{
				"test":  2,
				"test2": 3,
				"test3": 4,
			},
		},
	})
	require.NoError(t, err)
	beefSummary, err := beefGrpcHandler.GetBeefSummary(context.Background(), &pb.BeefSummaryRequest{})
	require.NoError(t, err)
	assert.Equal(t, map[string]int32{
		"test":  2,
		"test2": 3,
		"test3": 4,
	}, beefSummary.Beef)
}

func TestGetBeefSummaryError(t *testing.T) {
	beefGrpcHandler, err := beef.NewBeefGrpcHandler(&beefUsercaseMock{
		GetSummaryError: assert.AnError,
	})
	require.NoError(t, err)
	_, err = beefGrpcHandler.GetBeefSummary(context.Background(), &pb.BeefSummaryRequest{})
	require.Error(t, err)
}
