package main

import (
	"os"
	"path/filepath"

	"github.com/bbfh-dev/mend/internal"
	"github.com/bbfh-dev/mend/mend"
	cp "github.com/otiai10/copy"
)

const INPUT_DIR = 1
const OUTPUT_DIR = 2

func main() {
	inputPath, outputPath := getArgs()

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		if err := os.MkdirAll(outputPath, os.ModePerm); err != nil {
			internal.Error(err, "Creating output dir")
		}
	}

	if err := internal.RemoveContents(outputPath); err != nil {
		internal.Error(err, "Clearing output dir")
	}

	for _, entry := range internal.ReadDir(inputPath) {
		processEntry(inputPath, outputPath, entry)
	}

	for path, body := range mend.TEMPLATES {
		internal.WriteFile(filepath.Join(outputPath, path), mend.ParseTemplate(body).Output())
	}
}

func getArgs() (string, string) {
	if len(os.Args) < 2 {
		internal.Error(nil, "No dir is specified")
	}

	if len(os.Args) < 3 {
		internal.Error(nil, "No output dir is specified")
	}

	inputPath, err := filepath.Abs(os.Args[INPUT_DIR])
	if err != nil {
		internal.Error(err, "Resolving input path")
	}

	outputPath, err := filepath.Abs(os.Args[OUTPUT_DIR])
	if err != nil {
		internal.Error(err, "Resolving output path")
	}

	return inputPath, outputPath
}

func processEntry(inputDir string, outputDir string, entry os.DirEntry) {
	switch entry.Name() {
	case "assets":
		err := cp.Copy(
			filepath.Join(inputDir, entry.Name()),
			filepath.Join(outputDir, entry.Name()),
		)
		if err != nil {
			internal.Error(err, "Copying assets")
		}
	case "icons":
		for _, entry := range internal.ReadDir(inputDir, "icons") {
			mend.ICONS[internal.Base(entry)] = internal.ReadFile(inputDir, "icons", entry.Name())
		}
	case "base.html":
		mend.BASE_HTML = internal.ReadFile(inputDir, entry.Name())
	case "components":
		for _, entry := range internal.ReadDir(inputDir, "components") {
			reader := internal.FileReader(inputDir, "components", entry.Name())
			components := mend.ParseComponent(reader)
			mend.COMPONENTS[internal.Base(entry)] = components.Output()
		}
	default:
		if entry.IsDir() {
			for _, subentry := range internal.ReadDir(inputDir, entry.Name()) {
				mend.TEMPLATES[filepath.Join(entry.Name(), subentry.Name())] = internal.ReadFile(
					inputDir,
					entry.Name(),
					subentry.Name(),
				)
			}
		} else {
			mend.TEMPLATES[entry.Name()] = internal.ReadFile(inputDir, entry.Name())
		}
	}
}
