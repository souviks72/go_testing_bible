package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterUnique(t *testing.T) {
	input := []Developer{
		Developer{Name: "Elliot"},
		Developer{Name: "Elliot"},
		Developer{Name: "David"},
		Developer{Name: "Alexander"},
		Developer{Name: "Eva"},
		Developer{Name: "Alan"},
	}

	expected := []string{
		"Elliot",
		"David",
		"Alexander",
		"Eva",
		"Alan",
	}

	assert.Equal(t, expected, FilterUnique(input))
}
