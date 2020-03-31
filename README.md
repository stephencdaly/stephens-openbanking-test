# stephens-openbanking-test
Playing around with creating an open banking demo in Go


## Building
```go build -o stephens-openbanking-test .```

## Running

```docker run -d -p 5432:5432 --name openbanking -e POSTGRES_PASSWORD=mysupersecretpassword postgres```

```DATABASE_URL=postgres://postgres:mysupersecretpassword@localhost:5432/?sslmode=disable ./stephens-openbanking-test```