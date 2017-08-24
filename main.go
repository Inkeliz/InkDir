package main

import (
	"github.com/Inkeliz/InkDir/Directory"
	"github.com/Inkeliz/InkDir/Walker"
	"github.com/Inkeliz/InkDir/Types"
	"github.com/Inkeliz/InkDir/Display"
	"github.com/Inkeliz/InkDir/Setup"
)

func main() {

	Param := Setup.Init()

	switch {
	case Param.InstallPath:
		Display.Error(Param.SetPath())
		Display.Install()
		return
	case Param.InstallMenu:
		Display.Error(Param.SetMenu())
		Display.Install()
		return
	}

	channelResult := make(chan *Types.Info)
	channelSearch := make(chan Walker.Search, Param.LimitQueue)
	channelFinish := make(chan int)
	channelEnd := make(chan *Types.Info)

	names, err := Param.InitDirNames()
	Display.Error(err)

	go Walker.Listener(names, channelResult, channelEnd)

	for i := 0; i < Param.LimitCPU; i++ {
		go Walker.Start(channelSearch, channelResult, channelFinish)
	}

	for _, name := range names {
		r := Walker.Search{}
		r.Main = name
		r.Current = Directory.CreatePath(Param.StartPath, name)

		channelSearch <- r
	}

	Walker.Wait(Param, channelSearch, channelResult, channelEnd, channelFinish)

}
