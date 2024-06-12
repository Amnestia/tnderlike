package authrepo

const (
	auth = `
		SELECT
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
			acc_email,
			acc_password,
			created_by,
			created_at,
			updated_by,
			updated_at
		) VALUES (:email, :password, :email, NOW(), :email, NOW())
		RETURNING id
	`
)
