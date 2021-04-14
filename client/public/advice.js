
function giveMeAdvices(topic, amount) {

    let apiUrl = "http://localhost:5000/api/advice"

    var inputData = {
        id: 1,
        jsonrpc: "2.0",
        method: "RPCService.GiveMeAdvice",
        params: {
            "topic": topic,
            "amount": parseInt(amount)
        }
    }

    console.log(inputData)

    fetch(apiUrl, {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        method: 'POST',
        body: JSON.stringify(inputData),
        dataType: 'json',
    }).then(response => response.json())
        .then(data => {
            if (data.result != null) {
                console.log('Success:', data);
                var myObject = data.result.adviceList
                var parsed = "";
                for (i = 0; i < myObject.length; i++) {
                    var myobj = myObject[i];
                    parsed += (i + 1 + " : " + myobj + "\n\n");
                }
                $("#display").val(parsed);
            } else {
                confirm("An error occurred:\n\n" + data.error.message)
            }
        })
        .catch((error) => {
            console.error('Error:', error);
            confirm("An error occurred:\n\n" + error)
        });
}