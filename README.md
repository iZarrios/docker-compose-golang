## Acessing docker container 

```bash
    docker compose run --service-ports web bash
```

**Notes: .env is just an example folder**

## Trying websocket in browser

### Connecting

```js
// where this url will work for my basic handler function only
var ws = new WebSocket("ws://localhost:3000/ws/123?v=1.0");
 

  ws.onmessage = function(e) {
    console.log('Message:', e.data);
  };

  ws.onclose = function(e) {
    console.log('Socket is closed. Reconnect will be attempted in 1 second.', e.reason);
    setTimeout(function() {
      connect();
    }, 1000);
  };

  ws.onerror = function(err) {
    console.error('Socket encountered error: ', err.message, 'Closing socket');
    ws.close();
  };

```

### Sending text to the server

```js
ws.send("hello world")
```
