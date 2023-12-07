package composite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	imageEditor := NewImageEditor()
	imageEditor.load()

	assert.Equal(t, 2, len(imageEditor.all.children))

	components := []InterfaceGraphic{imageEditor.all.children[0], imageEditor.all.children[1]}
	imageEditor.groupSelected(components)

	assert.Equal(t, 1, len(imageEditor.all.children))
}
