package metadata

import (
	"bufio"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/xml"
	"log"
	"os"
	"strings"
)

// >> TODO: duplicated from lingo-extremote/extremote.go
type PlayerState byte

const (
	PlayerStateStopped PlayerState = 0x00
	PlayerStatePlaying PlayerState = 0x01
	PlayerStatePaused  PlayerState = 0x02
	PlayerStateError   PlayerState = 0xff
)

// <<

type Item struct {
	XMLName xml.Name `xml:"item"`
	Type    string   `xml:"type"`
	Code    string   `xml:"code"`
	Length  int      `xml:"length"`
	Data    string   `xml:"data"`
}

type MetadataParser struct {
	IndexStr   string
	TrackIndex int32
	Artist     string
	Album      string
	Title      string
	Length     int
	Status     PlayerState
}

func (mp *MetadataParser) Start() {
	// TODO: make a parameter
	file, err := os.Open("/tmp/shairport-sync-metadata")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	var tag string
	for scanner.Scan() {
		chunk := scanner.Text()
		tag += chunk

		isClosing := strings.HasSuffix(chunk, "</item>")
		if isClosing {
			var item Item
			xml.Unmarshal([]byte(tag), &item)

			var decodedItem = decodeItem(item)

			switch decodedItem.Code {
			case "asar":
				log.Printf("Artist: %s", decodedItem.Data)
				mp.Artist = decodedItem.Data
			case "asal":
				log.Printf("Album Name: %s", decodedItem.Data)
				mp.Album = decodedItem.Data
			case "minm":
				log.Printf("Title: %s", decodedItem.Data)
				mp.Title = decodedItem.Data
			case "astm":
				trackLength := binary.BigEndian.Uint32([]byte(decodedItem.Data))
				log.Printf("Track length: %d", trackLength)
				mp.Length = int(trackLength)
			case "pffr":
				log.Printf(">> Play")
				mp.Status = PlayerStatePlaying
			case "paus":
			case "pend":
				log.Printf(">> Pause")
				mp.Status = PlayerStatePaused
			case "pres":
				log.Printf(">> Resume")
				mp.Status = PlayerStatePlaying
			case "mdst":
				log.Printf("Metadata bundle start")
			case "mden":
				log.Printf("Metadata bundle end")
				indexStr := mp.Album + mp.Artist + mp.Title
				if indexStr != mp.IndexStr {
					mp.TrackIndex += 1
					mp.IndexStr = indexStr
				}
			default:
			}

			tag = ""
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (mp MetadataParser) Stop() {
	_, w, _ := os.Pipe()

	w.Close()
}

func decodeItem(item Item) Item {
	iCode, _ := hex.DecodeString(item.Code)
	item.Code = string(iCode)

	iType, _ := hex.DecodeString(item.Type)
	item.Type = string(iType)

	iData, _ := base64.StdEncoding.DecodeString(item.Data)
	item.Data = string(iData)

	return item
}
