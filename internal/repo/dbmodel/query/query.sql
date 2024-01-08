-- name: CreateUser :one
INSERT INTO users (
    name,
    email,
    password
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: CreateAccount :one
INSERT INTO accounts (
    name,
    bank
) VALUES (
    $1, $2
 ) RETURNING *;

-- name: CreateTransaction :one
INSERT INTO transactions (
    account_id,
    amount,
    transaction_type
) VALUES (
    $1, $2, $3
) RETURNING *;

/* name: GetAccountForUpdate :one */
SELECT * FROM accounts WHERE id = $1 FOR UPDATE;

/* name: UpdateAccountBalance :exec */
UPDATE accounts SET balance = $1 where id = $2;