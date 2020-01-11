package geoPost

import (
	"github.com/ldugdale/dropper/pkg/log"
	pb "github.com/LDugdale/Dropper/proto"
	"golang.org/x/net/context"
)

type geoPostRepository interface {
	AddGeoPost()
}

type GeoPostService struct {
	logger log.Logger
	repository geoPostRepository
}

func NewGeoPostService(logger log.Logger) GeoPostService {
	return GeoPostService{
		logger: logger,
	}
}

func (geoPostService GeoPostService) AddGeoPost(context context.Context, params *pb.AddGeoPostParameters) (*pb.AddGeoPostReturn, error){
	
	geoPostService.logger.LogTrace("AddGeoPost")

	values := &pb.AddGeoPostReturn {
		Version: "test",
	}

	return values, nil
}