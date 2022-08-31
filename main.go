package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/netip"
	"os"

	"github.com/huin/goupnp/dcps/ocf/internetgateway2"
)

func main() {

	internal := flag.String("i", "", "internal ip:port")
	external := flag.String("e", "", "external ip:port")
	duration := flag.Uint("d", 30, "duration in seconds")
	name := flag.String("n", "", "rule description/name")
	udp := flag.Bool("u", false, "is udp, otherwise tcp")

	flag.Parse()

	internalAddr, err := parseAddPort(*internal)
	if err != nil {
		fmt.Println("Please provide a valid internal ip:port, error:", err)
		os.Exit(1)
	}

	externalAddr, err := parseAddPort(*external)
	if err != nil {
		fmt.Println("Please provide a valid external ip:port, error:", err)
		os.Exit(1)
	}

	durationSeconds := uint32(*duration)

	ctx := context.Background()

	client, err := PickRouterClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	externalIP, err := client.GetExternalIPAddress()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your external IP address is: ", externalIP)

	var protocol string
	switch *udp {
	case true:
		protocol = "UDP"
	default:
		protocol = "TCP"
	}

	if err := client.AddPortMapping(
		externalAddr.Addr().String(),
		externalAddr.Port(),
		protocol,
		internalAddr.Port(),
		internalAddr.Addr().String(),
		true,
		*name,
		durationSeconds,
	); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stdout, "Mapped external %s:%d to internal %s:%d\n", externalAddr.Addr(), externalAddr.Port(), internalAddr.Addr(), internalAddr.Port())
}

func parseAddPort(s string) (*netip.AddrPort, error) {
	if s == "" {
		return nil, fmt.Errorf("empty string")
	}

	addrPort, err := netip.ParseAddrPort(s)
	if err != nil {
		return nil, err
	}

	return &addrPort, nil
}

func PickRouterClient(ctx context.Context) (*internetgateway2.WANIPConnection1, error) {
	ip1Clients, _, err := internetgateway2.NewWANIPConnection1Clients()
	if err != nil {
		return nil, err
	}

	return ip1Clients[0], nil
}
