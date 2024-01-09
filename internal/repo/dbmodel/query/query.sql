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
    bank,
    user_id
) VALUES (
    $1, $2, $3
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
SELECT * FROM accounts WHERE id = $1 and user_id = $2 FOR UPDATE;

/* name: GetAccount :one */
SELECT * FROM accounts WHERE id = $1 and user_id = $2 ;

/* name: GetUserByEmail :one */
SELECT * FROM users WHERE email = $1 ;

/* name: UpdateAccountBalance :exec */
UPDATE accounts SET balance = $1 where id = $2;

/* name: ListAccounts :many */
SELECT * from accounts;

/* name: ListTransactions :many */
SELECT * from transactions;