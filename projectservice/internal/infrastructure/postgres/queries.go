package postgres

var (
	QuerieSave   = "INSERT INTO projects(owner_id, name) VALUES($1, $2)"
	QuerieDelete = "DELETE FROM projects WHERE owner_id = $1 AND name = $2"
)
