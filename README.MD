# Gotely
Gotely is a Go module that provides a convenient way to interact with the Telegram Bot API, allowing you to set up a Long Polling bot or a Webhook server.

---

## Installation

```sh
go get -u github.com/bigelle/gotely@latest
```

---

## Dependencies

Nope.

---

## Changelog

##### API Version 8.3
- v1.0.0: first release

--- 

## Usage

It's actually simple:

### Sending a request to the API

```go
// Preparing request body
sm := methods.SendMessage{
    ChatId: "@some_username", // or fmt.Sprint(12345678) - if you want to use a user ID
    Text: "deez nuts",
}
var msg objects.Message // Store the result here, or pass nil if you don't care
err := gotely.SendRequest(sm, &msg, "MY-TOP-SECRET-TOKEN") // Replace with an actual one
if err != nil {
    // Handle the error
}
// Now you can do something with the result
```

### Sending a request using custom types

Let's say you don’t want to use predefined types—maybe you prefer maps, or a method is currently unimplemented. As long as your type implements `gotely.Method`, you can send requests with it.

```go
type SendMessage map[string]any

func (sm SendMessage) Validate() error {
    // Validation logic
    return nil
}

func (sm SendMessage) Reader() io.Reader {
    buf := new(bytes.Buffer)
    json.NewEncoder(buf).Encode(sm)
    return buf
}

func (sm SendMessage) ContentType() string {
    return "application/json"
}

func (sm SendMessage) Endpoint() string {
    return "sendMessage" // Replaces <method> in the API URL
}

func main() {
    var resp map[string]any
    err := gotely.SendRequest(
        SendMessage{
            "chat_id": "@cool_username",
            "text": "deez nuts",
        },
        &resp,
        "MY-TOP-SECRET-TOKEN",
    )
    if err != nil {
        panic(err)
    }
    // Do something with resp
}
```

### Running a Long Polling bot

First, define a type that implements `tgbot.Bot`:

```go
type MyBot struct {
    token string
}

func (b MyBot) Token() string {
    return b.token
}
}

func (b MyBot) ApiUrl() string {
    // Return a custom API URL template if needed.
    return gotely.DEFAULT_URL_TEMPLATE
}

func (b MyBot) Client() *http.Client {
    return http.DefaultClient // Or a custom one
}

func (b MyBot) OnUpdate(objects.Update) error {
    // Define bot logic here
}
```

Now create an instance of `LongPollingBot`:

```go
lb := longpolling.New(MyBot{token: "MY-TOP-SECRET-TOKEN"}) // You can pass options here
lb.Start()
```

...and done! Now your bot will react to updates according to `OnUpdate(objects.Update) error`.

### Webhook server

The process is similar. First, tell Telegram where to send updates, then define your bot and pass it to the constructor:

```go
swh := webhook.SetWebhook{
    Url: "https://example.com/webhook", // The URL where Telegram will send updates
}
var ok bool
err := gotely.SendRequestWith(swh, &ok, "MY-TOP-SECRET-TOKEN")
if err != nil {
    panic(err)
}
if !ok {
    panic("Telegram API responded with false")
}

hook := webhook.New(MyBot{token: "MY-TOP-SECRET-TOKEN"}) // You can pass options here
if err := hook.Start(); err != nil {
    panic(err)
}
```

By default, it listens on port 8080 and expects updates at `/webhook`, but you can configure this.

### Webhook with an existing server

The simplest way. If you already have a webhook server running, no matter what language or framework, just send a `setWebhook` request:

```go
swh := webhook.SetWebhook{
    Url: "https://example.com/webhook", // The URL where Telegram will send updates
}
var ok bool
err := gotely.SendRequestWith(swh, &ok, "MY-TOP-SECRET-TOKEN")
if err != nil {
    panic(err)
}
if !ok {
    panic("Telegram API responded with false")
}
```

...and that's it.

---

## License

This project is licensed under the MIT License, meaning you are free to use, modify, and distribute it as long as you retain the original copyright notice.

See the full text in the [LICENSE](./LICENSE) file.