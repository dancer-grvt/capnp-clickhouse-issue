init-stack:
	docker compose up -d
	sleep 10
	clickhouse client < clickhouseConfig/sqlSchema/00_oldTable.sql
	clickhouse client < clickhouseConfig/sqlSchema/01_newTable.sql

teardown-stack:
	docker compose down
