const { request, response } = require('express');
const fetch = require('node-fetch');

const dotenv = require('dotenv');
dotenv.config();
const apiBasehUrlPath = process.env.API_BASE_URL_PATH;

var express = require('express');
var router = express.Router();

// get route to home page
router.get('/', function (req, res, next) {
  res.render('index', { title: 'Express' });
});

// post route for data request to backend
router.post('/api', async (request, response) => {

  const data = request.body
  const topic = data.topic
  const amount = data.amount

  // init input data for request
  var inputData = {
    id: 1,
    jsonrpc: "2.0",
    method: "RPCService.GiveMeAdvice",
    params: {
      "topic": topic,
      "amount": parseInt(amount)
    }
  }

  // do request to backend
  const fetch_response = await fetch(apiBasehUrlPath, {
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    method: 'POST',
    body: JSON.stringify(inputData),
    dataType: 'json',
  });
  const json = await fetch_response.json();
  response.json(json)
})



module.exports = router;
