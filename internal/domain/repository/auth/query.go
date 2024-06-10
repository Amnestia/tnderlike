package authrepo

const (
	auth = `
		SELECT
			username,
			acc_email,
			acc_password
		FROM
			account
		WHERE
			acc_email = $1
			AND deleted_at IS NULL
		LIMIT 1
	`

	insertNewAccount = `
		INSERT INTO account(
			username,
			acc_email,
			acc_password,
			created_by,
			updated_by
		) VALUES (:username, :email, :password, :email, :email)
		RETURNING id
	`
)
