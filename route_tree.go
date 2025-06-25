package routetree

import (
	"encoding/json"
	"errors"
	"strings"
)

type RouteNode struct {
	Resource     interface{}
	IsQueryParam bool
	Segment      string
	Nodes        []RouteNode
}

type RouteTree struct {
	Root RouteNode
}

func NewRouteTree() *RouteTree {
	return &RouteTree{
		Root: RouteNode{
			Resource:     nil,
			IsQueryParam: false,
			Segment:      "/",
			Nodes:        []RouteNode{},
		},
	}
}

func (rt *RouteTree) AddRoute(path string, resource interface{}) error {
	parts, err := splitPath(path)
	if err != nil {
		return err
	}
	currentNode := &rt.Root

	for i, part := range parts {
		found := false
		for j, node := range currentNode.Nodes {
			if strings.HasPrefix(part, ":") || strings.HasPrefix(node.Segment, ":") {
				p1 := strings.TrimPrefix(part, ":")
				p2 := strings.TrimPrefix(node.Segment, ":")
				if p1 == p2 || node.IsQueryParam {
					return errors.New("conflict with existing route segment: " + node.Segment + " vs " + part)
				}
			}
			if node.Segment == part {
				currentNode = &currentNode.Nodes[j]
				found = true
				break
			}
		}

		if !found {
			newNode := RouteNode{
				IsQueryParam: strings.HasPrefix(part, ":"),
				Segment:      part,
				Nodes:        []RouteNode{},
			}
			if i == len(parts)-1 {
				newNode.Resource = resource // Set resource only for the last segment
			}
			currentNode.Nodes = append(currentNode.Nodes, newNode)
			currentNode = &currentNode.Nodes[len(currentNode.Nodes)-1]
		}
	}

	return nil
}

func (rt *RouteTree) FindRoute(path string) (interface{}, map[string]string, error) {
	parts, err := splitPath(path)
	if err != nil {
		return nil, nil, err
	}
	currentNode := &rt.Root
	params := make(map[string]string)

	for _, part := range parts {
		found := false
		for i, node := range currentNode.Nodes {
			if node.IsQueryParam {
				params[strings.TrimPrefix(node.Segment, ":")] = part
			}
			if node.Segment == part {
				currentNode = &currentNode.Nodes[i]
				found = true
				break
			}
		}

		if !found {
			return nil, params, errors.New("route not found")
		}
	}

	if currentNode.Resource == nil {
		return nil, params, errors.New("route found but no resource associated")
	}

	return currentNode.Resource, params, nil
}

func (rt *RouteTree) DumpTree() {
	t := rt.Root
	data, _ := json.MarshalIndent(t, "", "  ")
	println(string(data))
}

func splitPath(path string) ([]string, error) {
	if path == "" || path == "/" {
		return []string{"/"}, nil
	}

	// Remove leading and trailing slashes
	if path[0] == '/' {
		path = path[1:]
	}
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	// Split by slashes
	return append([]string{"/"}, splitBySlash(path)...), nil
}

func splitBySlash(path string) []string {
	parts := []string{}
	currentPart := ""

	for _, char := range path {
		if char == '/' {
			if currentPart != "" {
				parts = append(parts, currentPart)
				currentPart = ""
			}
		} else {
			currentPart += string(char)
		}
	}

	if currentPart != "" {
		parts = append(parts, currentPart)
	}

	return parts
}
