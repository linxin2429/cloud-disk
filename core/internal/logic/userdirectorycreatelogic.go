package logic

import (
	"context"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"
	"cloud_disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDirectoryCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDirectoryCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDirectoryCreateLogic {
	return &UserDirectoryCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDirectoryCreateLogic) UserDirectoryCreate(req *types.UserDirectoryCreateRequest, userIdentity string) (resp *types.UserDirectoryCreateResponse, err error) {
	cnt, err := l.svcCtx.Engine.Where("name = ? AND parent_id = ?", req.Name, req.ParentId).Count(new(models.UserRepository))
	if err != nil {
		return
	}
	if cnt > 0 {
		return nil, utils.ErrFileNameExist
	}
	data := &models.UserRepository{
		Identity:     utils.GenerateUUID(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
	}

	_, err = l.svcCtx.Engine.Insert(data)
	return
}
