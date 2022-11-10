package products

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/storage/internal/domain"
)

type Repository interface {
	GetByName(ctx context.Context, name string) (domain.Product, error)
	Store(ctx context.Context, p domain.Product) (int, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, p domain.Product) error
	Exists(context.Context, int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
	STORE        = "INSERT INTO products (name, type, count, price VALUES (?, ?, ?, ?, ?);"
	GET_BY_NAME  = "SELECT id, name, type, count, price FROM products WHERE name=?;"
	GET_ALL      = "SELECT id, name, type, count, price FROM products"
	DELETE       = "DELETE FROM products WHERE id=?"
	UPDATE       = "UPDATE products SET name=?, type=?, count=?, price=? WHERE id=?;"
	EXISTS_QUERY = "SELECT id FROM products WHERE id=?;"
)

func (r *repository) GetByName(ctx context.Context, name string) (domain.Product, error) {
	row := r.db.QueryRow(GET_BY_NAME, name)
	var product domain.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) Store(ctx context.Context, p domain.Product) (int, error) {
	stm, err := r.db.Prepare(STORE)
	if err != nil {
		return 0, err
	}
	defer stm.Close()
	result, err := stm.Exec(p.Name, p.Type, p.Count, p.Price)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	row, err := r.db.Query(GET_ALL)
	if err != nil {
		return []domain.Product{}, err
	}
	var products []domain.Product
	for row.Next() {
		var product domain.Product
		if err := row.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			return []domain.Product{}, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stm, err := r.db.Prepare(DELETE)
	if err != nil {
		return err
	}

	defer stm.Close() //cierro la sentencia para no gastar memoria

	_, err = stm.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(ctx context.Context, p domain.Product) error {
	stm, err := r.db.Prepare(UPDATE)
	if err != nil {
		return err
	}
	defer stm.Close() //cerramos para no perder memoria

	//ejecutamos la consulta con aquellos valores a remplazar en la sentencia
	result, err := stm.Exec(p.Name, p.Type, p.Count, p.Price)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("error: no affected rows")
	}
	return nil
}

func (r *repository) Exists(ctx context.Context, id int) bool {
	row := r.db.QueryRow(EXISTS_QUERY, id)
	err := row.Scan(&id)
	return err == nil
}
