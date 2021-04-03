.PHONY: migrateCreate

migrateCreate:
	migrate create -ext cql -dir revisions -seq $(MESSAGE)