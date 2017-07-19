package gogen

import (
	"testing"

	"github.com/glossina/gotify"
	"github.com/stretchr/testify/require"
)

func TestConstNameFromContent(t *testing.T) {
	g := NewGenerator(gotify.New(nil), nil)
	require.Equal(t, "const12", g.constNameFromContent("12"))
	require.Equal(t, "const12", g.constNameFromContent("12"))
	require.Equal(t, "done200", g.constNameFromContent("Done 200"))
	require.Equal(t, "done200", g.constNameFromContent("Done 200"))
	require.Equal(t, "done200Case2", g.constNameFromContent("Done  200"))
	require.Equal(t, "done200Case2", g.constNameFromContent("Done  200"))
	require.Equal(t, "comma", g.constNameFromContent(", "))
	require.Equal(t, "comma", g.constNameFromContent(", "))
	require.Equal(t, "commaAimsidEq", g.constNameFromContent(",aimsid="))
	require.Equal(t, "commaAimsidEq", g.constNameFromContent(",aimsid="))
}
