package service

import (
	"errors"
	"gorm.io/gorm"
)

type Service struct {
	Orm      *gorm.DB
	Identity *AuthIdentity
}

var (
	ErrDatabaseError = errors.New("db error")

	ErrInternalError = errors.New("internal error")

	ErrNoSuchObject = errors.New("对象不存在")

	ErrPermissionDenied = errors.New("对象不存在或无权查看")

	NoAuthIdentity = AuthIdentity{
		UserId:   -2,
		Username: "noauth",
		DeptId:   -2,
		RoleId:   -2,
		RoleKey:  "noauth",
		RoleName: "noauth",
	}

	WebIdentity = AuthIdentity{
		UserId:   -1,
		Username: "web",
		DeptId:   -1,
		RoleId:   -1,
		RoleKey:  "web",
		RoleName: "web",
	}
)

func (s *Service) ConstructFromDB(db *gorm.DB) {
	s.Orm = db
	s.Identity = &WebIdentity
}
