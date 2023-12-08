package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
)

const (
    Left      = 'L'
    Right     = 'R'
    startNode = "AAA"
    finalNode = "ZZZ"
)

type Network map[string]map[rune]string

func main() {
    pf, err := os.Open("example2.txt")
    if err != nil {
        log.Fatalf("while opening file %q: %s", pf.Name(), err)
    }
    defer pf.Close()

    scnr := bufio.NewScanner(pf)

    var instr string
    network := make(Network)
    reSteps := regexp.MustCompile(`\s*([LR]+)\s*`)
    reNode := regexp.MustCompile(`([A-Z]{3})\s=\s\(([A-Z]{3}),\s([A-Z]{3})\)`)

    for scnr.Scan() {
        line := scnr.Text()
        if reSteps.MatchString(line) {
            instr = reSteps.FindStringSubmatch(line)[1]
        }
        if reNode.MatchString(line) {
            nodes := reNode.FindStringSubmatch(line)
            network[nodes[1]] = make(map[rune]string)
            network[nodes[1]][Left] = nodes[2]
            network[nodes[1]][Right] = nodes[3]
        }
    }

    steps := 0
    nextNode := startNode

    for {
        for _, step := range instr {
            fmt.Println(nextNode)
            steps++
            nextNode = network[nextNode][step]
            if nextNode == finalNode {
                fmt.Println(nextNode)
                fmt.Println(steps)
                return
            }
        }
    }
}
