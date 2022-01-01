package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

const (
	Magic     = (('P' << 24) + ('S' << 16) + ('B' << 8) + 'I')
	HeaderLen = 160 // magic + version + lump metadata
	EntLump   = 0   // the location in the header
)

/**
 * Just simple error checking
 */
func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * Read 4 bytes as a Long
 */
func ReadLong(input []byte, start int) int32 {
	var tmp struct {
		Value int32
	}

	r := bytes.NewReader(input[start:])
	if err := binary.Read(r, binary.LittleEndian, &tmp); err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	return tmp.Value
}

/**
 * Make sure the first 4 bytes match the magic number
 */
func VerifyHeader(header []byte) {
	if ReadLong(header, 0) != Magic {
		panic("Invalid BPS file")
	}
}

/**
 * Find the entity lump in the BSP.
 * Return the location and length
 */
func LocateEntityLump(header []byte) (int, int) {
	var offsets [19]int
	var lengths [19]int

	pos := 8
	for i := 0; i < 18; i++ {
		offsets[i] = int(ReadLong(header, pos)) - HeaderLen
		pos = pos + 4
		lengths[i] = int(ReadLong(header, pos))
		pos = pos + 4
	}

	return offsets[EntLump] + HeaderLen, lengths[EntLump]
}

/**
 * Get a slice of the just the texture lump from the map file
 */
func GetEntityLump(f *os.File, offset int, length int) []byte {
	_, err := f.Seek(int64(offset), 0)
	check(err)

	lump := make([]byte, length)
	read, err := f.Read(lump)
	check(err)

	if read != length {
		panic("reading entity lump: hit EOF")
	}

	return lump
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <q2mapfile.bsp>\n", os.Args[0])
		return
	}

	bspname := os.Args[1]
	bsp, err := os.Open(bspname)
	check(err)

	header := make([]byte, HeaderLen)
	_, err = bsp.Read(header)
	check(err)

	VerifyHeader(header)

	offset, length := LocateEntityLump(header)
	ents := GetEntityLump(bsp, offset, length)
	fmt.Println(string(ents))
}
