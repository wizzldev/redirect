# Path to URL redirect 🔗

A simple http server for redirecting given paths to the corresponding url.

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

### Error handling

Our service also allows you to specify a custom error page when a redirect does not exists or the
service gets a bad request. Just edit the `error.html` file's content, and you're ready to **Go**. (pun intended 😆)

## Testing

Every feature is well tested. You can simply run `go test ./...` to check if everything is fine.

## Credits

[Martin Binder](https://mrtn.vip) - The creator of [Wizzl](https://wizzl.app). An aspiring Go developer.
