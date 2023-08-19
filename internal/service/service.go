package service

import (
	"github.com/google/wire"
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

// 服务端也用wire去管理

// service --> biz --> data
// 全部依赖由wire去管理

// controller -- service --- dal

// 实现了 api 定义的服务层，类似 DDD 的 application 层
// todo 唯一一个不清楚的点就是不能生成proto文件对应的实现方法

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
