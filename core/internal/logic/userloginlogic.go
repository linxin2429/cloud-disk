package logic

import (
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"
	"cloud_disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	user := new(models.UserBasic)
	has, err := l.svcCtx.Engine.Where("name = ?", req.Name).Get(user)
	if err != nil {
		return nil, utils.NewErrWrapper(err, "LoginLogic.Login")
	}

	if !has || !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, utils.ErrUsernameOrPasswd
	}

	token, err := utils.GenerateToken(user.Id, user.Identity, user.Name)
	if err != nil {
		return nil, utils.NewErrWrapper(err, "LoginLogic.Login")
	}
	resp = new(types.LoginResponse)
	resp.Token = token
	return
}
