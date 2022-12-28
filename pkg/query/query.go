package queryTools

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

type Interface interface {
	Query() *goqu.SelectDataset
	GetID() *goqu.SelectDataset
	SelectIsNotDeleted(b *goqu.SelectDataset) *goqu.SelectDataset
	SelectIsDeleted(b *goqu.SelectDataset) *goqu.SelectDataset
	UpdateIsNotDeleted(b *goqu.UpdateDataset, t string) *goqu.UpdateDataset
	DeleteIsDeleted()

	//SoftDeleteAll(t string) *goqu.UpdateDataset
}

type QueryTools struct {
	Builder *goqu.SelectDataset
	Table   string
}

func NewQueryTools(t string) Interface {
	var qt QueryTools
	qt.Dialect()
	qt.Table = t
	qt.Builder = qt.Dialect().From(t).Prepared(true)
	return qt
}

func (qt QueryTools) Query() *goqu.SelectDataset {
	return qt.Builder
}

func (qt QueryTools) Dialect() goqu.DialectWrapper {
	return goqu.Dialect("postgres")
}

func (qt QueryTools) GetID() *goqu.SelectDataset {
	return qt.Builder.From(qt.Table).Select(".id")
}

func (qt QueryTools) SelectIsNotDeleted(
	b *goqu.SelectDataset,
) *goqu.SelectDataset {
	// query = query.Where(goqu.C("is_deleted").IsNull())
	return b.Where(goqu.ExOr{qt.Table + ".is_deleted": false})
}

func (qt QueryTools) SelectIsDeleted(
	b *goqu.SelectDataset,
) *goqu.SelectDataset {
	return b.Where(goqu.ExOr{qt.Table + ".is_deleted": true})
}

func (qt QueryTools) UpdateIsNotDeleted(
	b *goqu.UpdateDataset,
	t string,
) *goqu.UpdateDataset {
	isDeleted := "is_deleted"
	if t != "" {
		isDeleted = t + ".is_deleted"
	}
	expr := goqu.Ex{isDeleted: false}
	return b.Where(expr)
}

func (qt QueryTools) DeleteIsDeleted() {
	isDeleted := "is_deleted"
	if qt.Table != "" {
		isDeleted = qt.Table + ".is_deleted"
	}
	fmt.Println(isDeleted)
	//expr := goqu.Ex{isDeleted: false}
	//b = b.Where(expr)
	//return b
}

//
//func (qt QueryTools) SoftDeleteAll(t string) *goqu.UpdateDataset {
//	r := goqu.Record{
//		"is_deleted": true,
//		"deleted_at": time.Now(),
//	}
//	b := qt.SelectIsNotDeleted(t)
//	b = b.Update().Set(r)
//	//b := qt.Dialect().Update(t).Set(r)
//	//b = qt.UpdateIsNotDeleted(b, t)
//	return b
//}
