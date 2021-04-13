# advice-app
A server application showing Go implementation of JSON-RPC 2.0.

In particular, the application retrieves advices about a given topic through [Advices Slip API](https://api.adviceslip.com/).
In addition, a Redis local cache is implemented, data is refreshed every 5 minutes.

## Usage

### Requirements

- Docker engine version >= 18.06.1
- Linux / Unix machine w/GNU make installed

**N.B.** The commands refer to bash and Make files, so please be sure you have bash installed on your machine and you can use `Make`.

### Run App
To start the application, run:
```bash
make start-api
```

The command will download Go and Redis images, get the dependencies, build the application and run it (in background).

To stop the application just run `make stop-api`.

The web server will listen on port 5000.

### Debugging

You can run the debug of the application by using the Debug panel in VSCode.

### Available endpoints

The JSON RPC server offers an HTTP `POST` `/api/advice` endpoint and a `RPCService.GiveMeAdvice` method which returns a list of advices from [Advices Slip API](https://api.adviceslip.com/) for a given topic and maximum amount.

#### Parameters:
- `topic` (string, mandatory): a topic for which is wanted to get advice.
- `amount` (integer, optional): a maximum amount of phrases to return.

### Examples

The [examples](https://github.com/marcocharlie/advice-app/tree/master/api/docs/examples) folder provides a REST file as example. It is meant to be used on [VSCode](https://code.visualstudio.com/) [REST Client plugin](https://github.com/Huachao/vscode-restclient).

## Automated tests

To test the application, run:
```bash
make test-api
```

## Application structure
The application is structured as follows:

- api (backend Go JSON-RPC application):
    - adapter: contains the logic for external web comunication (retrieving advices list from the external source).
    - cmd: contains the entrypoints of the main application and the logic to expose information on a JSON-RPC channel.
    - config: contains the default configuration file.
    - data: contains the logic to retrieve data from database (Redis).
    - internal: contains the internal logic.
    - models: contains the definitions of the entities involved.