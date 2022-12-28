package queryTools

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func Dialect() goqu.DialectWrapper {
	return goqu.Dialect("postgres")
}

func GetID(b *goqu.SelectDataset, t string) *goqu.SelectDataset {
	return b.From(t).Select(".id")
}

func SelectIsNotDeleted(b *goqu.SelectDataset, t string) *goqu.SelectDataset {
	// query = query.Where(goqu.C("is_deleted").IsNull())
	return b.Where(goqu.ExOr{t + ".is_deleted": false})
}

func SelectIsDeleted(b *goqu.SelectDataset, t string) *goqu.SelectDataset {
	return b.Where(goqu.ExOr{t + ".is_deleted": true})
}

func DeleteIsDeleted(t string) {
	isDeleted := "is_deleted"
	if t != "" {
		isDeleted = t + ".is_deleted"
	}
	fmt.Println(isDeleted)
	//expr := goqu.Ex{isDeleted: false}
	//b = b.Where(expr)
	//return b
}
