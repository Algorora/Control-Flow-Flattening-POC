# Control Flow Flattening Obfuscator for PowerShell Scripts
This repository contains a Go-based tool designed to apply control flow flattening to PowerShell scripts. Control flow flattening is an advanced obfuscation technique that makes the execution path of scripts more complex and harder to analyze. This tool is intended for educational purposes and to demonstrate how control flow flattening can be implemented in practice.

## Features
Control Flow Flattening: Transforms PowerShell scripts to obscure their execution flow by introducing complex conditional branches and random variable names.
Simple and Effective: Provides a straightforward implementation to showcase the control flow flattening technique without additional obfuscation methods.
Usage
Build the Executable:
```sh
go build control_flow_flattening.go
```
Run the Tool:
```sh
./control_flow_flattening -script /path/to/your/script.ps1
```
Replace /path/to/your/script.ps1 with the path to your PowerShell script.

## Example
To obfuscate a PowerShell script, place your script in the designated path and run the tool. The obfuscated script with flattened control flow will be output to the terminal.

## Disclaimer
This tool is provided for educational and research purposes only. Use it responsibly and ensure compliance with all applicable laws and regulations. The creator of this tool does not endorse or condone illegal activities.
