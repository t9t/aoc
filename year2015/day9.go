package year2015

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type route struct {
	from, to string
	distance int
}

func (r route) connectsTo(other route) bool {
	return r.from == other.from || r.from == other.to || r.to == other.from || r.to == other.to
}

func (r route) hasShorterDistanceThan(other route) bool {
	return r.distance < other.distance
}

const ridiculouslyLongDistance = math.MaxInt32

type destination struct {
	to       string
	distance int
}

var routeRegexp = regexp.MustCompile(`(\w+) to (\w+) = (\d+)`)

func Day9Part1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	routes := make([]route, len(lines))
	currentRoute := route{distance: ridiculouslyLongDistance}
	for i, line := range lines {
		route, err := parseRoute(line)
		if err != nil {
			return 0, err
		}
		routes[i] = route

		if route.hasShorterDistanceThan(currentRoute) {
			currentRoute = route
		}
	}

	visited := make(map[string]bool)
	visited[currentRoute.from] = true
	visited[currentRoute.to] = true

	for {
		found := false
		next := route{distance: ridiculouslyLongDistance}
		for _, route := range routes {
			if visited[route.from] && visited[route.to] {
				continue
			}

			if currentRoute.connectsTo(route) && route.hasShorterDistanceThan(next) {
				next = route
				found = true
			}
		}
		if !found {
			break
		}

		visited[next.from] = true
		visited[next.to] = true
		currentRoute.distance += next.distance
		if currentRoute.from == next.from {
			currentRoute.from = currentRoute.to
		} else if currentRoute.from == next.to {
			currentRoute.from = next.from
		} else if currentRoute.to == next.from {
			currentRoute.to = next.to
		} else {
			currentRoute.to = next.from
		}
	}

	return currentRoute.distance, nil
}

func parseRoute(s string) (route, error) {
	matches := routeRegexp.FindStringSubmatch(s)
	if len(matches) != 4 {
		return route{}, fmt.Errorf("invalid route %q", s)
	}
	distance, err := strconv.Atoi(matches[3])
	if err != nil {
		return route{}, fmt.Errorf("invalid distance in route %q: %w", s, err)
	}
	return route{
		from:     matches[1],
		to:       matches[2],
		distance: distance,
	}, nil
}

func Day9Part2(input string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}
