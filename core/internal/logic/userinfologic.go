package logic

import (
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"
	"cloud_disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	resp = new(types.UserInfoResponse)
	ub := new(models.UserBasic)
	has, err := models.Engine.Where("identity=?",req.Identity).Get(ub)
	if err!= nil {
		return nil, utils.NewErrWrapper(err, "UserInfo")
	}

	if !has {
		return nil, utils.ErrUserNotFound
	}

	resp.Name = ub.Name
	resp.Email = ub.Email
	return
}
