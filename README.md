## sensupluginsfile

## Commands
 * checkFileHandles

## Usage

### checkFileHandles
Check the number of open file handles for a process.

Ex. ` sensupluginsfile checkFileHandles --app foo --warn 20 --crit 30`

## Installation

1. godep go build -o bin/sensupluginsfile
1. chmod +x sensupluginsfile (*nix only)
1. cp sensupluginsfile /usr/local/bin

## Notes
