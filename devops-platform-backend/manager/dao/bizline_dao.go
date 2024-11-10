package dao

import (
	"context"
	"shared/model"
)

func (m *MySQL) CreateBizLine(ctx context.Context, b *model.BizLine) error {
	tx := m.db.WithContext(ctx)
	return tx.Create(b).Error
}

func (m *MySQL) DeleteBizLine(ctx context.Context, bid uint) (*model.BizLine, error) {
	tx := m.db.WithContext(ctx)
	b := &model.BizLine{}
	err := tx.Delete(b, bid).Error
	return b, err
}

func (m *MySQL) UpdateBizLine(ctx context.Context, b *model.BizLine) error {
	tx := m.db.WithContext(ctx)
	return tx.Save(b).Error
}

func (m *MySQL) GetBizLineByID(ctx context.Context, bid uint) (*model.BizLine, error) {
	tx := m.db.WithContext(ctx)
	b := &model.BizLine{}
	err := tx.Take(b, bid).Error
	return b, err
}

func (m *MySQL) GetBizlineByName(ctx context.Context, name string) (*model.BizLine, error) {
	tx := m.db.WithContext(ctx)
	b := &model.BizLine{}
	err := tx.Where("business_line_name = ?", name).First(b).Error
	return b, err
}

func (m *MySQL) GetBizLines(ctx context.Context, page, pageSize int) ([]*model.BizLine, int64, error) {
	tx := m.db.WithContext(ctx)
	bizs := make([]*model.BizLine, 0, pageSize)
	page = (page - 1) * pageSize

	var total int64
	tx.Model(&model.BizLine{}).Count(&total)
	err := tx.Offset(page).Limit(pageSize).Find(&bizs).Error

	return bizs, total, err
}
