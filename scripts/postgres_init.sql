-- user storage
CREATE TABLE IF NOT EXISTS %schema%.users (
  id UUID PRIMARY KEY,
  username TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
);

-- username / password authentication
CREATE TABLE IF NOT EXISTS %schema%.authentication_username_password (
  user_id UUID PRIMARY KEY REFERENCES users.id,
  password TEXT NOT NULL,
  domain TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  expiration_date TIMESTAMP
)

-- bearer token authentication
CREATE TABLE IF NOT EXISTS %schema%.authentication_tokens (
  user_id UUID PRIMARY KEY REFERENCES users.id,
  token TEXT NOT NULL,
  refresh_token TEXT,
  created_at TIMESTAMP NOT NULL,
  expiration_date TIMESTAMP
)
