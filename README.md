# ======> MAILER SERVICE <========
golang project provided to handle mailer service and has been integrated with `rabbitmq`, on channel `mailer_queue` and templateing with file `.html`, for provider sending email using smtp2go for the testing with template of `.env` is ready on this project,

the project using `enqueue` concept that set in file `rabbitmq/rabbitmq.go` will be automatically send `message` after consumming message from channel set on file `.env`

for example of list template email to send you can see in file `templates/*_template.html` for the lsit of template that will using message with variable `tag` on the message from channel for list of `tag` you can see at the file `handlers/handlers.go` and here this of list the `tag`
    
    1. login
    2. send_code
    3. reset_password
    4. api_key

that must be the same name on file `templates/*_template.html`, and if you want to customize the project you can update the file list

    1. handlers -> file to handle tag of message from consuming message from queue
    2. interfaces -> file to handling interfaces message like all variable they have
    3. mailes -> file to handling send mail with configuration template from message
    4. rabbitmq -> file to handling configuration of queue and subcribe to channel has been declared on rabbitmq queue
    5. templates -> file to handling templating style to show for user
    6. token -> file to handling decrypted message jwt from queue channel

flow of system on this project just like this :

the flow of system is started from your project

```
-> publish message json to rabbitmq channel when has been encrypted with jwt private key (example of json you see after this)
-> mailer service subcribe to channel and consume the message
-> message will decrypted first to check the tag 
-> system will generate template of message with using json message after decryption
-> system will send message to target email
```

project will using dual decrypted when `first decrypted` is for to get `tag` group that using `Record` interfaces and then project will decrypt again using file in `interfaces` to get list of json messages `when is suitable on tag and handlers`

```
example of json data
{
    record: {
        email: string, // email address user
        tag: string, // tag is for description to generate file
        subject: string // subject of message to subject on email send
    },
    iss: 'name_project',
    iat: <Time on Integer>,
    exp: <Time on Integer>,
    jti: SecureRandom.hex(10)
}
```

## ======> SETUP <========
Setup for this project is simple, you must prepare the configuration on file `.env`, configuration channel on the `rabbitmq`, the list of configuration you must following the step :
    
1. RabbitMQ

configuration rabbitmq you can follow this link `https://www.rabbitmq.com/download.html`

or you want to easly install RabbitMQ configuration you can running command on your terminal using

```bash
$ cd config && docker-compose up -Vd
```
after your config for rabbitmq installation don't forget to update sthe file `.env` on variable `RABBITMQ_HOST` with host to your `rabbitmq` has `installed`

2. Channel configuration

Channel configration you can using file in `config/rabbitmq_channel.yaml` on your project because this project just for consuming from config channel when you set in rabbitmq
you can create channel in rabbitmq in `exchange` and `queue` and don't forget to set the `queue_name` when you has been config in file `.env`

3. SMTP

Mailer using `github.com/go-mail/mail` for the package and for the configuration you can set in file `.env` on list variable in below
```
SMTP_HOST={smtp host }
SMTP_PORT={ smtp port }
SMTP_USER={ smtp username }
SMTP_PASS={ smtp password }
SMTP_SENDER_EMAIL={ smtp sender email }
SMTP_SENDER_NAME={ smttp sender name }
```
don't forget to change that `smtp configuration` with your smtp mailer you have, but on `sender email` must using `activated email` like `alone@gmail.com` or something else, if you want to customize that configuration you can change in file `mailers/mailers.go`

4. JWT

configuration jwt is easy, you just need to `generate` the `jwt_public_key` and `jwt_private_key` and set the `algorithm` on file `.env`, you can follow the step :

- Generate Private Key
```bash
$ openssl genpkey -algorithm RSA -out private_key.pem
```

- Generate Public Key
```bash
$ openssl rsa -in private_key.pem -pubout -out public_key.pem
```

but in this project you must to convert that file output to `base64` configuration `private_key` or `public_key`
- Convert private key to base64
```bash
$ openssl base64 -in private_key.pem
```
- Convert Public key to base64
```bash
$ openssl rsa -in private_key.pem -pubout -outform DER | openssl base64 -A
```

after that you can set that `private_key` on your project to protect the message from your project for executing on this `mailer-service`, and for `mailer-service` you can must to using that `public_key` to decrypted the message from your `project` after consuming channel has you declared in file `.env`

setup notification is using firebase with get server key
the base config of notification is on folder `notificaiton/notification.go`

format payload on jwt encode is
```json
{
  "record": {
    "users": {
      "email": "<email_user>",
      "device_token": [
        "<string>"
      ]
    },
    "title": "<string>",
    "message": "<string>"
  },
  "sub": "<string>",
  "aud": "'<string>'",
  "iss": "<string>",
  "iat": "timestamp",
  "exp": "timestamp",
  "jti": "generateString"
}
```

## ======> RUN PROJECT <========

project mailer-service you can run using the following command

download all depenencies project using the following command
```bash
$ go mod download
```

completely setup configuration for your project

after all is ready, run the following command

#### Running mailer is on
```bash
$ go run mailer.go
```
#### Running notification is on
```bash
$ go run push_notif.go
```