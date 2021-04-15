# advice-app
A simple server application showing Go implementation of JSON-RPC 2.0. A light Node Js client for API consumption is implemented.

In addition, a Redis local cache is implemented, data is refreshed every 5 minutes.

In particular, the application retrieves advice about a given topic through [Advice Slip API](https://api.adviceslip.com/). 

## Usage

### Requirements

In order to simplify the execution, a convenient Makefile is provided.

Docker and `make` is all you need to run both server and client.

### Run Server
To start the application, run:
```bash
make start-api
```

The command will download Go and Redis images, get the dependencies, build the application and run it (in background).

To stop the application just run `make stop-api`.

The web server will listen on port 5000.

#### Automated tests

To test the application, run:
```bash
make test-api
```

#### Debugging

You can run the debug of the application by using the Debug panel in VSCode.

#### Available endpoints

The JSON RPC server offers an HTTP `POST` `/api/advice` endpoint and a `RPCService.GiveMeAdvice` method which returns a list of advice from [Advice Slip API](https://api.adviceslip.com/) for a given topic and maximum amount.

##### Parameters:
- `topic` (string, mandatory): a topic for which is wanted to get advice.
- `amount` (integer, optional): a maximum amount of phrases to return.

#### Examples

The [examples](https://github.com/marcocharlie/advice-app/tree/master/api/docs/examples) folder provides a REST file as example. It is meant to be used on [VSCode](https://code.visualstudio.com/) [REST Client plugin](https://github.com/Huachao/vscode-restclient).

### Run Client
To start the application, run:
```bash
make start-client
```

The command will download Node Js image, get the dependencies, build the application and run it (in background).

Once the app starts, open `http://localhost:8080` in your browser, where you will be able to make requests to server.

To stop the application just run `make stop-client`.

## Application structure
The application is structured as follows:

- api (backend Go JSON-RPC application):
    - adapter: contains the logic for external web comunication (retrieving advice list from the external source).
    - cmd: contains the entrypoints of the main application and the logic to expose information on a JSON-RPC channel.
    - config: contains the default configuration file.
    - data: contains the logic to retrieve data from database (Redis).
    - internal: contains the internal logic.
    - models: contains the definitions of the entities involved.
- client (frontend Node Js application)