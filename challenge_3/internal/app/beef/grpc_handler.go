package beef

import (
	"context"

	"github.com/nithit-p/7-solutions-backend-challenge/challenge_3/pb"
)

type BeefGrpcHandler struct {
	beefService BeefService
	pb.UnimplementedBeefServiceServer
}

var _ pb.BeefServiceServer = (*BeefGrpcHandler)(nil)

func NewBeefGrpcHandler(beefService BeefService) (*BeefGrpcHandler, error) {
	return &BeefGrpcHandler{
		beefService: beefService,
	}, nil
}

func (h *BeefGrpcHandler) GetBeefSummary(ctx context.Context, input *pb.BeefSummaryRequest) (*pb.BeefSummary, error) {
	beefSummary, err := h.beefService.GetSummary(input.Text)
	if err != nil {
		return nil, err
	}
	newMap := make(map[string]int32, len(beefSummary.Beef))
	for key, value := range beefSummary.Beef {
		newMap[key] = int32(value)
	}
	resp := pb.BeefSummary{
		Beef: newMap,
	}
	return &resp, nil
}
