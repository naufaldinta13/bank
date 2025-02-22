package repository

import (
	"context"
	"math/rand"

	"github.com/naufaldinta13/bank/config"
	"github.com/naufaldinta13/bank/model"

	"github.com/uptrace/bun"
)

type CustomerRepository struct {
	db *bun.DB
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{
		db: config.GetDB(),
	}
}

func (r *CustomerRepository) Create(req *model.Customer) (e error) {
	_, e = r.db.NewInsert().Model(req).Exec(context.TODO())

	return
}

func (r *CustomerRepository) Show(id string) (mx *model.Customer, e error) {
	mx = new(model.Customer)
	if e = r.db.NewSelect().Model(mx).Where("id = ?", id).Scan(context.TODO()); e != nil {
		return nil, e
	}

	return
}

func (r *CustomerRepository) Update(req *model.Customer, fields ...string) (e error) {
	_, e = r.db.NewUpdate().Model(req).Column(fields...).WherePK().Exec(context.TODO())

	return
}

func (r *CustomerRepository) Delete(req *model.Customer) (e error) {
	req.IsDeleted = true

	_, e = r.db.NewUpdate().Model(req).Column("is_deleted").WherePK().Exec(context.TODO())

	return
}

func (r *CustomerRepository) FindByNik(nik string) (mx *model.Customer, e error) {
	mx = new(model.Customer)
	if e = r.db.NewRaw("SELECT * FROM customer WHERE nik = ? AND is_deleted = false ", nik).Scan(context.TODO(), mx); e != nil {
		return nil, e
	}

	return
}

func (r *CustomerRepository) FindByPhone(phone string) (mx *model.Customer, e error) {
	mx = new(model.Customer)
	if e = r.db.NewRaw("SELECT * FROM customer WHERE phone_number = ? AND is_deleted = false", phone).Scan(context.TODO(), mx); e != nil {
		return nil, e
	}

	return
}

func (r *CustomerRepository) FindByAccountNumber(an string) (mx *model.Customer, e error) {
	mx = new(model.Customer)
	if e = r.db.NewRaw("SELECT * FROM customer WHERE account_number = ? AND is_deleted = false", an).Scan(context.TODO(), mx); e != nil {
		return nil, e
	}

	return
}

func (r *CustomerRepository) GenerateAccountNumber() string {
	var letters = []rune("1234567890")
	b := make([]rune, 7)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func (r *CustomerRepository) SyncSaldo(id string) (e error) {
	query := `UPDATE customer uc
	SET saldo = x.saldo
	FROM customer c
	LEFT JOIN (
		select customer_id, sum(CASE WHEN type = 'deposit' THEN nominal ELSE -1 * nominal END) as saldo
		from transaction where customer_id = ? GROUP BY customer_id
	) as x ON x.customer_id = c.id
	WHERE uc.id = c.id AND c.id = ?`

	_, e = r.db.ExecContext(context.TODO(), query, id, id)

	return
}
