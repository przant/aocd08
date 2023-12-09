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
    pf, err := os.Open("input.txt")
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

    steps := make([]uint64, 0)

    for _, fNode := range finalNodes {
        for _, sNode := range startNodes {
            if network[sNode][Left] == network[fNode][Left] || network[sNode][Right] == network[fNode][Left] {
                steps = append(steps, walk(instr, sNode, fNode, network))
            }
        }
    }

    gcd := steps[0]
    for n := 1; n < len(steps); n++ {
        gcd = GCD(gcd, steps[n])

        if gcd == 1 {
            break
        }
    }

    result := steps[0]
    for i := 0; i < len(steps)-1; i++ {
        result *= (steps[i+1] / gcd)
    }

    fmt.Println()
    fmt.Println(result)

}

func walk(instr, sN, fN string, net Network) uint64 {
    steps := uint64(0)
    fmt.Println(sN)
    for {
        for _, step := range instr {
            steps++
            sN = net[sN][step]
            if sN == fN {
                return steps
            }
        }
    }
}

func GCD(a, b uint64) uint64 {
    for a > 0 && b > 0 {
        if a > b {
            a = a % b
        } else {
            b = b % a
        }
    }

    if a == 0 {
        return b
    }
    return a
}
