package Walker

import (
	"os"
	"github.com/Inkeliz/InkDir/Directory"
	"github.com/Inkeliz/InkDir/Types"
	"github.com/Inkeliz/InkDir/Display"
	"time"
	"fmt"
	"github.com/Inkeliz/InkDir/Setup"
)

type Search struct {
	Main string
	Current string
}

func start(s Search, ws chan Search, c chan *Types.Info) error {
	info, err := os.Lstat(s.Current)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		r := make(Types.Info)
		r[s.Main] = info.Size()

		c <- &r
		return nil
	}

	go more(s, ws, c)

	return nil
}

func more(s Search, ws chan Search, c chan *Types.Info) error {

	names, err := Directory.All(s.Current)
	if err != nil {
		return err
	}

	for _, name := range names {
		r := Search{}
		r.Main = s.Main

		r.Current = Directory.CreatePath(s.Current, name)

		ws <- r
	}

	return nil
}

func Start(ws chan Search, c chan *Types.Info, f chan int) {

FOR:
	for {
		select {
		case s := <-ws:
			start(s, ws, c)
		case <-time.After(time.Millisecond * 100):
			f <- 1
			break FOR
		}
	}

}

func Listener(names []string, c chan *Types.Info, e chan *Types.Info) {
	result := make(Types.Info)

	for _, name := range names {
		result[name] = 0
	}

	for r := range c {
		for index, size := range *r {
			result[index] += size
		}
	}

	e <- &result
}

func Wait(Param Setup.Parameters, ws chan Search, c chan *Types.Info, e chan *Types.Info, f chan int){
	finished := 0

	for range f {
		finished++
		if finished == Param.LimitCPU {
			close(c)
			close(f)
		}
	}

	Display.Result(*<-e)
	close(e)
	close(ws)
	Display.Path(Param.StartPath)
	Display.Time(time.Since(Param.Time).String())

	if Param.ForceWait {
		WaitType := ""
		fmt.Scanf("%s", &WaitType)
	}
}