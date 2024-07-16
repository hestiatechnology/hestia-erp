#!/bin/bash

# Wait for PostgreSQL to start
until pg_isready; do
  sleep 1;
done

# Generate a random password
RANDOM_PASSWORD=$(openssl rand -base64 12)

# Create a new superuser with the generated password
psql -U postgres -d test_db -c "CREATE ROLE hestia_erp_internal WITH LOGIN PASSWORD '$RANDOM_PASSWORD' SUPERUSER;"

# Print the username and password
echo "Superuser 'hestia_erp_internal' created with password: $RANDOM_PASSWORD"

# Use pg_restore to load the dump into the database
pg_restore -U postgres -d test_db /test_data.sql