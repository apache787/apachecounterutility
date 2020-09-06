package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/MakeNowJust/hotkey"
	"github.com/apache787/apachecounterutility/models"
)

//ApacheCounterUtil contains a list of counters
type ApacheCounterUtil struct {
	counters *models.Counters
}

//NewUtil returns a new ApacheCounterUtil
func NewUtil() *ApacheCounterUtil {
	util := ApacheCounterUtil{}
	config, err := loadConfig()
	if err != nil {
		return &util
	}
	util.counters = config
	return &util
}

func (acu *ApacheCounterUtil) start() {
	hkey := hotkey.New()

	quit := make(chan bool)

	mask, key, err := models.ParseHotkey(acu.counters.Quit)
	if err != nil {
		fmt.Println("Invalid Hokey assigned to Quit")
		quit <- true
	}
	hkey.Register(mask, key, func() {
		fmt.Println("Quit")
		quit <- true
	})
	fmt.Println(len(acu.counters.Counters))
	registerHotkeys(acu.counters, hkey)

	fmt.Println("Ready For Hotkeys")
	fmt.Println("Push Ctrl+Alt+Q to escape and quit")

	fmt.Println()
	for i := 0; i < len(acu.counters.Counters); i++ {
		fmt.Printf("%s%d\n", acu.counters.Counters[i].Prefix, acu.counters.Counters[i].Count)
	}
	<-quit
}

func registerHotkeys(counters *models.Counters, hkey *hotkey.Manager) {
	for i := 0; i < len(counters.Counters); i++ {
		fmt.Println("registering keyset")
		mask, key, err := models.ParseHotkey(counters.Counters[i].Hotkeys.Increase)
		if err != nil {
			fmt.Printf("Could not read hotkey %s; skipping\n", counters.Counters[i].Hotkeys.Increase)
		} else {
			idx := i
			hkey.Register(mask, key, func() {
				updateCounter(&counters.Counters[idx], 1)
			})
		}
		mask, key, err = models.ParseHotkey(counters.Counters[i].Hotkeys.Decrease)
		if err != nil {
			fmt.Printf("Could not read hotkey %s; skipping\n", counters.Counters[i].Hotkeys.Decrease)
		} else {
			idx := i
			hkey.Register(mask, key, func() {
				updateCounter(&counters.Counters[idx], -1)
			})
		}
		mask, key, err = models.ParseHotkey(counters.Counters[i].Hotkeys.Reset)
		if err != nil {
			fmt.Printf("Could not read hotkey %s; skipping\n", counters.Counters[i].Hotkeys.Reset)
		} else {
			idx := i
			hkey.Register(mask, key, func() {
				resetCounter(&counters.Counters[idx])
			})
		}
	}
}

func loadConfig() (*models.Counters, error) {
	fmt.Println("Loading counters.cfg")
	if !fileExists("counters.cfg") {
		fmt.Println("Config File not found, creating default")
		config := models.Counters{Quit: "Ctrl+Alt+Q"}
		defaultCounter1 := models.Counter{
			Prefix: "Deaths: ",
			Path:   "counters/deaths.txt",
			Hotkeys: models.Hotkeys{
				Increase: "Ctrl+Shift+F1",
				Decrease: "Ctrl+Shift+F2",
				Reset:    "Ctrl+Shift+F9",
			},
		}
		defaultCounter2 := models.Counter{
			Prefix: "Apache's Example Counter: ",
			Count:  787,
			Path:   "counters/example.txt",
			Hotkeys: models.Hotkeys{
				Increase: "Ctrl+Shift+F5",
				Decrease: "Ctrl+Shift+F6",
				Reset:    "Ctrl+Shift+F10",
			},
		}
		initCounter(&defaultCounter1)
		initCounter(&defaultCounter2)
		config.Counters = append(config.Counters, defaultCounter1, defaultCounter2)
		b, err := json.MarshalIndent(config, "", "	")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		err = ioutil.WriteFile("./counters.cfg", b, 0644)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		return &config, nil
	}

	configFile, err := os.Open("counters.cfg")
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened counters.cfg")
	defer configFile.Close()
	byteValue, _ := ioutil.ReadAll(configFile)
	var config models.Counters
	json.Unmarshal(byteValue, &config)
	for i := 0; i < len(config.Counters); i++ {
		initCounter(&config.Counters[i])
	}
	fmt.Println("Successfully Initialized Counters")
	return &config, nil

}

func initCounter(counter *models.Counter) {
	if !fileExists(counter.Path) {
		createDirectory(path.Dir(counter.Path))
		err := ioutil.WriteFile(counter.Path, []byte(counter.Prefix+strconv.Itoa(counter.Count)), 0644)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		content, err := ioutil.ReadFile(counter.Path)
		if err != nil {
			log.Fatal(err)
		}
		readText := string(content)
		if strings.HasPrefix(readText, counter.Prefix) { // Load Value
			i, err := strconv.Atoi(strings.TrimLeft(readText, counter.Prefix))
			if err != nil {
				err := ioutil.WriteFile(counter.Path, []byte(counter.Prefix+"0"), 0644)
				if err != nil {
					fmt.Println(err)
				}
			}
			counter.Count = i
		} else { //Override Counter
			err := ioutil.WriteFile(counter.Path, []byte(counter.Prefix+"0"), 0644)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func updateCounter(counter *models.Counter, i int) {
	counter.Count += i
	err := ioutil.WriteFile(counter.Path, []byte(counter.Prefix+strconv.Itoa(counter.Count)), 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s%d\n", counter.Prefix, counter.Count)
}
func resetCounter(counter *models.Counter) {
	counter.Count = 0
	err := ioutil.WriteFile(counter.Path, []byte(counter.Prefix+"0"), 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s%d\n", counter.Prefix, counter.Count)
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// createDirector checks if a directory exists, and if not, creates
// all directories required
func createDirectory(dirName string) bool {
	src, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirName, 0755)
		if errDir != nil {
			panic(err)
		}
		return true
	}

	if src.Mode().IsRegular() {
		fmt.Println(dirName, "already exist as a file!")
		return false
	}

	return false
}
