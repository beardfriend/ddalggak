package product

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/beardfriend/ddalggak/ent"
	"github.com/beardfriend/ddalggak/ent/product"
)

type Repo interface {
	Create(ctx context.Context, b *ent.Product) error
	Get(ctx context.Context, id int) (*ent.Product, error)
	Update(ctx context.Context, b *ent.Product) (err error)
	Delete(ctx context.Context, id int) (err error)
}

type repo struct {
	db *ent.Client
}

func NewRepo(db *ent.Client) Repo {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, b *ent.Product) (err error) {
	err = r.db.Product.Create().
		SetName(b.Name).
		SetUserID(b.UserID).
		SetPrice(b.Price).
		Exec(ctx)
	return
}

func (r *repo) Get(ctx context.Context, id int) (*ent.Product, error) {
	return r.db.Product.Get(ctx, id)
}

func (r *repo) Update(ctx context.Context, b *ent.Product) (err error) {
	err = r.db.Product.UpdateOne(b).
		SetName(b.Name).
		SetUserID(b.UserID).
		SetPrice(b.Price).
		Exec(ctx)
	return
}

func (r *repo) Delete(ctx context.Context, id int) (err error) {
	err = r.db.Product.DeleteOneID(id).Exec(ctx)
	return
}

type ListParams struct {
	Limit          int
	Offset         int
	OrderFieldName string
	OrderIsDesc    bool
}

func (r *repo) List(ctx context.Context, p *ListParams) ([]*ent.Product, error) {
	q := r.db.Product.
		Query().
		Limit(p.Limit).
		Offset(p.Offset)

	if p.OrderFieldName == "" {
		p.OrderFieldName = product.FieldUpdatedAt
	}

	q.Order(func(s *sql.Selector) {
		name := sql.Asc(p.OrderFieldName)
		if p.OrderIsDesc {
			name = sql.Desc(p.OrderFieldName)
		}
		s.OrderBy(name)
	})

	return q.All(ctx)
}
