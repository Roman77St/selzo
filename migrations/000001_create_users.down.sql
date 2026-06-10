DROP TRIGGER IF EXISTS users_set_updated_at
ON users;

DROP TRIGGER IF EXISTS user_credentials_set_updated_at
ON user_credentials;

DROP TRIGGER IF EXISTS user_profiles_set_updated_at
ON user_profiles;


DROP INDEX IF EXISTS idx_email_verification_tokens_expires_at;
DROP INDEX IF EXISTS idx_email_verification_tokens_user_id;

DROP INDEX IF EXISTS idx_password_reset_tokens_expires_at;
DROP INDEX IF EXISTS idx_password_reset_tokens_user_id;
DROP INDEX IF EXISTS idx_refresh_tokens_user_id;


DROP TABLE IF EXISTS user_profiles;

DROP TABLE IF EXISTS email_verification_tokens;

DROP TABLE IF EXISTS password_reset_tokens;

DROP TABLE IF EXISTS user_credentials;

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS refresh_tokens;


DROP FUNCTION IF EXISTS set_updated_at;


DROP EXTENSION IF EXISTS "citext";

DROP EXTENSION IF EXISTS "pgcrypto";
