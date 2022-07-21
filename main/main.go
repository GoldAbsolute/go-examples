package main

import (
	"ex.sov/arrays"
	"ex.sov/closing-channels"
	"ex.sov/closures"
	"ex.sov/constants"
	custom_sorting "ex.sov/custom-sorting"
	"ex.sov/embedding"
	"ex.sov/exbase64"
	"ex.sov/exdefer"
	"ex.sov/exepoch"
	"ex.sov/exerrors"
	"ex.sov/exfunctions"
	"ex.sov/exgenerics"
	"ex.sov/exhashes"
	"ex.sov/exjson"
	"ex.sov/exmethods"
	"ex.sov/exmutex"
	"ex.sov/expointers"
	"ex.sov/exrange"
	"ex.sov/exrecover"
	"ex.sov/exregular"
	"ex.sov/exrunes"
	"ex.sov/exstructs"
	"ex.sov/extime"
	"ex.sov/exxml"
	"ex.sov/forfor"
	"ex.sov/hello"
	"ex.sov/ifelse"
	"ex.sov/interfaces"
	"ex.sov/mapex"
	"ex.sov/multy"
	"ex.sov/nonblock"
	number_parsing "ex.sov/number-parsing"
	random_numbers "ex.sov/random-numbers"
	"ex.sov/rangeOverChannels"
	"ex.sov/recursion"
	"ex.sov/slisess"
	"ex.sov/sorting"
	string_format "ex.sov/string-format"
	string_library "ex.sov/string-library"
	"ex.sov/switchsov"
	text_templates "ex.sov/text-templates"
	time_parsing "ex.sov/time-parsing"
	url_parsing "ex.sov/url-parsing"
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
	// defer и запись/чтение/закрытие файла
	exdefer.Main()
	exrecover.Main()
	string_library.Main()
	string_format.Main()
	text_templates.Main()
	exregular.Main()
	// Кто-то из двух ниже грузит(скорее второй)
	exjson.Main()
	exxml.Main()
	extime.Main()
	exepoch.Main()
	time_parsing.Main()
	random_numbers.Main()
	number_parsing.Main()
	url_parsing.Main()
	exhashes.Main()
	exbase64.Main()

}
