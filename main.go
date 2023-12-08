package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strings"
)

const (
    Left      = 'L'
    Right     = 'R'
    startNode = "AAA"
    finalNode = "ZZZ"
)

var (
    startNodes []string
    finalNodes []string
)

type Network map[string]map[rune]string

func main() {
    pf, err := os.Open("example3.txt")
    if err != nil {
        log.Fatalf("while opening file %q: %s", pf.Name(), err)
    }
    defer pf.Close()

    scnr := bufio.NewScanner(pf)

    var instr string
    network := make(Network)
    reSteps := regexp.MustCompile(`^\s*([LR]+)\s*$`)
    reNode := regexp.MustCompile(`([A-Z1-9]{3})\s=\s\(([A-Z1-9]{3}),\s([A-Z1-9]{3})\)`)

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
            if strings.HasSuffix(nodes[1], "A") {
                startNodes = append(startNodes, nodes[1])
            }
            if strings.HasSuffix(nodes[1], "Z") {
                finalNodes = append(finalNodes, nodes[1])
            }
        }
    }

    steps := 0
    // nextNode := startNode
    nextNodes := make([]string, 0)

    nextNodes = append(nextNodes, startNodes...)

    for {
        for _, step := range instr {
            finish := true
            fmt.Println(nextNodes)
            steps++
            for idx, nextNode := range nextNodes {
                nextNodes[idx] = network[nextNode][step]
            }

            for _, nextNode := range nextNodes {
                if !strings.HasSuffix(nextNode, "Z") {
                    finish = false
                    break
                }
            }

            if finish {
                fmt.Println(nextNodes)
                fmt.Println(steps)
                return
            }
        }
    }
}
