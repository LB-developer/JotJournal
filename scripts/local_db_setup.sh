#!/bin/bash
set -e

DB_NAME=jotjournal
DB_USER=jotuser
DB_PASS=securepassword

echo "Creating Postgres DB and user..."

psql -U postgres <<EOF
CREATE USER $DB_USER WITH PASSWORD '$DB_PASS';
CREATE DATABASE $DB_NAME OWNER $DB_USER;
GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;
EOF

echo "Done. Make sure .env is configured with these values (this is done by default if following examples)."

