package main

import (
	"fmt"
	"go-trials/Embed_Types"
	"go-trials/FileNFolder"
	"go-trials/HashNCrypto"
	"go-trials/function"
	"go-trials/go_routine_trial"
	"go-trials/http_client_server"
	"go-trials/list_trial"
	"go-trials/map_trial"
	"go-trials/mascot"
	"go-trials/morestrings"
	"go-trials/rpc_client_server"
	"go-trials/slice"
	"go-trials/sort_trial"
	"go-trials/struct_trial"
	"go-trials/tcp_client_server"
	"os"

	"rsc.io/quote"
)

func main() {

	switch arg := os.Args[1]; arg {
	case "mascot":
		fmt.Println(mascot.BestMascot())
		fmt.Println(quote.Go())
	case "string":
		fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	case "slice":
		slice.Slice_trial()
		slice.Slice_append()
		slice.Slice_copy()
	case "map":
		map_trial.Map_basic()
		map_trial.Map_elements()
		map_trial.Map_of_Maps_elements()
	case "closure":
		add := func(x, y int) int {
			return x + y
		}

		fmt.Printf("add(1, 1):%d\n", add(1, 1))

		nextEven := function.MakeEvenGenerator()
		fmt.Println("Even Numbers from Even generator")
		fmt.Println(nextEven()) // 0
		fmt.Println(nextEven()) // 2
		fmt.Println(nextEven()) // 4
	case "function":
		xs := []float64{98, 93, 77, 82, 83}
		fmt.Printf("Average:%f\n", function.Average(xs))

		ints := []int{1, 2, 3}
		fmt.Printf("Addition result:%d\n", function.Add(ints...))

		fmt.Printf("Factorial of %d is %d\n", 5, function.Factorial(5))

		function.DeferFunction()
		function.PanicFunction()
		function.PointerTest()
		function.NewFunctionTest()

		value, isEven := function.Half(1)
		fmt.Printf("Half of 1:%d, isEven:%t\n", value, isEven)
		value, isEven = function.Half(2)
		fmt.Printf("Half of 1:%d, isEven:%t\n", value, isEven)

		squareValue := 1.5
		function.Square(&squareValue)
		fmt.Println(squareValue)

		swapx := 1
		swapy := 2

		function.Swap(&swapx, &swapy)
		fmt.Printf("After swap, swapx: %d, swapy: %d\n", swapx, swapy)
	case "struct":
		struct_trial.Circle_Test()
		struct_trial.Rectangle_Test()
		struct_trial.TotalArea_Test()
	case "embedType":
		Embed_Types.Embed_Test()
	case "FileNFolder":
		FileNFolder.FileOpenTest()
		FileNFolder.FileOpenUsingIOUtil()
		FileNFolder.ListFilesInDir()
		FileNFolder.ListFilesInDirRecurse()
	case "list":
		list_trial.ListTest()
	case "sort":
		sort_trial.SortPersonByName()
		sort_trial.SortPersonByAge()
	case "HashNCrypto":
		HashNCrypto.CRC_Test()
		HashNCrypto.HashTest()
		HashNCrypto.SHA1_Test()
	case "tcp":
		tcp_client_server.ClientServerTest()
	case "http":
		http_client_server.Http_client_server_test()
	case "rpc":
		rpc_client_server.RPC_client_server_test()
	case "goroutine":
		go_routine_trial.Run_Func_Concurrently()
	case "channel":
		go_routine_trial.Channel_Test()
		go_routine_trial.Biggest_Home_Page_Size()
	case "channelSelect":
		go_routine_trial.Channel_Select_Test()
	default:
		panic("un-recognized command")
	}
}
