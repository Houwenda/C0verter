package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

var (
	covertMode = "none"
	inTunnel   = ""
	outTarget  = "stdio/proxy"
	isProxy    = false
	crtPath    = ""
	keyPath    = ""
)

func init() {
	app := &cli.App{
		Name:  "C0verter",
		Usage: "Integrated tunneling tool with covert channel support.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "in",
				Value:    "",
				Aliases:  []string{"i"},
				Usage:    "`Address` inside the tunnel, local or remote address.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "out",
				Value:    "",
				Aliases:  []string{"o"},
				Usage:    "`Target` port outside the tunnel, will use command line IO if not defined, will work as proxy if --proxy is set to true.",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "covert",
				Value:    "",
				Aliases:  []string{"c"},
				Usage:    "Which `type` of covert channel to use, none as default.",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "proxy",
				Value:    false,
				Aliases:  []string{"p"},
				Usage:    "Use proxy, false by default.",
				Required: false,
			},
			&cli.StringFlag{
				Name:        "crt",
				Value:       "server.crt",
				Usage:       "Public-key file `path` when using https.",
				Required:    false,
				Destination: &crtPath,
			},
			&cli.StringFlag{
				Name:        "key",
				Value:       "server.key",
				Usage:       "Private-key file `path` when using https.",
				Required:    false,
				Destination: &keyPath,
			},
		},
		Action: func(c *cli.Context) error {
			// in
			if c.String("in") != "" {
				inTunnel = c.String("in")
				log.Println("param in:", inTunnel)
			}
			// out
			if c.String("out") != "" { // ip:port
				outTarget = c.String("out")
				log.Println("param out:", outTarget)
			} else { // standard IO
				log.Println("param out not defined, using command line IO or proxy")
			}
			// covert
			tmpCovert := strings.ToLower(c.String("covert"))
			if tmpCovert == "dns" || tmpCovert == "icmp" || tmpCovert == "http" { // dns/icmp/http
				covertMode = c.String("covert")
				log.Println("param covert:", covertMode)
			} else if tmpCovert == "https" { // https
				covertMode = c.String("covert")
				log.Println("param covert:", covertMode)
				log.Println("param crt:", crtPath)
				log.Println("param key:", keyPath)
			} else if tmpCovert == "" {
				log.Println("not using covert channel")
			} else { // undefined type
				log.Fatal("Undefined covert channel type")
			}
			// proxy
			if c.Bool("proxy") == true {
				isProxy = true
				log.Println("param proxy:", isProxy)
			} else {
				log.Println("not using proxy")
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: validate in-tunnel address
	// TODO: validate out-tunnel target(address and commands)
	// TODO: validate SSL path

}

func main() {
	log.Println("C0vert starts")
}
