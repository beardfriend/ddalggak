package product

import (
	"context"
	"net/http"

	"github.com/beardfriend/ddalggak/ent"
	"github.com/beardfriend/ddalggak/internal/common"
	"github.com/beardfriend/ddalggak/pkg/pagination"
)

type Usecase interface {
	Create(ctx context.Context, b *ent.Product) (err error)
	Get(ctx context.Context, id int) (result *ent.Product, err error)
	Update(ctx context.Context, b *ent.Product) (err error)
	Delete(ctx context.Context, id int) (err error)
	List(ctx context.Context, q *common.ListRequest) (result []*ent.Product, pgInfo *pagination.PaginationInfo, err error)
	GetByUserID(ctx context.Context, id int, userID int) (result *ent.Product, err error)
	DeleteOneByUserID(ctx context.Context, id int, userID int) (err error)
	UpdateOneByUserID(ctx context.Context, b *ent.Product) (result *ent.Product, err error)
	ListByUserID(ctx context.Context, userID int, q *common.ListRequest) (result []*ent.Product, pgInfo *pagination.PaginationInfo, err error)
}

type usecase struct {
	repo Repo
}

func NewUsecase(repo Repo) Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) Create(ctx context.Context, b *ent.Product) (err error) {
	err = u.repo.Create(ctx, b)
	if err != nil {
		err = common.NewUsecaseError(http.StatusInternalServerError, err, common.ErrDatabaseError)
		return
	}
	return
}

func (u *usecase) Get(ctx context.Context, id int) (result *ent.Product, err error) {
	result, err = u.repo.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			err = common.NewUsecaseError(http.StatusNotFound, err, common.ErrNotfound)
			return
		}
		err = common.NewUsecaseError(http.StatusInternalServerError, err, common.ErrDatabaseError)
		return
	}
	return
}

func (u *usecase) Update(ctx context.Context, b *ent.Product) (err error) {
	err = u.repo.Update(ctx, b)
	if err != nil {
		err = common.NewUsecaseError(http.StatusInternalServerError, err, common.ErrDatabaseError)
		return
	}
	return
}

func (u *usecase) Delete(ctx context.Context, id int) (err error) {
	err = u.repo.Delete(ctx, id)
	if err != nil {
		err = common.NewUsecaseError(http.StatusInternalServerError, err, common.ErrDatabaseError)
		return
	}
	return
}

func (u *usecase) List(ctx context.Context, q *common.ListRequest) (result []*ent.Product, pgInfo *pagination.PaginationInfo, err error) {
	total, err := u.repo.Total(ctx)
	if err != nil {
		err = common.NewUsecaseError(http.StatusInternalServerError, err, common.ErrDatabaseError)
		return
	}

	pg := pagination.NewPagination(q.PageSize, q.PageNo)
	pg.SetTotal(total)

	result, err = u.repo.List(ctx, &ListParams{
		Limit:          pg.GetLimit(),
		Offset:         pg.GetOffset(),
		OrderFieldName: q.OrderFieldName,
		OrderIsDesc:    q.IsDesc,
	})
	if err != nil {
		if ent.IsNotFound(err) {
			err = common.NewUsecaseError(http.StatusNotFound, err, common.ErrNotfound)
			return
		}
		err = common.NewUsecaseError(http.StatusInternalServerError, err, common.ErrDatabaseError)
		return
	}
	pgInfo = pg.GetInfo(len(result))
	return
}

func (u *usecase) GetByUserID(ctx context.Context, id int, userID int) (result *ent.Product, err error) {
	result, err = u.repo.GetByUserID(ctx, id, userID)
	if err != nil {
		err = common.NewUsecaseError(http.StatusInternalServerError, err, common.ErrDatabaseError)
		return
	}
	return
}

func (u *usecase) DeleteOneByUserID(ctx context.Context, id int, userID int) (err error) {
	err = u.repo.DeleteOneByUserID(ctx, id, userID)
	if err != nil {
		err = common.NewUsecaseError(http.StatusInternalServerError, err, common.ErrDatabaseError)
		return
	}
	return
}

func (u *usecase) UpdateOneByUserID(ctx context.Context, b *ent.Product) (result *ent.Product, err error) {
	err = u.repo.UpdateOneByUserID(ctx, b)
	if err != nil {
		err = common.NewUsecaseError(http.StatusInternalServerError, err, common.ErrDatabaseError)
		return
	}
	return
}

func (u *usecase) ListByUserID(ctx context.Context, userID int, q *common.ListRequest) (result []*ent.Product, pgInfo *pagination.PaginationInfo, err error) {
	total, err := u.repo.TotalByUserID(ctx, userID)
	if err != nil {
		err = common.NewUsecaseError(http.StatusInternalServerError, err, common.ErrDatabaseError)
		return
	}

	pg := pagination.NewPagination(q.PageSize, q.PageNo)
	pg.SetTotal(total)

	result, err = u.repo.ListByUserID(ctx, userID, &ListParams{
		Limit:          pg.GetLimit(),
		Offset:         pg.GetOffset(),
		OrderFieldName: q.OrderFieldName,
		OrderIsDesc:    q.IsDesc,
	})
	if err != nil {
		if ent.IsNotFound(err) {
			err = common.NewUsecaseError(http.StatusNotFound, err, common.ErrNotfound)
			return
		}
		err = common.NewUsecaseError(http.StatusInternalServerError, err, common.ErrDatabaseError)
		return
	}

	pgInfo = pg.GetInfo(len(result))
	return
}
