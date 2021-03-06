package models

import (
	"github.com/hzwy23/auth/dto"
	"github.com/hzwy23/auth/entity"
	"github.com/hzwy23/dbobj"
	"github.com/hzwy23/panda"
	)

type UserDomainPrivilege struct {
}

func (this *UserDomainPrivilege) GetByUserId(userId string) ([]dto.AccessDomainDTO, error) {
	var rst []dto.AccessDomainDTO
	if panda.IsAdmin(userId) {
		err := dbobj.QueryForSlice(sys_rdbms_053, &rst)
		return rst, err
	}
	err := dbobj.QueryForSlice(sys_rdbms_052, &rst, userId)
	return rst, err
}

func (this *UserDomainPrivilege) Get(userId string, domainId string) (entity.UserDomainData, error) {
	var row entity.UserDomainData
	err := dbobj.QueryForStruct(sys_rdbms_049, &row, userId, domainId)
	return row, err
}
