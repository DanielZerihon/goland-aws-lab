# s3-create

Small Go example that:

- creates the S3 bucket `bombo-s3` if it does not already exist
- uploads `upload-example.txt` to that bucket

The app uses region `us-east-1`.

## Prerequisite

Configure AWS credentials before running the app.

```bash
aws configure
```

## Commands

Install and clean dependencies:

```bash
go mod tidy
```

Run tests:

```bash
go test ./...
```

Run the app:

```bash
go run .
```

Recommended check before push:

```bash
go mod tidy
go test ./...
```
