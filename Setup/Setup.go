package Setup

import (
	"flag"
	"github.com/Inkeliz/InkDir/Directory"
	"runtime"
	"errors"
	"time"
)

type Parameters struct {
	StartPath string
	IncludeFiles, IncludeFolders bool
	LimitCPU, LimitQueue int
	ForceWait bool
	InstallPath, InstallMenu bool
	Time time.Time
}

func Init() Parameters {

	var Param Parameters
	
	flag.StringVar(&Param.StartPath, "path", "","Path to directory.")

	flag.BoolVar(&Param.IncludeFiles, "files", false, "If set to true, we'll check files on the given path.")
	flag.BoolVar(&Param.IncludeFolders, "folders", true, "If set to true, we'll check folders on the given path.")

	flag.IntVar(&Param.LimitCPU, "cpu", runtime.NumCPU(), "Set a limit of CPU cores available to use.")
	flag.IntVar(&Param.LimitQueue, "queue", 10000, "Set a limit of the queue, may reduce RAM but can create instability.")

	flag.BoolVar(&Param.ForceWait, "wait", false, "If set to true, we'll keep the window open until you close.")

	flag.BoolVar(&Param.InstallMenu, "installMenu", false, "If exist this flag, we'll install the ContextMenu. *")
	flag.BoolVar(&Param.InstallPath, "installPath", false, "If exist this flag, we'll set the Environment Variable. *")

	//TODO:
	/*
	- Output : Save result into a file
	- Format : Output format ([ASCII TABLE LIST]/JSON/XML)
	- Order : Create order based on [SIZE] (ASC/[DESC]) or ALPHABETIC (ASC/DESC)
	 */
	flag.Parse()

	Param.StartPath = Directory.CreatePath(Param.StartPath, "")
	Param.Time = time.Now()

	runtime.GOMAXPROCS(Param.LimitCPU)

	return Param
}


func (Setup *Parameters) InitDirNames() ([]string, error){

	switch {
	case Setup.IncludeFiles && Setup.IncludeFolders:
		return Directory.All(Setup.StartPath)
	case Setup.IncludeFiles:
		return Directory.Files(Setup.StartPath)
	case Setup.IncludeFolders:
		return Directory.Folders(Setup.StartPath)
	}

	return nil, errors.New("Oppps! What you need to do with this given path? Say `-h`. :)")
}