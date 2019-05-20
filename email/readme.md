# Email Service Example

A simple example gRPC based email service that sends emails via SMTP. Connections are kept open whilst sending email then closed for after 30 seconds of inactivity.

```
service Email {
  rpc Send (EmailRequest) returns (EmailResponse) {}
}
```

## Config

You can set SMTP details with the following env vars.

```
SMTP_HOST=smtp.sendgrid.net
SMTP_PORT=2525
SMTP_USERNAME=apikey
SMTP_PASSWORD=somepassword
```

Set `DEBUG=true` to see timing and connection information
