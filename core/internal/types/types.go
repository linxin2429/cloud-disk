// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserInfoRequest struct {
	Identity string `json:"identity"`
}

type UserInfoResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MailCaptchaRequest struct {
	Email string `json:"email"`
}

type MailCaptchaResponse struct {
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Captcha  string `json:"captcha"`
}

type UserRegisterResponse struct {
}

type FileUploadRequest struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext  string `json:"ext,optional"`
	Size int64  `json:"size,optional"`
	Path string `json:"path,optional"`
}

type FileUploadResponse struct {
	Identity string `json:"identity"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
}

type UserRepositorySaveRequest struct {
	ParentId           int    `json:"parent_id"`
	RepositoryIdentity string `json:"repository_identity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}

type UserRepositorySaveResponse struct {
	Identity string `json:"identity"`
}

type UserFileListRequest struct {
	Id   int `json:"id,optional"`
	Page int `json:"page,optional"`
	Size int `json:"size,optional"`
}

type UserFileListResponse struct {
	List  []*UserFile `json:"list"`
	Count int         `json:"count"`
}

type UserFile struct {
	Id                 int    `json:"id"`
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repository_identity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
	Path               string `json:"path"`
	Size               int64  `json:"size"`
}

type UserFileRenameRequest struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
}

type UserFileRenameResponse struct {
}

type UserDirectoryCreateRequest struct {
	ParentId int    `json:"parent_id"`
	Name     string `json:"name"`
}

type UserDirectoryCreateResponse struct {
	Identity string `json:"identity"`
}

type UserFileDeleteRequest struct {
	Identity string `json:"identity"`
}

type UserFileDeleteResponse struct {
}

type UserFileMoveRequest struct {
	Identity string `json:"identity"`
	ParentId int    `json:"parent_id"`
}

type UserFileMoveResponse struct {
}
