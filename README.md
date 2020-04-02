# stephens-openbanking-test
Playing around with creating an open banking demo in Go


## Building
```go build -o stephens-openbanking-test .```

## Running

```docker run -d -p 5432:5432 --name openbanking -e POSTGRES_PASSWORD=mysupersecretpassword postgres```

```DB_NAME=postgres DB_HOSTNAME=localhost:5432 DB_USERNAME=postgres DB_PASSWORD=mysupersecretpassword ./stephens-openbanking-test```