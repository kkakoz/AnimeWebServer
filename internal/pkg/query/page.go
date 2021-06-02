package query

import "gorm.io/gorm"

type Page struct {
	Page     int
	PageSize int
}

func (p *Page) LimitPage(db *gorm.DB) *gorm.DB {
	p.Page -= 1
	if p.Page < 0 { // 如果小于1
		p.Page = 0
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	return db.Limit(p.Page).Offset(p.PageSize * p.Page)
}
