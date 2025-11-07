-- Crear un índice único que compare emails en minúsculas
CREATE UNIQUE INDEX IF NOT EXISTS users_email_lower_idx
ON users (LOWER(email));

