package xorm_mysql

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm/core"
	"xorm.io/xorm/dialects"
	"xorm.io/xorm/schemas"
)

// Base represents a basic dialect and all real dialects could embed this struct
type Base struct {
	dialect dialects.Dialect
	uri     *dialects.URI
	quoter  schemas.Quoter
}

// Quoter returns the current database Quoter
func (db *Base) Quoter() schemas.Quoter {
	return db.quoter
}

// Init initialize the dialect
func (db *Base) Init(dialect dialects.Dialect, uri *dialects.URI) error {
	db.dialect, db.uri = dialect, uri
	return nil
}

// URI returns the uri of database
func (db *Base) URI() *dialects.URI {
	return db.uri
}

// FormatBytes formats bytes
func (db *Base) FormatBytes(bs []byte) string {
	return fmt.Sprintf("0x%x", bs)
}

// DropTableSQL returns drop table SQL
func (db *Base) DropTableSQL(tableName string) (string, bool) {
	quote := db.dialect.Quoter().Quote
	return fmt.Sprintf("DROP TABLE IF EXISTS %s", quote(tableName)), true
}

// HasRecords returns true if the SQL has records returned
func (db *Base) HasRecords(queryer core.Queryer, ctx context.Context, query string, args ...interface{}) (bool, error) {
	rows, err := queryer.QueryContext(ctx, query, args...)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}
	return false, nil
}

// IsColumnExist returns true if the column of the table exist
func (db *Base) IsColumnExist(queryer core.Queryer, ctx context.Context, tableName, colName string) (bool, error) {
	quote := db.dialect.Quoter().Quote
	query := fmt.Sprintf(
		"SELECT %v FROM %v.%v WHERE %v = ? AND %v = ? AND %v = ?",
		quote("COLUMN_NAME"),
		quote("INFORMATION_SCHEMA"),
		quote("COLUMNS"),
		quote("TABLE_SCHEMA"),
		quote("TABLE_NAME"),
		quote("COLUMN_NAME"),
	)
	return db.HasRecords(queryer, ctx, query, db.uri.DBName, tableName, colName)
}

// AddColumnSQL returns a SQL to add a column
func (db *Base) AddColumnSQL(tableName string, col *schemas.Column) string {
	s, _ := dialects.ColumnString(db.dialect, col, true)
	return fmt.Sprintf("ALTER TABLE %v ADD %v", db.dialect.Quoter().Quote(tableName), s)
}

// CreateIndexSQL returns a SQL to create index
func (db *Base) CreateIndexSQL(tableName string, index *schemas.Index) string {
	quoter := db.dialect.Quoter()
	var unique string
	var idxName string
	if index.Type == schemas.UniqueType {
		unique = " UNIQUE"
	}
	idxName = index.XName(tableName)
	return fmt.Sprintf("CREATE%s INDEX %v ON %v (%v)", unique,
		quoter.Quote(idxName), quoter.Quote(tableName),
		quoter.Join(index.Cols, ","))
}

// DropIndexSQL returns a SQL to drop index
func (db *Base) DropIndexSQL(tableName string, index *schemas.Index) string {
	quote := db.dialect.Quoter().Quote
	var name string
	if index.IsRegular {
		name = index.XName(tableName)
	} else {
		name = index.Name
	}
	return fmt.Sprintf("DROP INDEX %v ON %s", quote(name), quote(tableName))
}

// ModifyColumnSQL returns a SQL to modify SQL
func (db *Base) ModifyColumnSQL(tableName string, col *schemas.Column) string {
	s, _ := dialects.ColumnString(db.dialect, col, false)
	return fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN %s", tableName, s)
}

// ForUpdateSQL returns for updateSQL
func (db *Base) ForUpdateSQL(query string) string {
	return query + " FOR UPDATE"
}

// SetParams set params
func (db *Base) SetParams(params map[string]string) {
}
