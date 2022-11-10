package product

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Implementacion-Storage/Project/internal/domain"
	"github.com/Juram009/backpack-bcgow6-juan-ramirez/Implementacion-Storage/Project/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

// GO-TXDB
func TestSaveAndGetOneTXDB(t *testing.T) {
	db := utils.InitTxDB()
	defer db.Close()

	repo := NewRepository(db)

	ctx := context.TODO()

	expectedProduct := domain.Product{
		Name:        "Mesa",
		Type:        "Mueble",
		Count:       20,
		Price:       10000,
		IDWarehouse: 1,
	}

	// Act
	id, err := repo.Save(ctx, expectedProduct)
	assert.NoError(t, err)

	expectedProduct.ID = int(id)
	product, err := repo.Get(context.TODO(), int(id))

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, product)
	assert.Equal(t, expectedProduct.ID, product.ID)
}

func TestUpdateTXDB(t *testing.T) {
	db := utils.InitTxDB()
	defer db.Close()

	repo := NewRepository(db)

	ctx := context.TODO()

	productBefore := domain.Product{
		Name:        "Mesa",
		Type:        "Mueble",
		Count:       20,
		Price:       10000,
		IDWarehouse: 1,
	}

	// Act
	id, err := repo.Save(ctx, productBefore)
	assert.NoError(t, err)

	expectedProduct := domain.Product{
		ID:          int(id),
		Name:        "Mesaa",
		Type:        "Mueble",
		Count:       20,
		Price:       10000,
		IDWarehouse: 1,
	}
	productBefore.ID = int(id)

	err = repo.Update(context.TODO(), expectedProduct, int(id))
	assert.NoError(t, err)

	product, err := repo.Get(context.TODO(), int(id))
	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, product)
	assert.Equal(t, expectedProduct, product)
}

func TestDeleteTXDB(t *testing.T) {
	db := utils.InitTxDB()
	defer db.Close()

	repo := NewRepository(db)

	ctx := context.TODO()

	product := domain.Product{
		Name:        "Mesa",
		Type:        "Mueble",
		Count:       20,
		Price:       10000,
		IDWarehouse: 1,
	}

	// Act
	id, err := repo.Save(ctx, product)
	assert.NoError(t, err)
	product.ID = int(id)

	err = repo.Delete(context.TODO(), id)
	assert.NoError(t, err)

	product, err = repo.Get(context.TODO(), int(id))
	assert.Error(t, err)
	products, err := repo.GetAll(context.TODO())
	// Assert
	assert.Error(t, err)
	assert.Empty(t, product)
	assert.Empty(t, products)
	assert.Equal(t, domain.Product{}, product)
}

//GO-SQLMOCK

var product_test = domain.Product{
	ID:          1,
	Name:        "Mesa",
	Type:        "Mueble",
	Count:       20,
	Price:       10000,
	IDWarehouse: 1,
}

func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Store Ok", func(t *testing.T) {

		mock.ExpectPrepare(regexp.QuoteMeta(SAVE_PRODUCT))
		mock.ExpectExec(regexp.QuoteMeta(SAVE_PRODUCT)).WillReturnResult(sqlmock.NewResult(1, 1))

		columns := []string{"id", "name", "type", "count", "price", "id_warehouse"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(product_test.ID, product_test.Name, product_test.Type, product_test.Count, product_test.Price, product_test.IDWarehouse)
		mock.ExpectQuery(regexp.QuoteMeta(GET_PRODUCT)).WithArgs(1).WillReturnRows(rows)

		repository := NewRepository(db)
		ctx := context.TODO()

		newID, err := repository.Save(ctx, product_test)
		assert.NoError(t, err)

		res, err := repository.Get(ctx, int(newID))
		assert.NoError(t, err)

		assert.NotNil(t, res)
		assert.Equal(t, product_test.ID, res.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Store Fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		errSaving := errors.New("Error saving")
		mock.ExpectPrepare(regexp.QuoteMeta(SAVE_PRODUCT))
		mock.ExpectExec(regexp.QuoteMeta(SAVE_PRODUCT)).WillReturnError(errSaving)

		repository := NewRepository(db)
		ctx := context.TODO()

		id, err := repository.Save(ctx, product_test)

		assert.EqualError(t, err, errSaving.Error())
		assert.Empty(t, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Update Ok", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(UPDATE_PRODUCT))
		mock.ExpectExec(regexp.QuoteMeta(UPDATE_PRODUCT)).WithArgs(product_test.Name, product_test.Type, product_test.Count, product_test.Price, product_test.IDWarehouse, product_test.ID).WillReturnResult(sqlmock.NewResult(0, 1))

		//mock.ExpectQuery(regexp.QuoteMeta(GET_PRODUCT)).WithArgs(1).WillReturnRows(rows)

		repository := NewRepository(db)
		ctx := context.TODO()

		err = repository.Update(ctx, product_test, int(product_test.ID))
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Update Fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		errUpdating := errors.New("Error updating")
		mock.ExpectPrepare(regexp.QuoteMeta(UPDATE_PRODUCT))
		mock.ExpectExec(regexp.QuoteMeta(UPDATE_PRODUCT)).WillReturnError(errUpdating)

		repository := NewRepository(db)
		ctx := context.TODO()

		err = repository.Update(ctx, product_test, product_test.ID)

		assert.EqualError(t, err, errUpdating.Error())

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Delete Ok", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(DELETE_PRODUCT))
		mock.ExpectExec(regexp.QuoteMeta(DELETE_PRODUCT)).WithArgs(product_test.ID).WillReturnResult(sqlmock.NewResult(0, 1))

		repository := NewRepository(db)
		ctx := context.TODO()

		err = repository.Delete(ctx, int64(product_test.ID))
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Update Fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()
		errUpdating := errors.New("Error updating")
		mock.ExpectPrepare(regexp.QuoteMeta(UPDATE_PRODUCT))
		mock.ExpectExec(regexp.QuoteMeta(UPDATE_PRODUCT)).WillReturnError(errUpdating)

		repository := NewRepository(db)
		ctx := context.TODO()

		err = repository.Update(ctx, product_test, product_test.ID)

		assert.EqualError(t, err, errUpdating.Error())
		//assert.Empty(t,id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
