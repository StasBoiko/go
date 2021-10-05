#!/bin/bash
export PGPASSWORD="$POSTGRES_PASSWORD"
psql -h $POSTGRES_HOST -d $POSTGRES_DBNAME -U $POSTGRES_USER -f /app/schema.sql
