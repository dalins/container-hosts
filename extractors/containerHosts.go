package extractors

import (
	"errors"
	"github.com/docker/docker/api/types"
)

type containerHosts struct {
}

func (l *containerHosts) HasFilterLabel() bool {
	return true
}

func (l *containerHosts) FilterLabel() string {
	return "container-hosts.enable=true"
}

func (l *containerHosts) HostnameFromContainer(container types.Container) (string, error) {
	hostName, err := l.hostNameFromLabels(container.Labels)
	if err != nil {
		err = errors.New(err.Error() + " for container '" + container.ID + "'")
	}

	return hostName, err
}

func (l *containerHosts) hostNameFromLabels(labels map[string]string) (string, error) {
	for key, value := range labels {
		if key == "container-hosts" {
			return l.hostNameFromLabel(value)
		}
	}

	return "", errors.New("no host name label found")
}

func (l *containerHosts) hostNameFromLabel(labelContent string) (string, error) {
	return labelContent, nil
}
