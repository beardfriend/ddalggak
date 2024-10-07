package user

import (
	"context"
	"net/http"

	"github.com/beardfriend/ddalggak/ent"
	"github.com/beardfriend/ddalggak/internal/common"
	"github.com/beardfriend/ddalggak/pkg/pagination"
)

type Usecase interface {
	Create(ctx context.Context, b *ent.User) (err error)
	Get(ctx context.Context, id int) (result *ent.User, err error)
	Update(ctx context.Context, b *ent.User) (err error)
	Delete(ctx context.Context, id int) (err error)
	List(ctx context.Context, q *common.ListRequest) (result []*ent.User, pgInfo *pagination.PaginationInfo, err error)
}

type usecase struct {
	repo Repo
}

func NewUsecase(repo Repo) Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) Create(ctx context.Context, b *ent.User) (err error) {
	err = u.repo.Create(ctx, b)
	if err != nil {
		err = common.NewUsecaseError(http.StatusInternalServerError, err, common.ErrDatabaseError)
		return
	}
	return
}

func (u *usecase) Get(ctx context.Context, id int) (result *ent.User, err error) {
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

func (u *usecase) Update(ctx context.Context, b *ent.User) (err error) {
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

func (u *usecase) List(ctx context.Context, q *common.ListRequest) (result []*ent.User, pgInfo *pagination.PaginationInfo, err error) {
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
