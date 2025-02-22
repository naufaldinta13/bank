package repository

import (
	"context"

	"github.com/naufaldinta13/bank/config"
	"github.com/naufaldinta13/bank/model"

	"github.com/uptrace/bun"
)

type TransactionRepository struct {
	db *bun.DB
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{
		db: config.GetDB(),
	}
}

func (r *TransactionRepository) Create(req *model.Transaction) (e error) {
	_, e = r.db.NewInsert().Model(req).Exec(context.TODO())

	return
}

func (r *TransactionRepository) Show(id string) (mx *model.Transaction, e error) {
	mx = new(model.Transaction)
	if e = r.db.NewSelect().Model(mx).Where("id = ?", id).Scan(context.TODO()); e != nil {
		return nil, e
	}

	return
}

func (r *TransactionRepository) Update(req *model.Transaction, fields ...string) (e error) {
	_, e = r.db.NewUpdate().Model(req).Column(fields...).WherePK().Exec(context.TODO())

	return
}

func (r *TransactionRepository) Delete(req *model.Transaction) (e error) {
	req.IsDeleted = true

	_, e = r.db.NewUpdate().Model(req).Column("is_deleted").WherePK().Exec(context.TODO())

	return
}
