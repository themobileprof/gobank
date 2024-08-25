.PHONY: migration_up migration_down migration_test_up migration_test_down
ENV := $(PWD)/.env
TEST_ENV := $(PWD)/.env.testing

include $(ENV)
include $(TEST_ENV)

migration_up:
	migrate -path db/migrations -database "mysql://$(DBUSER):$(DBPASS)@tcp(localhost:3306)/$(DB)" -verbose up

migration_down:
	migrate -path db/migrations -database "mysql://$(DBUSER):$(DBPASS)@tcp(localhost:3306)/$(DB)" -verbose down

migration_test_up:
	migrate -path db/migrations -database "mysql://$(DBTESTUSER):$(DBTESTPASS)@tcp(localhost:3306)/$(DBTEST)" -verbose up

migration_test_down:
	migrate -path db/migrations -database "mysql://$(DBTESTUSER):$(DBTESTPASS)@tcp(localhost:3306)/$(DBTEST)" -verbose down