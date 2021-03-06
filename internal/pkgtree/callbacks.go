package pkgtree

import (
	"github.com/AOSC-Dev/ciel/internal/abstract"
)

func (t *Tree) MountHandler(i abstract.Instance, mount bool) {
	if mount {
		t.Mount(i.MountPoint())
	} else {
		t.Unmount(i.MountPoint())
	}
}
