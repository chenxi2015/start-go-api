package entity

import "github.com/voyager-go/start-go-api/entity/internal"

type SysUser internal.SysUser

const SysUserStatusTrue = 1  // 启用
const SysUserStatusFalse = 0 // 禁用

// SysUserServiceCreateReq 创新sys_user输入参数
type SysUserServiceCreateReq struct {
	Nickname string
	Phone    string
	Password string
	Status   int8
}
