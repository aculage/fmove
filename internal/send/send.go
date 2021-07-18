package send

import (
	"encoding/binary"
	"log"
	"net"
	"os"
)

func uint64tobyte(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	return b
}

func SendFile(conn net.Conn, path string) {
	defer conn.Close()

	buffer := make([]byte, 4096)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err.Error())
	}
	conn.Write(uint64tobyte(uint64(len(fileInfo.Name()))))
	conn.Write([]byte(fileInfo.Name()))               //ack needed
	conn.Write(uint64tobyte(uint64(fileInfo.Size()))) //ack needed
	for {
		r, err := file.Read(buffer)
		if r == 0 {
			break
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("Sending %vB to %v\n", r, conn.RemoteAddr())
		conn.Write(uint64tobyte(uint64(r)))
		conn.Write(buffer)
	}
	log.Println("File sent")
}
