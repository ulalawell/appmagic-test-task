package reverser

import (
	"fmt"
	"sync"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type task struct {
	index int
	text  string
}

type result struct {
	index    int
	threadId int
	reversed string
}

// ProcessTexts concurrently reverses strings and returns ordered output lines.
func ProcessTexts(texts []string, numThreads int) []string {
	taskCh := make(chan task, len(texts))
	resultCh := make(chan result, len(texts))
	var wg sync.WaitGroup

	// Workers
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(threadId int) {
			defer wg.Done()
			for t := range taskCh {
				resultCh <- result{
					index:    t.index,
					threadId: threadId,
					reversed: reverseString(t.text),
				}
			}
		}(i + 1)
	}

	// Send tasks
	for i, text := range texts {
		taskCh <- task{index: i, text: text}
	}
	close(taskCh)

	// Close result channel after workers are done
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Collect and sort results
	results := make([]result, len(texts))
	for r := range resultCh {
		results[r.index] = r
	}

	var lines []string
	for i, r := range results {
		lines = append(lines, fmt.Sprintf("line %d, thread %d: \"%s\"", i+1, r.threadId, r.reversed))
	}
	lines = append(lines, "done")
	return lines
}
