package tree

import (
	"errors"
	"sort"
)

type Node struct {
	ID       int
	Children []*Node
}

type Record struct {
	ID     int
	Parent int
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	// Sort the records so we can enforce some rules.
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make([]Node, len(records))

	for idx, record := range records {
		nodes[idx].ID = record.ID
		if record.Parent == 0 && idx == 0 && record.ID == 0 {
			continue
		}

		if err := validate(idx, record); err != nil {
			return nil, err
		}

		parent := &nodes[record.Parent]
		parent.Children = append(parent.Children, &nodes[idx])
	}
	return &nodes[0], nil
}

func validate(idx int, record Record) error {
	if idx != record.ID {
		return errors.New("gap or duplicate in record IDs")
	}
	if record.ID == 0 {
		return errors.New("non-root node has root ID")
	}
	if record.ID <= record.Parent {
		return errors.New("invalid parent ID")
	}
	return nil
}


