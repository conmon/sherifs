package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/conmon/sherifs/settings"
	"github.com/conmon/sherifs/version"
)

var (
	versionFlag  = flag.Bool("version", false, "Prints the version and exits")
	hostnameFlag = flag.String("hostname", "", "Define the local system hostname. Default: from /etc/hostname")
)

func main() {
	var err error
	flag.Parse()

	// Print the version and quit
	if *versionFlag {
		fmt.Printf("sherif version %s%s, %s\n", version.Version, version.Release, version.GoVersion)
		os.Exit(0)
	}

	// Set the hostname
	if len(*hostnameFlag) == 0 {
		settings.Hostname, err = os.Hostname()
		if err != nil {
			fmt.Printf("Sherif failed: unable to get hostname: %s\n", err)
		}
	} else {
		settings.Hostname = *hostnameFlag
	}

	fmt.Println(settings.Hostname)
}
