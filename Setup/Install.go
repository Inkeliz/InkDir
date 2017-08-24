package Setup

import (
	"os"
	"os/exec"
	"github.com/Inkeliz/InkDir/Directory"
	"errors"
)

func isAdmin() error {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return errors.New("This action need administrator rights, open it as a admin.")
	}
	return nil
}

func (Setup *Parameters) SetPath() error {

	err := isAdmin()
	if err != nil {
		return err
	}

	path, err := Directory.Local()
	if err != nil {
		return err
	}

	cmd := exec.Command(`REG`, `ADD`, `HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment`, `/v`, `InkDir`, `/t`, `REG_EXPAND_SZ`, `/d`, Directory.ClearDir(path), `/f`)
	err = cmd.Start()

	return err
}

func (Setup *Parameters) SetMenu() error {

	err := isAdmin()
	if err != nil {
		return err
	}

	path, err := Directory.Local()
	if err != nil {
		return err
	}

	RegDirectory := []string{`HKCR\Directory\Background\shell\InkDir`, `HKCR\Directory\shell\InkDir`}

	for _, Reg := range RegDirectory {
		cmd := exec.Command(`REG`, `ADD`, Reg, `/v`, `MUIVerb`, `/t`, `REG_SZ`, `/d`, `InkDir`, `/f`)
		if err = cmd.Start(); err != nil {
			return err
		}

		cmd = exec.Command(`REG`, `ADD`, Reg, `/v`, `SubCommands`, `/t`, `REG_SZ`, `/d`, `InkDir.All;InkDir.Folders;InkDir.Files`, `/f`)
		if err = cmd.Start(); err != nil {
			return err
		}
	}

	RegAction := map[string]string{
		"All": "-files=true -folders=true",
		"Folders": "-files=false -folders=true",
		"Files": "-files=true -folders=false",
	}

	for name, action := range RegAction {
		cmd := exec.Command(`REG`, `ADD`, `HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\CommandStore\shell\InkDir.` + name, `/d`, `List ` + name, `/f`)
		if err = cmd.Start(); err != nil {
			return err
		}

		cmd = exec.Command(`REG`, `ADD`, `HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\CommandStore\shell\InkDir.` + name + `\command`, `/d`, path+" -wait -path %V " + action, `/f`)
		if err = cmd.Start(); err != nil {
			return err
		}
	}

	return err
}

