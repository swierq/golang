# Echopoc

just a poc project for echo web framework but might be used to debug entra id oauth integration

## runing with cli

```
❯ bin/echopoc -h
NAME:
   echopoc - A new cli application

USAGE:
   echopoc [global options] command [command options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port value          port (default: 1323) [$ECHOPOC_PORT]
   --tenantID value      tenant ID for OAuth2 [$ECHOPOC_TENANT_ID]
   --clientID value      client ID for OAuth2 [$ECHOPOC_CLIENT_ID]
   --clientSecret value  client Secret for OAuth2 [$ECHOPOC_CLIENT_SECRET]
   --redirectURL value   redirect url for OAuth2 [$ECHOPOC_REDIRECT_URL]
   --appScope value      appScope url for OAuth2 [$ECHOPOC_APP_SCOPE]
   --help, -h            show help

```

## docker

```
docker run -it --rm \
-p 8080:8080 \
-e ECHOPOC_PORT=8080 \
-e ECHOPOC_TENANT_ID=xxxxxxxxxxxxxxx \
-e ECHOPOC_CLIENT_ID=fxxxxxxxxxxxxxxx \
-e ECHOPOC_CLIENT_SECRET=xxxxxxxxxxxxxxx \
-e ECHOPOC_REDIRECT_URL=http://localhost:8080/callback \
-e ECHOPOC_APP_SCOPE=api://xxxxxxxxxxxxxxx \
docker.io/swierq/golang echopoc

```
