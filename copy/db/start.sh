#!/bin/sh
pg_ctl restart -D data
psql -U postgres -f db_init.sql 
psql
