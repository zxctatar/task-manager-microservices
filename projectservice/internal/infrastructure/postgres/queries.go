package postgres

var (
	QuerieSave = "INSERT INTO projects(owner_id, name) VALUES($1, $2)"
)
