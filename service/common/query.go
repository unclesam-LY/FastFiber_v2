package common

import (
	"FastFiber_v2/global"
	"FastFiber_v2/models"
	"fmt"
	"gorm.io/gorm"
)

type QueryOptions struct {
	models.PageInfo
	Likes    []string
	Where    *gorm.DB
	Preloads []string
	Debug    bool
}

func QueryList[T any](model T, option QueryOptions) (list []T, count int64, err error) {
	list = make([]T, 0)

	query := global.DB.Where(model)

	if option.Where != nil {
		query = query.Where(option.Where)
	}

	// 模糊匹配
	if option.Key != "" {
		if len(option.Likes) != 0 {
			likeQuery := global.DB.Where("")
			for _, column := range option.Likes {
				likeQuery.Or(
					fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			}
			query.Where(likeQuery)

		}

	}
	// 预加载
	for _, preload := range option.Preloads {
		query = query.Preload(preload)
	}

	// 	分页
	if option.Page <= 0 {
		option.Page = 1
	}

	if option.Limit <= 0 {
		option.Limit = -1
	}

	offset := (option.Page - 1) * option.Limit

	if option.Order == "" {
		option.Order = "created_at desc"
	}

	if option.Debug {
		query = query.Debug()
	}

	query.Limit(option.Limit).Offset(offset).Order(option.Order).Find(&list)

	global.DB.Model(model).Where(query).Count(&count)
	return
}
