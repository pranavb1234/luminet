# luminet
Implement a basic TCP server in Go to listen for incoming connections.
Implement a TCP client to connect to the local server.
Forward traffic between the remote client and the local server.
Implement an HTTP server to handle incoming HTTP requests.
Forward HTTP requests to the local server and return responses to the remote client.
Add WebSocket support using a library like gorilla/websocket.
Design a protocol for communication between the client and server (e.g., JSON over WebSocket or raw TCP).
Implement authentication (e.g., API keys or tokens) to secure the connection.
Add heartbeat/ping-pong mechanism to keep connections alive.
Implement subdomain allocation for public endpoints (e.g., user123.jprq.io).
Use a DNS library or API to dynamically create subdomains.
Map subdomains to specific tunnels.
Add TLS/SSL support to secure traffic between the client and server.
Implement rate limiting to prevent abuse.
Add logging and monitoring to detect suspicious activity.
Write unit tests for core functionality (e.g., tunneling, HTTP forwarding).
Write integration tests to test the client-server interaction.
Test edge cases (e.g., high traffic, connection drops).
Containerize the project using Docker.
Deploy the server to a cloud platform (e.g., AWS, GCP, or Heroku).
``` Plain Text
            Public Client
                 │
                 ▼
        🌍 Public Internet
                 │
                 ▼
        🌐 Tunneling Server
            - Listens for public connections
            - Forwards data between phone & local server
                 │
                 ▼
        🏠 Local Network
                 │
                 ▼
        📞 Local TCP Server
            - Listens on port 8000
            - Receives & responds to messages
```