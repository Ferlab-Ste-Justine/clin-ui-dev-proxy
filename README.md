# clin-ui-dev-proxy

## Description

Enable to run clin-frontend and clin-ui on the same port without configuring a proxy.

This solve authentification and iframe communication

## Requirements

You need either Go ([to install](https://go.dev/doc/install)) or Docker

Update clin-ui `.env.development` and add `PUBLIC_URL="/clinui-static"`

## Installation
### Command line

    go install github.com/Ferlab-Ste-Justine/clin-ui-dev-proxy@latest

### Docker

1. Build the image

        docker build --tag clin-ui-dev-proxy .

2. make sure you can see the image

        docker image ls

## Development

To build and run

      go run proxy.go

## Run proxy

    go install
    clin-ui-dev-proxy

it expect the following defaults (configurable)

```bash
> clin-ui-dev-proxy --help

  -clinui-host string
        clin-ui host name or ip plus the port if not 80 (default "http://0.0.0.0:2005")
  -clinui-staticpath string
        clin-frontend development static ressources url (default "/clinui-static")
  -frontend-host string
        clin-frontend host name or ip plus the port if not 80 (default "http://0.0.0.0:2002")
  -frontend-staticpath string
        clin-frontend development static ressources url (default "/static")
  -help
        Display default commands
  -port int
        Proxy Port. Normaly 2000. Auth should redirect there  (default 2000)
  -verbose
        Display more information, access files
```
