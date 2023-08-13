db.run:
	docker run --rm --name godb -e POSTGRES_PASSWORD=mysecretpassword -p 5433:5432 -d postgres:15.4-alpine3.18
db.stop:
	docker stop godb