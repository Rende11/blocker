package pkg

import (
	"errors"
	"log"

	"github.com/goodhosts/hostsfile"
)

const target = "0.0.0.0"

func getHostsFile(path string) (*hostsfile.Hosts, error) {
	if path == "" {
		return hostsfile.NewHosts()
	}
	return hostsfile.NewCustomHosts(path)
}

func BlockSites(sites []string, path string) {

	if len(sites) == 0 {
		log.Fatal(errors.New("No sites for blocking"))
	}

	hosts, err := getHostsFile(path)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err := hosts.Add(target, sites...); err != nil {
		log.Fatal(err.Error())
	}

	if err := hosts.Flush(); err != nil {
		log.Fatal(err.Error())
	}
}

func UnBlockSites(sites []string, path string) {

	if len(sites) == 0 {
		log.Fatal(errors.New("No sites for blocking"))
	}

	hosts, err := getHostsFile(path)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err := hosts.Remove(target, sites...); err != nil {
		log.Fatal(err.Error())
	}

	if err := hosts.Flush(); err != nil {
		log.Fatal(err.Error())
	}
}

func GetHosts(path string) string {
	hosts, err := getHostsFile(path)

	if err != nil {
		log.Fatal(err.Error())
	}

	return hosts.String()
}
