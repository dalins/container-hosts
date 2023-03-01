# Overview
container-hosts extracts host names from containers and combines them with an ip address.

## Motivation
Reduce the management effort required for name resolution when adding or removing containers.

## Current features
- Extraction of the host of a traefik router
- Extraction of host names via a self defined label (container-hosts / no traefik needed)
- Create and update a linux hosts file

## Planned features
- Extraction of container name to use it as host name

## Example configuration for container-hosts extractor
This example uses the container-hosts extractor to create a hosts file:
````yaml
---
services:
  container-hosts:
    build: .
    container_name: container-hosts
    environment:
      - TZ=Europe/Berlin
      - HOST_IP=192.168.178.2
      - EXTRACTOR=ContainerHosts
    restart: unless-stopped
    volumes:
      - data:/data
      - /var/run/docker.sock:/var/run/docker.sock:ro

  busybox:
    image: busybox
    command: sleep 30
    container_name: busybox
    labels:
      - container-hosts.enable=true
      - container-hosts=busybox

version: '3.5'
volumes:
  data: null
...
````

Result:
````ini
cat /data/hosts 
192.168.178.2	busybox
````

## Tested with following services as DNS provider:
- [blocky](https://0xerr0r.github.io/blocky/) as [Hosts file](https://0xerr0r.github.io/blocky/configuration/#hosts-file)
- [dnsmasq](https://thekelleys.org.uk/dnsmasq/doc.html) as [addn-hosts](https://thekelleys.org.uk/dnsmasq/docs/dnsmasq-man.html)
