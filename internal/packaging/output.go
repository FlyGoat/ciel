package packaging

import (
	d "ciel/display"
	"ciel/internal/abstract"
	"ciel/proc-api"
	"os"
	"path"
	"syscall"
)

const (
	OutputPath = "/debs"
)

type Tree struct {
	Parent   abstract.Ciel
	BasePath string
}

func (t *Tree) Mount(mountPoint string) {
	if _, err := os.Stat(t.BasePath); os.IsNotExist(err) {
		err = os.MkdirAll(t.BasePath, 0755)
		if err != nil {
			return
		}
	}
	outputMountPoint := path.Join(mountPoint, OutputPath)
	os.MkdirAll(outputMountPoint, 0755)
	syscall.Mount(t.BasePath, outputMountPoint, "", syscall.MS_BIND, "")
}

func (t *Tree) Unmount(mountPoint string) {
	outputMountPoint := path.Join(mountPoint, OutputPath)
	if _, err := os.Stat(outputMountPoint); os.IsNotExist(err) {
		return
	}
	if !proc.Mounted(outputMountPoint) {
		return
	}
	d.ITEM("unmount output")
	err := syscall.Unmount(outputMountPoint, 0)
	d.WARN(err)
	if err != nil {
		return
	}
	d.ITEM("remove output mount point")
	err = os.Remove(outputMountPoint)
	d.WARN(err)
}

func (t *Tree) MountHandler(i abstract.Instance, mount bool) {
	if mount {
		t.Mount(i.MountPoint())
	} else {
		t.Unmount(i.MountPoint())
	}
}