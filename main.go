package main

import (
	"sync"

	concurrency "github.com/jaswanth-gorripati/go-examples/concurrency"
	goInterface "github.com/jaswanth-gorripati/go-examples/interfaces"
	httpServer "github.com/jaswanth-gorripati/go-examples/webServers"
	log "github.com/sirupsen/logrus"
)

type GoExamples struct {
	goInterface.Shape
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	// Example 1 -- interface
	goInterface.BasicInterfaceExample()
	ci := &goInterface.Circle{Radius: 23}

	log.Debugf("The calculated area : %v", goInterface.GetArea(ci))

	// Example 2 -- Print Interface
	txt := goInterface.TextPrinter{Text: "This Message is from Main function"}
	goInterface.PrintTextFromInterface(&txt)

	// Example - 3 -- Sortable Interface
	var sortableIntCollection goInterface.Sortable
	ele := goInterface.IntCollection{Elements: []int{1, 3, 4, 2, 5, 2, 3, 42, 4, 5, 1, 2, 3}}
	sortableIntCollection = &ele

	log.Debugf(" Length Of sortable elements : %v", sortableIntCollection.Len())
	log.Debugf(" Compare the elements in collection: %v", sortableIntCollection.Less(0, 1))
	log.Debugf("Before Swap of elements : %v", ele.Elements)

	sortableIntCollection.Swap(0, 1)
	log.Debugf("After Swap of elements : %v", ele.Elements)

	concurrency.SequentialPrint()
	concurrency.ConcurrentPrint()
	concurrency.BasicUsagePrint()
	concurrency.ConcurrentSummation()

	urls := concurrency.WebURLS{Urls: concurrency.WEB_URLs}
	urls.GetAllStatus()

	var wg sync.WaitGroup
	wg.Add(1)
	go httpServer.StartHttpServer(&wg)
	wg.Add(1)
	go httpServer.StartGorillaMux(&wg)
	wg.Wait()

}
