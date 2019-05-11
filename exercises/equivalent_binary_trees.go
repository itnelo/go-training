package exercises

import (
	"fmt"
	"sync"
	_ "sync/atomic"

	"golang.org/x/tour/tree"
)

var (
	lock = sync.RWMutex{}
)

// channel can be used as an alternative to mutex
// ch <- true
// *work*
// <- ch
func concurrentMapWrite(m map[int]int, key, delta int) {
	lock.Lock()
	defer lock.Unlock()

	if m[key] += delta; 0 == m[key] {
		delete(m, key)
	}
}

func walkBranch(
	branch *tree.Tree,
	valuesChannel chan<- int,
	waitGroup *sync.WaitGroup,
) {
	defer waitGroup.Done()

	if nil == branch {
		return
	}

	valuesChannel <- branch.Value

	waitGroup.Add(2)
	go walkBranch(branch.Left, valuesChannel, waitGroup)
	go walkBranch(branch.Right, valuesChannel, waitGroup)
}

// Walk walks the tree t sending all values
// from the tree to the channel.
func Walk(root *tree.Tree, valuesChannel chan<- int) {
	defer close(valuesChannel)

	if nil == root {
		return
	}

	valuesChannel <- root.Value

	var waitGroup *sync.WaitGroup = &sync.WaitGroup{}
	waitGroup.Add(2)

	go walkBranch(root.Left, valuesChannel, waitGroup)
	go walkBranch(root.Right, valuesChannel, waitGroup)

	waitGroup.Wait()
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	occurrenceMap := make(map[int]int)

	var t1ValuesChannel, t2ValuesChannel = make(chan int), make(chan int)
	var done = make(chan bool)

	go func() {
		go Walk(t1, t1ValuesChannel)

		for value := range t1ValuesChannel {
			concurrentMapWrite(occurrenceMap, value, 1)
		}

		done <- true
	}()

	go func() {
		go Walk(t2, t2ValuesChannel)

		for value := range t2ValuesChannel {
			concurrentMapWrite(occurrenceMap, value, -1)
		}

		done <- true
	}()

	<-done
	<-done

	fmt.Println(occurrenceMap)

	return 0 == len(occurrenceMap)
}

func TestWalk() {
	var valuesChannel chan int = make(chan int)

	go Walk(tree.New(2), valuesChannel)

	defer fmt.Println()

	for value := range valuesChannel {
		fmt.Print(value, " ")
	}
}

func TestSame() {
	fmt.Println(Same(tree.New(1), tree.New(1))) // should return true
	fmt.Println(Same(tree.New(1), tree.New(2))) // should return false.
}

func equivalentBinaryTrees() {
	TestWalk()
	TestSame()
}
