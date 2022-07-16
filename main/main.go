package main

import (
	"ex.sov/arrays"
	"ex.sov/closing-channels"
	"ex.sov/closures"
	"ex.sov/constants"
	custom_sorting "ex.sov/custom-sorting"
	"ex.sov/embedding"
	"ex.sov/exdefer"
	"ex.sov/exerrors"
	"ex.sov/exfunctions"
	"ex.sov/exgenerics"
	"ex.sov/exmethods"
	"ex.sov/exmutex"
	"ex.sov/expointers"
	"ex.sov/exrange"
	"ex.sov/exrunes"
	"ex.sov/exstructs"
	"ex.sov/forfor"
	"ex.sov/hello"
	"ex.sov/ifelse"
	"ex.sov/interfaces"
	"ex.sov/mapex"
	"ex.sov/multy"
	"ex.sov/nonblock"
	"ex.sov/rangeOverChannels"
	"ex.sov/recursion"
	"ex.sov/slisess"
	"ex.sov/sorting"
	"ex.sov/switchsov"
	"ex.sov/values"
	"ex.sov/variables"
	"ex.sov/variadic"
	"fmt"
)

func main() {
	fmt.Println("Main0!")
	hello.Hello()
	values.Values()
	variables.Main()
	constants.Main()
	forfor.Main()
	ifelse.Main()
	switchsov.Main()
	arrayssov.Main()
	//testmath.Main()
	slisess.Main()
	mapex.Main()
	exrange.Main()
	exfunctions.Main()
	multy.Main()
	variadic.Main()
	closures.Main()
	recursion.Main()
	expointers.Main()
	exrunes.Main()
	exstructs.Main()
	exmethods.Main()
	interfaces.Main()
	embedding.Main()
	exgenerics.Main()
	exerrors.Main()
	// Долго отрабатывают, ассинхронность, задержки
	//goroutines.Main()
	//exchannels.Main()
	//buffering.Main()
	//exsync.Main()
	//directions.Main()
	// Рандом конкуренси
	//exselect.Main()
	// Долго отрабатывают, ассинхронность, задержки
	//extimeouts.Main()
	nonblock.Main()
	closing_channels.Main()
	rangeOverChannels.Main()
	// таймеры
	//extimer.Main()
	// задержка (таймер)
	//extickers.Main()
	// Мои любимые воркеры <3
	//workers.Main()
	//wait_groups.Main()
	//exrate.Main()
	//ex_atomic.Main()
	exmutex.Main()
	//stateful.Main()
	sorting.Main()
	custom_sorting.Main()
	// panic
	//expanic.Main()
	exdefer.Main()

}
