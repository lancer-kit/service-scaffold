# notifier
This package is used for sending notifications to NATS.

## Sender
`Sender` is a structure that works with the notification service. The structure contains `*Config`, `*nats.Conn` and `*logrus.Entry`.

#### Method list:
- `Send(*Message) error` – main functional of `Sender`. Publish message to service.
- `SetConfig(*Config) *Sender` – set new configuration.
- `SetLogger(*logrus.Entry) *Sender` – set logger.
- `ErrorLog(error, string)` – write error message into log.
- `IsConnected() error` – check connection to NATS.
- `Disconnect() *Sender` – disconnec.

## SenderI
`SenderI` is an interface for `Sender`.
#### Method list:
- `SetLogger(*logrus.Entry) *Sender`
- `SetConfig(*Config) *Sender`
- `Disconnect() *Sender`
- `Send(*Message) error`

## Config
`Config` is a structure that contains configuration for `Sender`. It has the following fields:  

| Parameter | Type   | Description                               |
|-----------|--------|-------------------------------------------|
| Chanel    | string | Chanel (topic) name. Default = "notifier" |
| Url       | string | NATS connection url                       |
| User      | string | NATS user                                 |
| Password  | string | NATS password                             |
| Token     | string | NATS auth token                           |

## Message
`Message` – main message structure. Used to send notifications to `socket-notifier`.  
`Message` has the following fields:

| Parameter   | Type        | Description                  | JSON name   |
|-------------|-------------|------------------------------|-------------|
| UserId      | int         | User ID for notifications    | userId      |
| Command     | string      | "Command" for front-end      | command     |
| IsBroadcast | bool        | Private or broadcast message | isBroadcast |
| Data        | interface{} | Optional data                | Data        |

`Command` field shows which type of update and which structure is used in `Data` field. Here is the list of possible commands:

- "operation" – update of operation status. Data: WSOperationMessage.
- "payment" – update of payment status. Data: WSPaymentMessage.
- "transaction" – update of transaction status. Data: WSTxMessage.
