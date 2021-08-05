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

type destination struct {
	to       string
	distance int
}

var routeRegexp = regexp.MustCompile(`(\w+) to (\w+) = (\d+)`)

func Day9Part1(input string) (int, error) {
	return findBestRouteDistance(input, math.MaxInt32, func(left, right int) bool {
		return left < right
	})
}

func Day9Part2(input string) (int, error) {
	return findBestRouteDistance(input, 0, func(left, right int) bool {
		return left > right
	})
}

func findBestRouteDistance(input string, initDistance int, isLeftBetter func(left, right int) bool) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	routes := make([]route, len(lines))
	currentRoute := route{distance: initDistance}
	for i, line := range lines {
		route, err := parseRoute(line)
		if err != nil {
			return 0, err
		}
		routes[i] = route

		if isLeftBetter(route.distance, currentRoute.distance) {
			currentRoute = route
		}
	}

	visited := make(map[string]bool)
	visited[currentRoute.from] = true
	visited[currentRoute.to] = true

	for {
		found := false
		next := route{distance: initDistance}
		for _, route := range routes {
			if visited[route.from] && visited[route.to] {
				continue
			}

			if currentRoute.connectsTo(route) && isLeftBetter(route.distance, next.distance) {
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
			currentRoute.from = next.to
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
	return route{from: matches[1], to: matches[2], distance: distance}, nil
}
