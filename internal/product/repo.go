
package product

import (
    "github.com/beardfriend/ddalggak/ent"
    "github.com/beardfriend/ddalggak/ent/product"
    "context"
    "entgo.io/ent/dialect/sql"
)

type Repo interface {
    Create(ctx context.Context, b *ent.Product) error
    Get(ctx context.Context, id int) (*ent.Product, error)
    Update(ctx context.Context, b *ent.Product) (err error)
    Delete(ctx context.Context, id int) (err error)
    List(ctx context.Context, p *ListParams) ([]*ent.Product, error)
    Total(ctx context.Context) (int, error)
    GetByUserID(ctx context.Context, id int, userID int) (*ent.Product, error)
    DeleteOneByUserID(ctx context.Context, id int, userID int) (err error)
    UpdateOneByUserID(ctx context.Context, b *ent.Product) (err error)
    TotalByUserID(ctx context.Context, userID int) (int, error)
    ListByUserID(ctx context.Context, userID int, p *ListParams) ([]*ent.Product, error)
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
    Limit int
    Offset int
    OrderFieldName string
    OrderIsDesc bool
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

func (r *repo) Total(ctx context.Context) (int, error) {
    return r.db.Product.Query().Count(ctx)
}

// func by userID

func (r *repo) GetByUserID(ctx context.Context, id int, userID int) (*ent.Product, error) {
    return r.db.Product.Query().Where(
        product.And(
            product.IDEQ(id),
            product.UserIDEQ(userID),
        ),
    ).Only(ctx)
}

func (r *repo) DeleteOneByUserID(ctx context.Context, id int, userID int) (err error) {
    _, err = r.db.Product.Delete().Where(
        product.And(
            product.IDEQ(id),
            product.UserIDEQ(userID),
        ),
    ).Exec(ctx)
    return 
}

func (r *repo) UpdateOneByUserID(ctx context.Context, b *ent.Product) (err error) {
    err = r.db.Product.Update().
        SetName(b.Name).
        SetPrice(b.Price).
    Where(
        product.And(
            product.IDEQ(b.ID),
            product.UserIDEQ(b.UserID),
        ),
    ).Exec(ctx)
    return 
}


func (r *repo) TotalByUserID(ctx context.Context, userID int) (int, error) {
    return r.db.Product.
        Query().
        Where(product.UserIDEQ(userID)).
        Count(ctx)
}

func (r *repo) ListByUserID(ctx context.Context, userID int, p *ListParams) ([]*ent.Product, error) { 
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

    return q.Where(product.UserIDEQ(userID)).All(ctx)
}
