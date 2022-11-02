package extractors

import (
	"errors"
	"regexp"

	"github.com/docker/docker/api/types"
)

type traefik struct {
}

func (l *traefik) HasFilterLabel() bool {
	return true
}

func (l *traefik) FilterLabel() string {
	return "traefik.enable=true"
}

func (l *traefik) HostnameFromContainer(container types.Container) (string, error) {
	hostName, err := l.hostNameFromLabels(container.Labels)
	if err != nil {
		err = errors.New(err.Error() + " for container '" + container.ID + "'")
	}

	return hostName, err
}

func (l *traefik) hostNameFromLabels(labels map[string]string) (string, error) {
	regexStr := `^traefik\.(.+)\.routers\.(.+)\.rule$`
	regex := regexp.MustCompile(regexStr)
	for key, value := range labels {
		if regex.MatchString(key) {
			return l.hostNameFromLabel(value)
		}
	}

	return "", errors.New("no host name label found")
}

func (l *traefik) hostNameFromLabel(labelContent string) (string, error) {
	regexStr := `^Host\(\s*\W((([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9]))\W\s*\)$`
	regex := regexp.MustCompile(regexStr)
	submatches := regex.FindStringSubmatch(labelContent)
	if len(submatches) > 2 {
		return submatches[1], nil
	} else {
		return "", errors.New("failed to parse host name")
	}
}
