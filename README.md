# Path to URL redirect 🔗

## Usage

### Configuration

Define your routes in `.data/redirects.json` like this example:
```json
{
  "/instagram": "https://instagram.com",
  "/discord": "https://discord.com"
}
```
This simple http server will redirect all routes to the corresponding url.

### Startup

You can start the server with ease by running `docker compose up`.
You can also download `go` (tested with go1.22.6) and run `make run` to start the server, or `make build` to create
a single binary from the code (you'll have to add the `.data` folder to the correct path).

### Flags

Our service allows you to specify the server's listen address by using the `--listenAddr :3001` flag.

## Testing

Every feature is well tested. You can simply run `go test ./...` to check if everything is fine.
