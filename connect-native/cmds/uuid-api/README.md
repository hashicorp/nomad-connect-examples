# uuid-api


#### Build

```bash
DOCKER_BUILDKIT=1 docker build --no-cache --ssh default -t hashicorpnomad/uuid-api:<version> .
```

#### Publish

```bash
docker login
docker push hashicorpnomad/uuid-api:<version>
```
