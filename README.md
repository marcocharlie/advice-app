# advice-app
A simple application offering:
- Go implementation of JSON-RPC 2.0.
- a light frontend Node.js server for API consumption.

In addition, a Redis local cache is implemented, data is refreshed every 5 minutes.

In particular, the application retrieves advice about a given topic through [Advice Slip API](https://api.adviceslip.com/). 

## Usage

### Requirements

In order to simplify the execution, a convenient Makefile is provided.

Docker and `make` is all you need to run both backend and frontend apps.

### Run Application
To start the application, run:
```bash
make start-app
```

The command will download Go, Redis and Node.js images, get the dependencies, build the application and run it (in background).

The backend web server will listen on port 5000.

Once the app starts, open `http://localhost:8080` in your browser, where you will be able to make requests to backend server (suggested queries: `life`,`love`,`cars`,`spiders`,`wine`,`lover`,`smile`).

To stop the application just run `make stop-app`.

#### Backend automated tests

To test the application, run:
```bash
make test-backend
```

#### Backend Debugging

You can run the debug of the application by using the Debug panel in VSCode.

#### Backend available endpoints

The JSON RPC server offers an HTTP `POST` `/api/advice` endpoint and a `RPCService.GiveMeAdvice` method which returns a list of advice from [Advice Slip API](https://api.adviceslip.com/) for a given topic and maximum amount.

##### Parameters:
- `topic` (string, mandatory): a topic for which is wanted to get advice.
- `amount` (integer, optional): a maximum amount of phrases to return.

#### Examples

The [examples](https://github.com/marcocharlie/advice-app/tree/master/api/docs/examples) folder provides a REST file as example. It is meant to be used on [VSCode](https://code.visualstudio.com/) [REST Client plugin](https://github.com/Huachao/vscode-restclient).

## Application structure

The GO application structure is inspired by the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).
The Node.js application skeleton was created through the [express-generator](https://expressjs.com/it/starter/generator.html) tool.

The application is structured as follows:

- api (backend Go JSON-RPC application):
    - adapter: contains the logic for external web comunication (retrieving data from the external source).
    - cmd: contains the entrypoints of the main application and the logic to expose information on a JSON-RPC channel.
    - config: contains the default configuration file.
    - data: contains the logic to retrieve data from database (Redis).
    - internal: contains the internal logic.
    - models: contains the definitions of the entities involved.
- app (frontend Node.js application):
    - the `package.json` file defines the application dependencies and other information.
    - bin: stores the application entry point `/bin/www` which sets up some error handling and then loads `app.js` to do the rest of the work. 
    - routes: stores the app routes.
    - views: containns the templates. 
    - public: contains the `.css` and other useful files