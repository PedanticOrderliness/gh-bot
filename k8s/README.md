# K8s files

TODO: Add Helm

### Secrets

The service requires that an access token and GH App Key PEM file be configured.

```
export BOT_ACCESS_TOKEN=$(echo -n <token> | base64 -w 0)
envsubst < k8s/secrets.yaml | kubectl apply -f -
```

Replace `<key file>` in k8s/configmap.yaml
```
kubectl apply -f k8s/configmap.yaml
```

### Deploy

```
docker build .
export TAG_NAME=$(docker images --format='{{.ID}}' | head -1)
docker tag $TAG_NAME docker.pedanticorderliness.com/gh-bot:$TAG_NAME
docker push docker.pedanticorderliness.com/gh-bot:$TAG_NAME
envsubst < k8s/deployment.yaml | kubectl apply -f -
```

### Setup service and ingress

```
kubectl apply -f k8s/service.yaml
kubectl apply -f k8s/ingress.yaml
```
