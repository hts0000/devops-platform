package dao

import (
	"shared/model"
)

func (m *MySQL) CreateBizLine(b *model.BizLine) (uint, error) {
	err := m.db.Create(b).Error
	return b.ID, err
}

func (m *MySQL) DeleteBizLine(bid uint) (*model.BizLine, error) {
	b := &model.BizLine{}
	err := m.db.Delete(b, bid).Error
	return b, err
}

func (m *MySQL) UpdateBizLine(b *model.BizLine) error {
	return m.db.Save(b).Error
}

func (m *MySQL) GetBizLine(bid uint) (*model.BizLine, error) {
	b := &model.BizLine{}
	err := m.db.Take(b, bid).Error
	return b, err
}

func (m *MySQL) GetBizLines(offset, limit uint) ([]*model.BizLine, error) {

}
