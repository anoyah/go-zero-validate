package logic

import (
	"context"

	"valid/internal/svc"
	"valid/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *user.CreateRequest) (*user.CreateResponse, error) {
	// TODO: 手动调用protoc-gen-validate验证
	// if err := in.ValidateAll(); err != nil {
	// 	return &user.CreateResponse{
	// 		Msg: err.Error(),
	// 	}, nil
	// }

	return &user.CreateResponse{
		Msg: "Ok",
	}, nil
}
