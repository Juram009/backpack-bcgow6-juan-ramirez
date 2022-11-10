package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Implementacion-Storage/Project/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Save(ctx context.Context, b domain.Product) (int64, error)
	Exists(ctx context.Context, id int) bool
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, b domain.Product, id int) error
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
	SAVE_PRODUCT = "INSERT INTO products (name, type, count, price, id_warehouse) VALUES (?,?,?,?,?);"

	GET_ALL_PRODUCTS = "SELECT id , name, type, count, price FROM products;"

	GET_PRODUCT = "SELECT id, name, type, count, price, id_warehouse FROM products WHERE id=?;"

	EXIST_PRODUCT = "SELECT id FROM products WHERE id=?;"

	DELETE_PRODUCT = "DELETE FROM products WHERE id=?;"

	UPDATE_PRODUCT = "UPDATE products SET name=?, type=?, count=?, price=?, id_warehouse=? WHERE id=?;"
)

func (r *repository) Exists(ctx context.Context, id int) bool {
	rows := r.db.QueryRow(EXIST_PRODUCT, id)
	err := rows.Scan(&id)
	return err == nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	rows, err := r.db.Query(GET_ALL_PRODUCTS)
	if err != nil {
		return nil, err
	}

	var products []domain.Product

	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.IDWarehouse); err != nil {
			return []domain.Product{}, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) Save(ctx context.Context, p domain.Product) (int64, error) {
	stm, err := r.db.Prepare(SAVE_PRODUCT)
	if err != nil {
		return 0, err
	}

	//ejecutamos la consulta con aquellos valores a remplazar en la sentencia
	res, err := stm.Exec(p.Name, p.Type, p.Count, p.Price, p.IDWarehouse)
	if err != nil {
		return 0, err
	}

	//obtenemos el ultimo id
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Product, error) {
	row := r.db.QueryRow(GET_PRODUCT, id)
	var product domain.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.IDWarehouse); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) Update(ctx context.Context, p domain.Product, id int) error {
	stm, err := r.db.Prepare(UPDATE_PRODUCT)
	if err != nil {
		return err
	}
	defer stm.Close() //cerramos para no perder memoria

	//ejecutamos la consulta con aquellos valores a remplazar en la sentencia
	result, err := stm.Exec(p.Name, p.Type, p.Count, p.Price, p.IDWarehouse, id)
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

func (r *repository) Delete(ctx context.Context, id int64) error {
	stm, err := r.db.Prepare(DELETE_PRODUCT)
	if err != nil {
		return err
	}
	defer stm.Close()

	result, err := stm.Exec(id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("error: no affected rows")
	}

	return nil
}
