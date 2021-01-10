# uuid-api

## Configuration

#### Environment variables

The `uuid-api` command is configured with just the `PORT` environment variable.

```bash
$ PORT=2001 go run main.go
```

## Docker

#### Usage

Run `uuid-api` with docker (listens at `0.0.0.0:PORT`).

```bash
$ docker run --env PORT=2001 --net=host --rm hashicorpnomad/uuid-api:v4
```

#### Publish

Build image and push to docker hub.

`<version>` format is `v<n>-<arch>` (e.g. `v1-amd64`).

```bash
$ docker build --no-cache -t hashicorpnomad/uuid-api:<version> .
$ docker push hashicorpnomad/uuid-api:<version>
```

Also build and publish for `arm64` (e.g. Graviton)

Publish a manifest for `v<n>`.

```bash
$ docker manifest create hashicorpnomad/uuid-api:<v> --amend hashicorpnomad/uuid-api:<v>-arm64 --amend hashicorpnomad/uuid-api:<v>-amd64
$ docker manifest push hashicorpnomad/uuid-api:<v>
```
