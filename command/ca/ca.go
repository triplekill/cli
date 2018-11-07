package ca

import (
	"github.com/smallstep/cli/command"
	"github.com/smallstep/cli/command/ca/provisioner"
	"github.com/urfave/cli"
)

// init creates and registers the ca command
func init() {
	cmd := cli.Command{
		Name:      "ca",
		Usage:     "initialize and manage a certificate authority",
		UsageText: "step ca <subcommand> [arguments] [global-flags] [subcommand-flags]",
		Description: `**step ca** command group provides facilities initialize a certificate
authority, sign and renew certificate, ...

## Examples

Create the configuration for a new certificate authority:
'''
$ step ca init
'''

Download the root_ca.crt:
'''
$ step ca root root_ca.crt \
  --ca-url https://ca.smallstep.com \
  --fingerprint 0d7d3834cf187726cf331c40a31aa7ef6b29ba4df601416c9788f6ee01058cf3
'''

Create a new certificate using a token:
'''
$ TOKEN=$(step ca new-token internal.example.com)
$ step ca new-certificate internal.example.com internal.crt internal.key \
  --token $TOKEN --ca-url https://ca.smallstep.com --root root_ca.crt
'''

Renew the certificate while is still valid:
'''
$ step ca renew internal.crt internal.key \
  --ca-url https://ca.smallstep.com --root root_ca.crt
'''

Configure the ca-url and root in the environment:
'''
$ cp root_ca.crt $STEPPATH/secrets/
$ cat \> $STEPPATH/config/defaults.json
{
    "ca-url": "https://ca.smallstep.com",
    "root": "/home/user/.step/secrets/root_ca.crt"
}
'''`,
		Subcommands: cli.Commands{
			initCommand(),
			newTokenCommand(),
			newCertificateCommand(),
			signCertificateCommand(),
			rootComand(),
			renewCertificateCommand(),
			provisioner.Command(),
		},
	}

	command.Register(cmd)
}

// common flags used in several commands
var (
	caURLFlag = cli.StringFlag{
		Name:  "ca-url",
		Usage: "<URI> of the targeted Step Certificate Authority.",
	}

	rootFlag = cli.StringFlag{
		Name:  "root",
		Usage: "The path to the PEM <file> used as the root certificate authority.",
	}

	fingerprintFlag = cli.StringFlag{
		Name:  "fingerprint",
		Usage: "The <fingerprint> of the targeted root certificate.",
	}

	tokenFlag = cli.StringFlag{
		Name: "token",
		Usage: `The one-time <token> used to authenticate with the CA in order to create the
certificate.`,
	}

	notBeforeFlag = cli.StringFlag{
		Name: "not-before",
		Usage: `The <time|duration> set in the NotBefore (nbf) property of the token. If a
<time> is used it is expected to be in RFC 3339 format. If a <duration> is
used, it is a sequence of decimal numbers, each with optional fraction and a
unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns",
"us" (or "µs"), "ms", "s", "m", "h".`,
	}

	notAfterFlag = cli.StringFlag{
		Name: "not-after",
		Usage: `The <time|duration> set in the Expiration (exp) property of the token. If a
<time> is used it is expected to be in RFC 3339 format. If a <duration> is
used, it is a sequence of decimal numbers, each with optional fraction and a
unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns",
"us" (or "µs"), "ms", "s", "m", "h".`,
	}

	provisionerKidFlag = cli.StringFlag{
		Name:  "kid",
		Usage: "The provisioner <kid> to use.",
	}

	provisionerIssuerFlag = cli.StringFlag{
		Name:  "issuer",
		Usage: "The provisioner <name> to use.",
	}

	passwordFileFlag = cli.StringFlag{
		Name: "password-file",
		Usage: `The path to the <file> containing the password to decrypt the one-time token
generating key.`,
	}
)