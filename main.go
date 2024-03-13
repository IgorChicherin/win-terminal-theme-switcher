//go:generate go build -o=switcher.exe
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

func ParseArgs(programName string, args []string) (string, string) {
	flags := flag.NewFlagSet(programName, flag.ContinueOnError)

	var settingsPath, themeName string

	flags.StringVarP(
		&settingsPath,
		"settings",
		"s",
		"C:\\Users\\r00t\\AppData\\Local\\Packages\\Microsoft.WindowsTerminal_8wekyb3d8bbwe\\LocalState\\settings.json",
		"Terminal settings path")

	flags.StringVarP(&themeName, "theme", "t", "One Half Dark", "Theme name")

	if err := flags.Parse(args); err != nil {
		log.Fatalln(err)
	}

	return settingsPath, themeName
}

func SetConfig(path string, theme string) error {
	jsonFile, err := os.Open(path)

	if err != nil {
		log.Errorln(err)
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var configs map[string]interface{}
	if err := json.Unmarshal(byteValue, &configs); err != nil {
		return err
	}
	profiles := configs["profiles"].(map[string]interface{})
	defaults := profiles["defaults"].(map[string]interface{})

	fmt.Printf("Switching theme `%s`", theme)

	defaults["colorScheme"] = theme

	bytes, err := json.MarshalIndent(configs, "", " ")
	if err != nil {
		log.Fatalf("Error marshaling config: %s", err)
		return err
	}

	err = os.WriteFile(path, bytes, 0644)
	if err != nil {
		log.Fatalf("Error writing file: %s", err)
		return err
	}

	return nil
}

func main() {
	settingsPath, themeName := ParseArgs(os.Args[0], os.Args[1:])

	if err := SetConfig(settingsPath, themeName); err != nil {
		log.Fatalln(err)
	}

}
