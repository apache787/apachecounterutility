package models

//Counter defines the path and assigned Hotkeys
type Counter struct {
	Prefix  string
	Path    string
	Count   int `json:"-"`
	Hotkeys Hotkeys
}

//Counters defines an array of Coutner
type Counters struct {
	Quit     string
	Counters []Counter
}
