# authnz [Educational Purpose Only]

Authorization and authentication. Learning go by writing a simple authentication and authorization service.

## Getting started

Run ``make rsa run`` to generate a rsa key and then start the server directly using go.
Run ``make build rsa start`` to build the binary, generate a rsa key and then start the build binary
Run ``make image`` to build the docker image.
Run ``make up`` to run the docker image. Note that you need to inject the private key to start the service.


## API

To see a list of available APIs see [api.http](api.http). You can use [Rest Client](https://github.com/Huachao/vscode-restclient) to directly invoke them on VSCode and its derivatives.


## Tasks

To see a list of tasks, run ``make help``. 

## Tests

There are few unit tests in the repo. Run ``make test`` to run them. However, there are no integration tests right now as this was written as an educational exercise.

## LICENSE

Apache 2 or MIT