package get

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"os"
)

func GetFile(conn net.Conn, path string) {
	b := make([]byte, 8)

	_, err := conn.Read(b)
	if err != nil {
		log.Fatal(err.Error())
	}
	fileName := make([]byte, binary.LittleEndian.Uint64(b))
	_, err = conn.Read(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	file, err := os.Create(path + string(fileName))
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = conn.Read(b)
	if err != nil {
		log.Fatal(err.Error())
	}
	fileSize := binary.LittleEndian.Uint64(b)

	buffer := make([]byte, 4096)
	for {
		_, err := conn.Read(b)
		if err != nil {
			log.Println(err.Error())
		}
		bufsize := binary.LittleEndian.Uint64(b)
		i, err := conn.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		if i > 0 {
			log.Printf("Read %vB\n", bufsize)
			file.Write(buffer[:bufsize-1])
		}
	}

	if info, _ := os.Stat(string(fileName)); info.Size() == int64(fileSize) {
		log.Println("Transaction complete")
	} else {
		log.Println("Transaction failure")
	}
	conn.Close()
}
