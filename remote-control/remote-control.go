package remotecontrol

import (
	"fmt"
	"os/exec"
)

type RemoteControl struct {
}

func (rc RemoteControl) RunCommand(name string) error {
	cmd := fmt.Sprintf("dbus-send --system --print-reply --type=method_call --dest=org.gnome.ShairportSync '/org/gnome/ShairportSync' org.gnome.ShairportSync.RemoteControl.%s", name)

	err := exec.Command("bash", "-c", cmd).Run()

	return err
}

func (rc RemoteControl) Next() error {
	return rc.RunCommand("Next")
}

func (rc RemoteControl) Previous() error {
	return rc.RunCommand("Previous")
}
