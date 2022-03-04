package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the command line app ready to run
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "command line app"
	app.Usage = "search IPs and web server names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "search IPs in web",
			Flags:  flags,
			Action: searchIps,
		},

		{
			Name:   "server",
			Usage:  "search name server in web",
			Flags:  flags,
			Action: searchServer,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func searchServer(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host) // name serve web
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
