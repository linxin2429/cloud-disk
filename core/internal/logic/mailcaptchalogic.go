package logic

import (
	"context"
	"time"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"
	"cloud_disk/utils"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
)

type MailCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCaptchaLogic {
	return &MailCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCaptchaLogic) MailCaptcha(req *types.MailCaptchaRequest) (resp *types.MailCaptchaResponse, err error) {
	cnt, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return
	}
	if cnt > 0 {
		err = utils.ErrEmailRegistered
		return
	}

	_, err = l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		if err != redis.Nil {
			return
		}
	} else {
		err = utils.ErrCaptchaRepeated
		return
	}

	code := utils.RandCode()
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(models.CaptchaExpiration))
	if err := utils.SendEmailCaptcha(req.Email, code); err != nil {
		return nil, utils.NewErrWrapper(err, "MailCaptcha")
	}
	return
}
