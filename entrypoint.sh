#!/bin/sh
set -e
env | grep '^POSTGRES_' | sed 's/^POSTGRES_/PG/' > /app/.env
echo "Waiting..."
  sleep 1

exec ./contact-list