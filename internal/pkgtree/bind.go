package pkgtree

import (
	"os"
	"path"
	"path/filepath"
	"syscall"

	d "github.com/AOSC-Dev/ciel/display"
	"github.com/AOSC-Dev/ciel/proc-api"
)

const (
	TreePath = "/tree"
)

func (t *Tree) Mount(mountPoint string) {
	if _, err := os.Stat(t.BasePath); os.IsNotExist(err) {
		return
	}
	treeMountPoint := path.Join(mountPoint, TreePath)
	os.MkdirAll(treeMountPoint, 0755)
	if !proc.Mounted(treeMountPoint) {
		syscall.Mount(t.BasePath, treeMountPoint, "", syscall.MS_BIND, "")
	}
}

func (t *Tree) Unmount(mountPoint string) {
	treeMountPoint := path.Join(mountPoint, TreePath)
	if _, err := os.Stat(treeMountPoint); os.IsNotExist(err) {
		return
	}
	if !proc.Mounted(treeMountPoint) {
		return
	}
	d.ITEM("unmount tree")
	result, err := filepath.Abs(treeMountPoint)
	if err != nil {
		d.WARN(err)
		return
	}
	for proc.Mounted(result) {
		err = syscall.Unmount(result, syscall.MNT_FORCE)
		if err != nil {
			d.WARN(err)
			return
		}
	}
	d.WARN(err)
	if err != nil {
		return
	}
	d.ITEM("remove tree mount point")
	err = os.Remove(treeMountPoint)
	d.WARN(err)
}
