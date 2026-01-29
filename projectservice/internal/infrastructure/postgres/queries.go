package postgres

var (
	QuerieFindByName = "SELECT id, owner_id, name, created_at FROM projects WHERE owner_id = $1 AND name = $2"
	QuerieSave       = "INSERT INTO projects(owner_id, name) VALUES($1, $2)"
)
