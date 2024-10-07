
package user

import (
    "github.com/beardfriend/ddalggak/ent"
    "github.com/beardfriend/ddalggak/ent/user"
    "context"
    "entgo.io/ent/dialect/sql"
)

type Repo interface {
    Create(ctx context.Context, b *ent.User) error
    Get(ctx context.Context, id int) (*ent.User, error)
    Update(ctx context.Context, b *ent.User) (err error)
    Delete(ctx context.Context, id int) (err error)
    List(ctx context.Context, p *ListParams) ([]*ent.User, error)
    Total(ctx context.Context) (int, error)
}

type repo struct {
    db *ent.Client
}

func NewRepo(db *ent.Client) Repo {
    return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, b *ent.User) (err error) {
    err = r.db.User.Create().
        SetEmail(b.Email).
        SetPassword(b.Password).
        SetNickname(b.Nickname).
        Exec(ctx)
    return 
}

func (r *repo) Get(ctx context.Context, id int) (*ent.User, error) {
    return r.db.User.Get(ctx, id)
}

func (r *repo) Update(ctx context.Context, b *ent.User) (err error) {
    err = r.db.User.UpdateOne(b).
        SetEmail(b.Email).
        SetPassword(b.Password).
        SetNickname(b.Nickname).
        Exec(ctx)
    return 
}

func (r *repo) Delete(ctx context.Context, id int) (err error) {
    err = r.db.User.DeleteOneID(id).Exec(ctx)
    return 
}


type ListParams struct {
    Limit int
    Offset int
    OrderFieldName string
    OrderIsDesc bool
}

func (r *repo) List(ctx context.Context, p *ListParams) ([]*ent.User, error) {
    q := r.db.User.
        Query().
        Limit(p.Limit).
        Offset(p.Offset)

    if p.OrderFieldName == "" {
		p.OrderFieldName = user.FieldUpdatedAt
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
    return r.db.User.Query().Count(ctx)
}
