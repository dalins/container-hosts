package main

import (
	"container-hosts/extractors"
	"context"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func extractHostNames(extractor extractors.Extractor, ctx context.Context, cli *client.Client) []string {
	opts := types.ContainerListOptions{All: true}
	if extractor.HasFilterLabel() {
		opts.Filters = filters.NewArgs()
		opts.Filters.Add("label", extractor.FilterLabel())
	}
	containers, err := cli.ContainerList(ctx, opts)
	if err != nil {
		panic(err)
	}

	var hostNames []string
	for _, container := range containers {
		hostName, err := extractor.HostnameFromContainer(container)
		if err == nil {
			hostNames = append(hostNames, hostName)
		} else {
			log.Fatal(err.Error())
		}
	}

	return hostNames
}

func writeHostsFile(hostFilePath string, hostIp string, hostNames []string) {
	f, err := os.OpenFile(hostFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for _, hostName := range hostNames {
		line := hostIp + "\t" + hostName
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	hostIpEnvName := "HOST_IP"
	hostIp, envIsSet := os.LookupEnv(hostIpEnvName)
	if !envIsSet {
		panic(hostIpEnvName + " is not set")
	}

	hostsFilePathEnvName := "HOSTS_FILEPATH"
	hostsFilePath, envIsSet := os.LookupEnv(hostsFilePathEnvName)
	if !envIsSet {
		panic(hostsFilePathEnvName + " is not set")
	}

	extractorEnvName := "EXTRACTOR"
	extractorName, extractorIsSet := os.LookupEnv(extractorEnvName)
	if !extractorIsSet {
		panic(extractorEnvName + " is not set")
	}

	extractor, err := extractors.CreateExtractor(extractorName)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	writeHostsFile(hostsFilePath, hostIp, extractHostNames(extractor, ctx, cli))

	eventOpts := types.EventsOptions{}
	eventOpts.Filters = filters.NewArgs()
	if extractor.HasFilterLabel() {
		eventOpts.Filters.Add("label", extractor.FilterLabel())
	}
	eventOpts.Filters.Add("type", "container")
	msgs, errs := cli.Events(ctx, eventOpts)
	for {
		select {
		case err := <-errs:
			print(err)
		case msg := <-msgs:
			if msg.Action == "create" || msg.Action == "destroy" {
				writeHostsFile(hostsFilePath, hostIp, extractHostNames(extractor, ctx, cli))
			}
		}
	}
}
