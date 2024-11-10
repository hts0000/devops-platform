package manager

import (
	"context"
	managerpb "manager/api/gen/v1"
	"manager/dao"
	"shared/server"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:123456@tcp(127.0.0.1:13306)/devops-platform-test?charset=utf8mb4&parseTime=True&loc=Local"

func TestCreateBizLine(t *testing.T) {
	s := newService(t)
	ctx := context.Background()

	res, err := s.CreateBizLine(ctx, &managerpb.BizLine{
		Name:          "hello1",
		ResponsibleId: 1,
		Description:   "create from testing",
	})
	if err != nil {
		t.Fatalf("create bizline failed, err: %v", err)
	}

	t.Logf("res: %#v\n", res)
}

func TestDeleteBizLine(t *testing.T) {
	s := newService(t)
	ctx := context.Background()

	res, err := s.DeleteBizLine(ctx, &managerpb.DeleteBizLineRequest{
		Id: 1,
	})
	if err != nil {
		t.Fatalf("delete bizline failed, err: %v", err)
	}

	t.Logf("res: %#v\n", res)
}

func TestUpdateBizLine(t *testing.T) {
	s := newService(t)
	ctx := context.Background()

	res, err := s.UpdateBizLine(ctx, &managerpb.BizLineEntity{
		Id: 2,
		BizLine: &managerpb.BizLine{
			Name:          "update",
			ResponsibleId: 99,
			Description:   "update from testing",
		},
	})
	if err != nil {
		t.Fatalf("update bizline failed, err: %v", err)
	}

	t.Logf("res: %#v\n", res)
}

func TestGetBizLine(t *testing.T) {
	s := newService(t)
	ctx := context.Background()

	res, err := s.GetBizLine(ctx, &managerpb.GetBizLineRequest{
		Id: 2,
	})
	if err != nil {
		t.Fatalf("get bizline failed, err: %v", err)
	}

	t.Logf("res: %#v\n", res)
}

func TestGetBizLines(t *testing.T) {
	s := newService(t)
	ctx := context.Background()

	res, err := s.GetBizLines(ctx, &managerpb.GetBizLinesRequest{
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		t.Fatalf("get bizlines failed, err: %v", err)
	}

	for _, biz := range res.Bizlines {
		t.Logf("id: %v, name: %v, rid: %v, desc: %v\n", biz.Id, biz.BizLine.Name, biz.BizLine.ResponsibleId, biz.BizLine.Description)
	}
}

func newService(t *testing.T) *BizLineService {
	logger, err := server.NewZapLogger()
	if err != nil {
		t.Fatalf("create logger failed, err: %v", err)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("open mysql database failed, err: %v", err)
	}

	return &BizLineService{
		Logger: logger,
		MySQL:  dao.NewMySQL(db),
	}
}
