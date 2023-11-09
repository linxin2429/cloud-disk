package logic

import (
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"
	"cloud_disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveResponse, err error) {
	parentData := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Where("id = ? AND user_identity = ?", req.ParentId, userIdentity).Get(parentData)
	if err != nil {
		return
	}
	if !has {
		err = utils.ErrDirNotFound
		return
	}
	_, err = l.svcCtx.Engine.Where("identity = ?", req.Identity).Update(models.UserRepository{
		ParentId: parentData.Id,
	})
	return
}
