// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/txrps/next-golang-project/model"
)

func newTmMenu(db *gorm.DB, opts ...gen.DOOption) tmMenu {
	_tmMenu := tmMenu{}

	_tmMenu.tmMenuDo.UseDB(db, opts...)
	_tmMenu.tmMenuDo.UseModel(&model.TmMenu{})

	tableName := _tmMenu.tmMenuDo.TableName()
	_tmMenu.ALL = field.NewAsterisk(tableName)
	_tmMenu.MenuID = field.NewInt32(tableName, "MenuID")
	_tmMenu.ParentID = field.NewInt32(tableName, "ParentID")
	_tmMenu.MenuName = field.NewString(tableName, "MenuName")
	_tmMenu.MenuType = field.NewInt32(tableName, "MenuType")
	_tmMenu.Level = field.NewInt32(tableName, "Level")
	_tmMenu.Route = field.NewString(tableName, "Route")
	_tmMenu.Icon = field.NewString(tableName, "Icon")
	_tmMenu.IsDisplay = field.NewBool(tableName, "IsDisplay")
	_tmMenu.IsSetPermission = field.NewBool(tableName, "IsSetPermission")
	_tmMenu.IsDisable = field.NewBool(tableName, "IsDisable")
	_tmMenu.IsView = field.NewBool(tableName, "IsView")
	_tmMenu.IsManage = field.NewBool(tableName, "IsManage")
	_tmMenu.IsShowbreadcrumb = field.NewBool(tableName, "IsShowbreadcrumb")
	_tmMenu.IsActive = field.NewBool(tableName, "IsActive")
	_tmMenu.Order_ = field.NewFloat64(tableName, "Order")

	_tmMenu.fillFieldMap()

	return _tmMenu
}

type tmMenu struct {
	tmMenuDo

	ALL              field.Asterisk
	MenuID           field.Int32
	ParentID         field.Int32
	MenuName         field.String
	MenuType         field.Int32
	Level            field.Int32
	Route            field.String
	Icon             field.String
	IsDisplay        field.Bool
	IsSetPermission  field.Bool
	IsDisable        field.Bool
	IsView           field.Bool
	IsManage         field.Bool
	IsShowbreadcrumb field.Bool
	IsActive         field.Bool
	Order_           field.Float64

	fieldMap map[string]field.Expr
}

func (t tmMenu) Table(newTableName string) *tmMenu {
	t.tmMenuDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tmMenu) As(alias string) *tmMenu {
	t.tmMenuDo.DO = *(t.tmMenuDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tmMenu) updateTableName(table string) *tmMenu {
	t.ALL = field.NewAsterisk(table)
	t.MenuID = field.NewInt32(table, "MenuID")
	t.ParentID = field.NewInt32(table, "ParentID")
	t.MenuName = field.NewString(table, "MenuName")
	t.MenuType = field.NewInt32(table, "MenuType")
	t.Level = field.NewInt32(table, "Level")
	t.Route = field.NewString(table, "Route")
	t.Icon = field.NewString(table, "Icon")
	t.IsDisplay = field.NewBool(table, "IsDisplay")
	t.IsSetPermission = field.NewBool(table, "IsSetPermission")
	t.IsDisable = field.NewBool(table, "IsDisable")
	t.IsView = field.NewBool(table, "IsView")
	t.IsManage = field.NewBool(table, "IsManage")
	t.IsShowbreadcrumb = field.NewBool(table, "IsShowbreadcrumb")
	t.IsActive = field.NewBool(table, "IsActive")
	t.Order_ = field.NewFloat64(table, "Order")

	t.fillFieldMap()

	return t
}

func (t *tmMenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tmMenu) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 15)
	t.fieldMap["MenuID"] = t.MenuID
	t.fieldMap["ParentID"] = t.ParentID
	t.fieldMap["MenuName"] = t.MenuName
	t.fieldMap["MenuType"] = t.MenuType
	t.fieldMap["Level"] = t.Level
	t.fieldMap["Route"] = t.Route
	t.fieldMap["Icon"] = t.Icon
	t.fieldMap["IsDisplay"] = t.IsDisplay
	t.fieldMap["IsSetPermission"] = t.IsSetPermission
	t.fieldMap["IsDisable"] = t.IsDisable
	t.fieldMap["IsView"] = t.IsView
	t.fieldMap["IsManage"] = t.IsManage
	t.fieldMap["IsShowbreadcrumb"] = t.IsShowbreadcrumb
	t.fieldMap["IsActive"] = t.IsActive
	t.fieldMap["Order"] = t.Order_
}

