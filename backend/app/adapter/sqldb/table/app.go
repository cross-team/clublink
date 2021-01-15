package table

// App represents database table columns for 'app' table
var App = struct {
	TableName       string
	ColumnID        string
	ColumnName      string
	ColumnCreatedAt string
}{
	TableName:       "app",
	ColumnID:        "id",
	ColumnName:      "name",
	ColumnCreatedAt: "created_at",
}
