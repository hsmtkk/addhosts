package main

import (
	"fmt"
	"log"

	"github.com/hsmtkk/addhosts/getip"
	"github.com/spf13/cobra"
)

func main() {
	ipv6 := false
	command := &cobra.Command{
		Use:  "addhosts hostnames...",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			run(args, ipv6)
		},
	}
	command.Flags().BoolVar(&ipv6, "ipv6", false, "prefer IPv6")
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}

func run(hostnames []string, preferIPv6 bool) {
	getter := getip.New()
	for _, hostname := range hostnames {
		var ip string
		var err error
		if preferIPv6 {
			ip, err = getter.GetIP(hostname, getip.IPv6)
		} else {
			ip, err = getter.GetIP(hostname, getip.IPv4)
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s %s\n", hostname, ip)
	}

}
