#!/bin/sh
set -e

echo "Running migrations"

for file in /app/migrations/*; do
  echo "Applying $file..."
  psql -U $POSTGRES_USER -d $POSTGRES_DB -f $file
done

echo "Migrations completed"
