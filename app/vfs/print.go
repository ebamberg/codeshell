package vfs

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func PrintDirectoryTree[T VFSEntry](sl []T, rowMapper func(T) (int, string)) {
	treeData := pterm.LeveledList{}
	for _, row := range sl {
		level, text := rowMapper(row)
		treeData = append(treeData, pterm.LeveledListItem{Level: level, Text: text})
	}
	// Convert the leveled list into a tree structure.
	tree := putils.TreeFromLeveledList(treeData)
	tree.Text = "Directory" // Set the root node text.

	// Render the tree structure using the default tree printer.
	pterm.DefaultTree.WithRoot(tree).Render()
}
