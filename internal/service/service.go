package service

import (
	"github.com/google/wire"
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// RealWorldService is a greeter service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc *biz.RealWorldUsecase
}

// NewRealWorldService new a greeter service.
func NewRealWorldService(uc *biz.RealWorldUsecase) *RealWorldService {
	return &RealWorldService{uc: uc}
}
