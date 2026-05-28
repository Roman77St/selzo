-- Enable pgcrypto extension for gen_random_uuid().
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "citext";

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Users table.
--
-- Stores application identities and authorization roles.
-- Business-related profile data must be stored separately.
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    email CITEXT NOT NULL
        CONSTRAINT users_email_unique UNIQUE,

    email_verified_at TIMESTAMPTZ,

    role TEXT NOT NULL
        CONSTRAINT users_role_check
        CHECK (
            role IN (
                'buyer',
                'supplier',
                'admin'
            )
        ),

    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


-- User authentication credentials.
--
-- Separated from users table to keep authentication logic isolated
-- and allow future support for OAuth, MFA, WebAuthn, etc.
CREATE TABLE user_credentials (
    user_id UUID PRIMARY KEY
        CONSTRAINT user_credentials_user_id_fkey
        REFERENCES users(id)
        ON DELETE CASCADE,

    password_hash TEXT NOT NULL,

    password_changed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


-- Password reset tokens.
--
-- Raw reset tokens must never be stored in the database.
-- Only token hashes are persisted.
CREATE TABLE password_reset_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL
        CONSTRAINT password_reset_tokens_user_id_fkey
        REFERENCES users(id)
        ON DELETE CASCADE,

    token_hash TEXT NOT NULL,

    expires_at TIMESTAMPTZ NOT NULL,

    used_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


-- Email verification tokens.
--
-- Used for account email confirmation during registration
-- and email change operations.
CREATE TABLE email_verification_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL
        CONSTRAINT email_verification_tokens_user_id_fkey
        REFERENCES users(id)
        ON DELETE CASCADE,

    token_hash TEXT NOT NULL
    CONSTRAINT email_verification_tokens_token_hash_unique
    UNIQUE,

    expires_at TIMESTAMPTZ NOT NULL,

    used_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


-- User personal profile data.
--
-- Stores personal information about a specific user account.
-- Organization/company data must be stored separately.
CREATE TABLE user_profiles (
    user_id UUID PRIMARY KEY
        CONSTRAINT user_profiles_user_id_fkey
        REFERENCES users(id)
        ON DELETE CASCADE,

    first_name TEXT,
    last_name TEXT,
    phone TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


-- Indexes.

CREATE INDEX idx_password_reset_tokens_user_id
ON password_reset_tokens(user_id);

CREATE INDEX idx_password_reset_tokens_expires_at
ON password_reset_tokens(expires_at);

CREATE INDEX idx_email_verification_tokens_user_id
ON email_verification_tokens(user_id);

CREATE INDEX idx_email_verification_tokens_expires_at
ON email_verification_tokens(expires_at);

CREATE TRIGGER users_set_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER user_credentials_set_updated_at
BEFORE UPDATE ON user_credentials
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER user_profiles_set_updated_at
BEFORE UPDATE ON user_profiles
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();