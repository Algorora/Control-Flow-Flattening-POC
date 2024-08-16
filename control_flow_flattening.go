package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func randomizeVariable(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	name := make([]byte, 8)
	for i := range name {
		name[i] = chars[rand.Intn(len(chars))]
	}
	return prefix + string(name)
}

func flattenControlFlow(script string) string {
	var sb strings.Builder
	controlVariables := []string{randomizeVariable("var1_"), randomizeVariable("var2_"), randomizeVariable("var3_")}
	sb.WriteString(fmt.Sprintf("$%s = 0\n", controlVariables[0]))

	parts := strings.Split(script, "\n")
	for i, part := range parts {
		if i%2 == 0 {
			sb.WriteString(fmt.Sprintf("$%s = %d\n", controlVariables[1], i))
		} else {
			sb.WriteString(fmt.Sprintf("if ($%s -eq %d) {\n%s\n} else {\n", controlVariables[1], i, part))
		}
	}
	sb.WriteString(fmt.Sprintf("}\n$%s = 1\n", controlVariables[2]))
	return sb.String()
}

func main() {
	scriptPath := flag.String("script", "", "Path to the PowerShell script (.ps1) file")
	flag.Parse()

	if *scriptPath == "" {
		fmt.Println("Please provide a path to the PowerShell script using the -script flag.")
		os.Exit(1)
	}

	scriptContent, err := ioutil.ReadFile(*scriptPath)
	if err != nil {
		fmt.Printf("Error reading script file: %v\n", err)
		os.Exit(1)
	}

	obfuscatedScript := flattenControlFlow(string(scriptContent))
	fmt.Println("Control Flow Flattened PowerShell Script:")
	fmt.Println(obfuscatedScript)
}
