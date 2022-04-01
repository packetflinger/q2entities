# Q2Entities
A simple command-line utility to extract the entities string from a Quake 2 map file. 
## Entities?
Binary Space Partitioning maps (.bsp) contain a blog of text containing the metadata for each game object (entity). Their class, location, direction, and various other attributes are included.
## Compiling
To build for your current system:

`# go build q2entities.go` 

To build for another OS or ARCH (ex: 32bit windows):

`# GOOS=windows GOARCH=386 CGO_ENABLED=0 go build q2entities.go`

## Usage
`# q2ents [-c] <map.bsp>`

The `-c` flag will force a sorted count of each entity. If this flag is missing, the entire entity block will be printed to STDOUT.
