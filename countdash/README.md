# countdash

Nomad makes use of the [demo-consul-101](https://github.com/hashicorp/demo-consul-101)
education material throughout our examples for using Consul Connect. There are many
example jobs that launch tasks with [counting-service](https://github.com/hashicorp/demo-consul-101/tree/master/services/counting-service)
and [dashboard-service](https://github.com/hashicorp/demo-consul-101/tree/master/services/dashboard-service)
configured to communicate under various conditions.

## Publishing

When making changes to the `countdash` demo, open a PR against the [demo-consul-101](https://github.com/hashicorp/demo-consul-101)
project. Be sure to preserve backwards compatibility, as these demo files are used
extensively, including with other Consul integrations.

The docker images specific to Nomad can be rebuilt using [counter-dashboard](counter-dashboard/Dockerfile)
and [counter-api](counter-api/Dockerfile).

#### Testing locally
```bash
$ docker build -t hashicorpnomad/counter-api:test .
$ docker build -t hashicorpnomad/counter-dashboard:test .

$ docker run --rm --net=host hashicorpnomad/counter-api:test
$ docker run --rm --net=host --env COUNTING_SERVICE_URL="http://localhost:9001" hashicorpnomad/counter-dashboard:test
```

#### Publish

Build image and push to docker hub.

`<version>` format is `v<n>-<arch>` (e.g. `v1-amd64`).

```bash
$ docker build --no-cache -t hashicorpnomad/counter-api:<version>
$ docker push hashicorpnomad/counter-api:<version>
```

Also build and publish for `arm64` (e.g. Graviton)

Publish a manifest for `v<n>`.

```bash
$ docker manifest create hashicorpnomad/counter-api:<v> --amend hashicorpnomad/counter-api:<v>-arm64 --amend hashicorpnomad/counter-api:<v>-amd64
$ docker manifest push hashicorpnomad/counter-api:<v>

Repeat for `counter-dashboard`.
```
