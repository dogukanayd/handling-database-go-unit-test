package product

import (
	"database/sql"
	"github.com/dogukanayd/handling-database-go-unit-test/internal/tests/mysqlunit"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

var activeProduct = Product{
	Name: "active product",
	Status: true,
}

var inActiveProduct = Product{
	Name: "inactive product",
	Status: false,
}

func productFactory(connection *sql.DB, product Product) {
	_, _ = connection.Exec("INSERT INTO products (`name`, `status`) VALUES (?, ?)", product.Name, product.Status)
}

func TestCheckActivationStatusByName(t *testing.T) {
	connection, tearDown := mysqlunit.NewUnit(t)
	defer tearDown()
	productFactory(connection, activeProduct)
	productFactory(connection, inActiveProduct)

	type args struct {
		db   *sql.DB
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:    "it should return true when active product name give",
			args:    args{
				db:   connection,
				name: activeProduct.Name,
			},
			want:    true,
		},
		{
			name: "it should return false when inactive product name given",
			args: args{
				db:   connection,
				name: inActiveProduct.Name,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckActivationStatusByName(tt.args.db, tt.args.name); got != tt.want {
				t.Errorf("CheckActivationStatusByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
