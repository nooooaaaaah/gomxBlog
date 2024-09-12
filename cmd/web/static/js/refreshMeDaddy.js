function setupWebSocket() {
  var ws = new WebSocket("ws://192.168.88.214:6900/refreshMeDaddy"); // Replace with your server's address and port

  ws.onmessage = function (event) {
    if (event.data === "reload") {
      setTimeout(function () {
        window.location.reload();
      }, 1000); // Wait one second before reloading
    }
  };

  ws.onclose = function () {
    console.log("WebSocket closed. Attempting to reconnect...");
    setTimeout(setupWebSocket, 1000); // Attempt to reconnect after a delay
  };

  ws.onerror = function (err) {
    console.error("WebSocket encountered an error:", err);
    ws.close(); // Ensure WebSocket is closed after an error
  };
}

setupWebSocket();
