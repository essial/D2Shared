package d2datadict

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/D2Shared/d2common/d2resource"

	"github.com/OpenDiablo2/D2Shared/d2common/d2interface"

	dh "github.com/OpenDiablo2/D2Shared/d2helper"
)

type LevelTypeRecord struct {
	Name      string
	Id        int
	Files     [32]string
	Beta      bool
	Act       int
	Expansion bool
}

var LevelTypes []LevelTypeRecord

func LoadLevelTypes(fileProvider d2interface.FileProvider) {
	data := strings.Split(string(fileProvider.LoadFile(d2resource.LevelType)), "\r\n")[1:]
	LevelTypes = make([]LevelTypeRecord, len(data))
	for i, j := 0, 0; i < len(data); i, j = i+1, j+1 {
		idx := -1
		inc := func() int {
			idx++
			return idx
		}
		if len(data[i]) == 0 {
			continue
		}
		parts := strings.Split(data[i], "\t")
		if parts[0] == "Expansion" {
			j--
			continue
		}
		LevelTypes[j].Name = parts[inc()]
		LevelTypes[j].Id = dh.StringToInt(parts[inc()])
		for fileIdx := range LevelTypes[i].Files {
			LevelTypes[j].Files[fileIdx] = parts[inc()]
			if LevelTypes[j].Files[fileIdx] == "0" {
				LevelTypes[j].Files[fileIdx] = ""
			}

		}
		LevelTypes[j].Beta = parts[inc()] != "1"
		LevelTypes[j].Act = dh.StringToInt(parts[inc()])
		LevelTypes[j].Expansion = parts[inc()] != "1"
	}
	log.Printf("Loaded %d LevelType records", len(LevelTypes))
}
