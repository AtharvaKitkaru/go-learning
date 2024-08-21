package goroutines

import(
	"fmt"
	"sync"
	"time"
)

func fetchData(wg *sync.WaitGroup,url string){
	defer wg.Done()
	time.Sleep(5*time.Second)
	fmt.Printf("Fetching data from %s \n", url)
}

func CallApis () (){
	var apis []string
 	apis = []string{"api1","api2","api3"}

	var wg sync.WaitGroup
	
	for i:=0; i < len(apis); i++{
		wg.Add(1)
		go fetchData(&wg,apis[i])
	}


	wg.Wait()

	fmt.Println("Done fetching data")
	
	// alternative
	// for _, api := range apis{
	// 	go fetchData(api)
	// }
}
