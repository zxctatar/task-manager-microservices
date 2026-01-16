package postgres

const (
	QueryFindByEmail = `
	SELECT 
		id, 
		first_name, 
		middle_name, 
		last_name, 
		hash_password, 
		email 
	FROM users 
	WHERE email = $1`

	QuerySaveUser = `
	INSERT INTO users (
		first_name,
		middle_name,
		last_name,
		hash_password,
		email
	) VALUES ($1, $2, $3, $4, $5) RETURNING id`
)
