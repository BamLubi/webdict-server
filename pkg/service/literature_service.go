package service

import (
	"context"
	"gorm.io/gorm"
	"time"
	"unicode"
	"webdict-server/pkg/config"
	"webdict-server/pkg/entity"
	"webdict-server/pkg/utils"
)

type LiteratureService struct {
	db *gorm.DB
}

func NewLiteratureService() *LiteratureService {
	return &LiteratureService{db: config.DB}
}

func (service LiteratureService) Insert(ctx context.Context, literature entity.Literature) bool {
	literature.Id = utils.IdGen()
	literature.CreateTime = time.Now().UTC()
	literature.UpdateTime = time.Now().UTC()
	service.db.Save(&literature)
	return true
}

func (service LiteratureService) FindAll(ctx context.Context) []entity.Literature {
	var literatures []entity.Literature
	service.db.Find(&literatures)
	return literatures
}

func (service LiteratureService) Update(ctx context.Context, literature entity.Literature) bool {
	literature.UpdateTime = time.Now().UTC()
	service.db.Model(&entity.Literature{}).Where("id = ?", literature.Id).Updates(
		entity.Literature{
			Eng:        literature.Eng,
			Zh:         literature.Zh,
			Rate:       literature.Rate,
			Num:        literature.Num,
			Available:  literature.Available,
			UpdateTime: literature.UpdateTime})
	return true
}

func (service LiteratureService) FuzzyFind(ctx context.Context, key string) []entity.Literature {
	var literatures []entity.Literature
	// 确定是否包含中文，如果包含就使用ZH判断
	var isChinese bool
	var queryCol string
	isChinese = false
	for _, v := range key {
		if unicode.Is(unicode.Han, v) {
			isChinese = true
			break
		}
	}
	if isChinese {
		queryCol = "zh"
	} else {
		queryCol = "eng"
	}
	service.db.Model(&entity.Literature{}).Where("available = '1' AND "+queryCol+" like ?", "%"+key+"%").Order("rate DESC").Limit(50).Find(&literatures)
	return literatures
}
