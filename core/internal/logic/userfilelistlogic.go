package logic

import (
	"context"
	"time"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"cloud_disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListResponse, err error) {
	uf := make([]*types.UserFile, 0)
	resp = new(types.UserFileListResponse)

	size := req.Size
	if size == 0 {
		size = models.DefaultPageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Select("user_repository.id,user_repository.identity,user_repository.repository_identity,user_repository.ext,user_repository.name,"+
			"repository_pool.path,repository_pool.size").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.delete_at = ? OR user_repository.delete_at is NULL",time.Time{}.Format("2006-01-02 15:04:05")).
		Limit(size, offset).
		Find(&uf)
	if err != nil {
		return
	}
	resp.List = uf
	resp.Count = len(uf)
	return
}
