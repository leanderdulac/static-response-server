# Static Response Server

A very simple static HTTP server with support for web-sockets.

- Any request, independent of content, will always return a static body and content type.
- Any message sent from a web-socket client will be replied with same static response.
- Visit `/.ws` for a basic UI to connect and send web-socket messages.
- The `CONTENT_TYPE` and `CONTENT_BODY` defines response's content type and body.
- The `PORT` environment variable sets the server port.
- No TLS support yet :(

To run as a container:

```
docker run --env CONTENT_TYPE=application/json \
           --env CONTENT_BODY='{"upstream": "production"}' \
           --detach -P pagarme/static-response-server
```

Also works great with `docker-compose`:

```yaml
services:
  production:
    image: pagarme/static-response-server
    environment:
      - PORT=8080
      - CONTENT_TYPE=application/json
      - CONTENT_BODY={"upstream":"production"}
    ports:
      - "8080:8080"

  sandbox:
    image: pagarme/static-response-server
    environment:
      - PORT=8081
      - CONTENT_TYPE=application/json
      - CONTENT_BODY={"upstream":"sandbox"}
    ports:
      - "8081:8081"
```
