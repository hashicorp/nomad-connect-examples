# uuid-fe

## Configuration

#### Environment variables

The `uuid-fe` command is configured with `PORT` and `UPSTREAM` environment variables.

```bash
$ PORT=2002 UPSTREAM=uuid-api go run main.go
```

## Docker

#### Usage

Run `uuid-fe` with docker (listens at `0.0.0.0:PORT`)

```bash
$ docker run --env PORT=2002 --env UPSTREAM=uuid-api --net=host --rm hashicorpnomad/uuid-fe:v4
```

#### Publish

Build image and push to docker hub.

`<version> format is `v<n>-<arch>` (e.g. `v1-amd64`).

```bash
$ docker build -t hashicorpnomad/uuid-fe:<version> .
$ docker push hashicorpnomad/uuid-fe:<version>
```

Also build and publish for `arm64` (e.g. Graviton)

Publish a manifest for `v<n>`.

```bash
$ docker manifest create hashicorpnomad/uuid-fe:<v> --amend hashicorpnomad/uuid-fe:<v>-arm64 --amend hashicorpnomad/uuid-fe:<v>-amd64
$ docker manifest push hashicorpnomad/uuid-fe:<v>
```
