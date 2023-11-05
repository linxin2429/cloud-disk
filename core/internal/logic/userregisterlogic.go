package logic

import (
	"context"
	"log"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"
	"cloud_disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil || code != req.Captcha {
		err = utils.ErrCaptchaNotMatch
		return
	}

	cnt, err := l.svcCtx.Engine.Where("name = ?", req.Name).Count(new(models.UserBasic))
	if err != nil {
		err = utils.NewErrWrapper(err, "UserRegister")
		return
	}
	if cnt > 0 {
		err = utils.ErrUserExist
		return
	}

	passwd, err := utils.HashPassword(req.Password)
	if err != nil {
		err = utils.NewErrWrapper(err, "UserRegister")
		return
	}

	user := &models.UserBasic{
		Name:     req.Name,
		Email:    req.Email,
		Identity: utils.GenerateUUID(),
		Password: passwd,
	}

	n, err := l.svcCtx.Engine.Insert(user)
	if err != nil {
		err = utils.NewErrWrapper(err, "UserRegister")
	}

	log.Printf("insert {%v}, row %d\n", user, n)
	return
}
