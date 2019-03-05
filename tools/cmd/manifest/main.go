package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Graphic struct {
	Filename    string   `json:"file_name"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type Manifest []*Graphic

func (m Manifest) FindByFilename(name string) *Graphic {
	for _, g := range m {
		if g.Filename == name {
			return g
		}
	}
	return nil
}

func main() {

	svgDir := flag.String("svg-path", "./svg", "Path to directory containing SVGs")
	manifestPath := flag.String("manifest-path", "./gallery/manifest.json", "Location of manifest file")
	flag.Parse()

	_, err := updateManifest(*svgDir, *manifestPath)
	if err != nil {
		log.Fatalf("Failed to update manifest: %s", err.Error())
	}
	fmt.Println("Updated Manifest!")
}

func updateManifest(svgDir string, manifestPath string) (Manifest, error) {

	manifest := Manifest{}

	// parse existing manifest
	manifestFile, err := os.OpenFile(manifestPath, os.O_RDWR|os.O_CREATE, 0655)
	if err == nil {
		defer manifestFile.Close()
		stat, err := manifestFile.Stat()
		if err != nil {
			return nil, fmt.Errorf("failed to stat manifest: %s", err)
		}
		if stat.Size() >0{
			if err := json.NewDecoder(manifestFile).Decode(&manifest); err != nil {
				return nil, fmt.Errorf("failed to decode manifest: %s", err)
			}
		}
	} else {
		return nil, err
	}

	// merge in any new graphics
	files, err := ioutil.ReadDir(svgDir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".svg") {
			continue
		}
		g := &Graphic{
			Filename: f.Name(),
			Name: strings.ToTitle(strings.Replace(strings.TrimSuffix(f.Name(), ".svg"), "-", " ", -1)),
		}
		if manifest.FindByFilename(g.Filename) == nil {
			log.Printf("ADDED %s", g.Filename)
			manifest = append(manifest, g)
		}
	}

	// clean up any entries which do not have a graphic any more
	cleanManifest := Manifest{}
	for _, g := range manifest {
		if svgInSlice(files, g.Filename) {
			cleanManifest = append(cleanManifest, g)
		}
	}
	manifest = cleanManifest

	// write the new version
	b, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return nil, err
	}
	if err := manifestFile.Truncate(0); err != nil {
		return nil, err
	}
	if _, err := manifestFile.WriteAt(b, 0); err != nil {
		return nil, err
	}
	return manifest, err
}


func svgInSlice(files []os.FileInfo, findName string) bool {
	for _, f := range files {
		if f.Name() == findName && strings.HasSuffix(f.Name(), ".svg") {
			return true
		}
	}
	return false
}