func (t tmMenu) clone(db *gorm.DB) tmMenu {
	t.tmMenuDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tmMenu) replaceDB(db *gorm.DB) tmMenu {
	t.tmMenuDo.ReplaceDB(db)
	return t
}

type tmMenuDo struct{ gen.DO }

func (t tmMenuDo) Debug() *tmMenuDo {
	return t.withDO(t.DO.Debug())
}

func (t tmMenuDo) WithContext(ctx context.Context) *tmMenuDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tmMenuDo) ReadDB() *tmMenuDo {
	return t.Clauses(dbresolver.Read)
}

func (t tmMenuDo) WriteDB() *tmMenuDo {
	return t.Clauses(dbresolver.Write)
}

func (t tmMenuDo) Session(config *gorm.Session) *tmMenuDo {
	return t.withDO(t.DO.Session(config))
}

func (t tmMenuDo) Clauses(conds ...clause.Expression) *tmMenuDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tmMenuDo) Returning(value interface{}, columns ...string) *tmMenuDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tmMenuDo) Not(conds ...gen.Condition) *tmMenuDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tmMenuDo) Or(conds ...gen.Condition) *tmMenuDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tmMenuDo) Select(conds ...field.Expr) *tmMenuDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tmMenuDo) Where(conds ...gen.Condition) *tmMenuDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tmMenuDo) Order(conds ...field.Expr) *tmMenuDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tmMenuDo) Distinct(cols ...field.Expr) *tmMenuDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tmMenuDo) Omit(cols ...field.Expr) *tmMenuDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tmMenuDo) Join(table schema.Tabler, on ...field.Expr) *tmMenuDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tmMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tmMenuDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tmMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) *tmMenuDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tmMenuDo) Group(cols ...field.Expr) *tmMenuDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tmMenuDo) Having(conds ...gen.Condition) *tmMenuDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tmMenuDo) Limit(limit int) *tmMenuDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tmMenuDo) Offset(offset int) *tmMenuDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tmMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tmMenuDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tmMenuDo) Unscoped() *tmMenuDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tmMenuDo) Create(values ...*model.TmMenu) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tmMenuDo) CreateInBatches(values []*model.TmMenu, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tmMenuDo) Save(values ...*model.TmMenu) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tmMenuDo) First() (*model.TmMenu, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TmMenu), nil
	}
}

func (t tmMenuDo) Take() (*model.TmMenu, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TmMenu), nil
	}
}

func (t tmMenuDo) Last() (*model.TmMenu, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TmMenu), nil
	}
}

func (t tmMenuDo) Find() ([]*model.TmMenu, error) {
	result, err := t.DO.Find()
	return result.([]*model.TmMenu), err
}

func (t tmMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TmMenu, err error) {
	buf := make([]*model.TmMenu, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tmMenuDo) FindInBatches(result *[]*model.TmMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tmMenuDo) Attrs(attrs ...field.AssignExpr) *tmMenuDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tmMenuDo) Assign(attrs ...field.AssignExpr) *tmMenuDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tmMenuDo) Joins(fields ...field.RelationField) *tmMenuDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tmMenuDo) Preload(fields ...field.RelationField) *tmMenuDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tmMenuDo) FirstOrInit() (*model.TmMenu, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TmMenu), nil
	}
}

func (t tmMenuDo) FirstOrCreate() (*model.TmMenu, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TmMenu), nil
	}
}

func (t tmMenuDo) FindByPage(offset int, limit int) (result []*model.TmMenu, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tmMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tmMenuDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tmMenuDo) Delete(models ...*model.TmMenu) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tmMenuDo) withDO(do gen.Dao) *tmMenuDo {
	t.DO = *do.(*gen.DO)
	return t
}
