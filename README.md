# GOTP

A Time-based One-Time-Password tool for Go.

### Installing

```
go get github.com/josegomezr/gotp
```

After that, just use the `gotp` command. 

```
$ gotp -h
NAME:
   gotp - A TOTP Server & Utility

USAGE:
   gotp [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     server, s   Runs GOTP HTTP server
     check, c    Checks a TOTP code
     gencode, g  Generates a TOTP code
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port value, -p value  GOTP HTTP Server Port (default: "2444") [$PORT, $GOTP_PORT]
   --host value            GOTP HTTP Server Host (default: "localhost") [$HOST, $GOTP_HOST]
   --secret value          Secret key to check/generate the code against
   --code value            Code to be checked
   --help, -h              show help
   --version, -v           print only the version

```

## Server API

The server expose 3 endpoints:

- `/code` Given a Secret Key or Data Payload it will return a pin CODE.
- `/qr` Given a Secret Key or Data Payload it will return a QR image to use it with the Authenticator app.
- `/verify` Given a Secret Key or Data Payload and a PIN code it will perform the check and output the result.

All the endpoints can:
- Respond to **GET** and **POST** verbs.
- Recieve input from:
  - QueryString (only GET)
  - JSON
  - XML
  - multipart/formdata 
  - plain old form data.
- Output in JSON, but it can be changed using the `Accept` header to:
    - Plain Text (`text/plain`)
    - XML (`text/xml` or `application/xml`)


## Secret vs Data Payload

The correct use case for an OTP authentication would be to use a secret key, which consist of a Base32 string. But to provide ease of integration, **gotp** can use an arbitary data (up to 128 bytes) as secret code, internally it'll:

- Hash the contents (`md5` right now, although it could by any hash algo)
- Base32-encode the result

And use the Base32 result as the secret key for then on.

## Built With

* [dgryski/dgoogauth](https://github.com/dgryski/dgoogauth) - HOTP/TOTP Algorithm Implementation.
* [gin-gonic/gin](https://github.com/gin-gonic/gin) - golang HTTP web framework
* [urfave/cli](https://github.com/gin-gonic/gin) - CLI library
* [stretchr/testify](https://github.com/stretchr/testify) - Testing framework

## Contributing

PR's are welcome, I have literally no standard nor procedure. If your PR makes sense to me and all test pass, you're good to go.

## Acknowledgments

* SUSE's Hackweek: [GH](https://github.com/topics/hackweek), [Twitter](https://twitter.com/search?q=%23Hackweek&src=typed_query)
