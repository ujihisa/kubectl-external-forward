// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/int128/kubectl-socat/pkg/cmd"
	"github.com/int128/kubectl-socat/pkg/portforwarder"
)

// Injectors from di.go:

func NewCmd() cmd.Interface {
	portForwarder := &portforwarder.PortForwarder{}
	cmdCmd := &cmd.Cmd{
		PortForwarder: portForwarder,
	}
	return cmdCmd
}