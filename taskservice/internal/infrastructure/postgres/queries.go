package postgres

var (
	QuerieCreate = "INSERT INTO tasks (project_id, description, deadlone) VALUES($1, $2, $3) RETURNING id;"
)
