# uuid-fe


#### Build

```bash
DOCKER_BUILDKIT=1 docker build --ssh default -t hashicorpnomad/uuid-fe:<version> .
```

#### Publish

```bash
docker login
docker push hashicorpnomad/uuid-fe:<version>
```
