# Native javascript client

This example was built using only javascript and [webpack](https://webpack.js.org/), and the connecting with the blockchain through the [web3](https://github.com/ethereum/web3.js/) API.

## Build the example
```
npm run build
```

The command above will generate the `dist` directory with your application. We use webpack to bundle all the dependencies and generate only one javascript (i.e. `app.js`) that is used in the `index.html`.

## Running the example
```
// Serves the front-end on http://localhost:8080
npm run dev
```
