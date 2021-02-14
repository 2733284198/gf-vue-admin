package model

import (
	"gorm.io/gorm"
	"time"
)

type Authority struct {
	CreatedAt time.Time      `orm:"created_at" json:"CreatedAt"`
	UpdatedAt time.Time      `orm:"updated_at" json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `orm:"deleted_at" json:"-" gorm:"index"`

	ParentId      string `json:"parentId" gorm:"comment:父角色ID"`
	AuthorityId   string `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	AuthorityName string `json:"authorityName" gorm:"comment:角色名"`
	DefaultRouter string `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"`

	Children      []Authority `json:"children" gorm:"-"`
	BaseMenus     []BaseMenu  `json:"menus" gorm:"many2many:authorities_menus;foreignKey:AuthorityId;joinForeignKey:AuthorityId;References:ID;JoinReferences:BaseMenuID"`
	DataAuthority []Authority `json:"dataAuthorityId" gorm:"many2many:data_authorities;foreignKey:AuthorityId;joinForeignKey:AuthorityId;References:AuthorityId;JoinReferences:DataAuthority"`
}

func (a *Authority) TableName() string {
	return "authorities"
}

