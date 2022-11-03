# Overview
container-hosts extracts host names from containers and combines them with an ip address.

## Motivation
Reduce the management effort required for name resolution when adding or removing containers.

## Current features
- Extraction of the host of a traefik router
- Create and update a linux hosts file

## Planned features
- Extraction of host names from a self defined label (No dependency to traefik)
- Extraction of container name to use it as host name

## Tested with following services as DNS provider:
- [blocky](https://0xerr0r.github.io/blocky/) as [Hosts file](https://0xerr0r.github.io/blocky/configuration/#hosts-file)
- [dnsmasq](https://thekelleys.org.uk/dnsmasq/doc.html) as [addn-hosts](https://thekelleys.org.uk/dnsmasq/docs/dnsmasq-man.html)
