package megaclient

import (
	"github.com/t3rm1n4l/go-mega"
)

const (
	PATH_WIDTH = 50
	SIZE_WIDTH = 10
)

func download() string {
	client := mega.New()

	client.SetAPIUrl("")
	children, _ := client.FS.GetChildren(client.FS.GetRoot())

	for i := range children {

	}
}
