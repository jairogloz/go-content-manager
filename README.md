# go-content-manager

## How to set config values

1. Make a copy of `example.env` and name it `.env`.

```bash
cp example.env .env
```

2. Populate the new `.env` file with your own configuration values.

3. Export the whole file to your environment by running:

```bash
export $(cat .env | xargs)
```

## How to run with Docker

1. Build your Docker image with:

```bash
docker build -t content-manager .
```

2. Run your container with:

```bash
docker run -d --name content-manager-container --env-file .env -p 8080:8080 content-manager
```
