package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID int
	Parent int
}

type Node struct {
	ID int
	Children []*Node
}

func Build(r []Record) (*Node, error) {

	if len(r) == 0 {
		return nil, nil
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i].ID < r[j].ID
	})

	if r[0].ID != 0 || r[0].Parent != 0 {
		return nil, errors.New("invalid root node")
	}

	nodes := make([]*Node, len(r))

	for i, record := range r {

		if i != record.ID {
			return nil, errors.New("sequence of IDs is broken")
		}

		newNode := &Node{ID: record.ID}
		nodes[record.ID] = newNode

		if record.ID == 0 {
			continue
		}

		if record.Parent >= record.ID {
			return nil, errors.New("invalid parent id")
		}

		parentNode := nodes[record.Parent]
		parentNode.Children = append(nodes[record.Parent].Children, newNode)
	}

	return nodes[0], nil
}