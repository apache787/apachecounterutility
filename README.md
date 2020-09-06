# Apache's Counter Utility
 This utility provides the ability to configure text files with hotkeys to increment, decrement, and reset the counters.  Intended use is for OBS or other streaming utilities that allow reading from a text file.

# Disclaimer
 This program was written fast, probably could use some code organization, and not really optimized at all.

# Installation
 1) Download the utility to a folder that is accessible and easily found.  Recommended to place it in its own folder.
 2) Run the utility once in order to generate the default config file and text counter files
 3) Edit the counters.cfg file in your favorite text editor.

 # Example Configuration
 This is an example configuration:
```
{
  "Counters": [
    {
      "Prefix": "Deaths: ",
      "Path": "counters/deaths.txt",
      "Hotkeys": {
        "Increase": "Ctrl+Alt+F1",
        "Decrease": "Ctrl+Alt+F2",
        "Reset": "Ctrl+Alt+F10"
      }
    },
    {
      "Prefix": "Apache's Example Counter: ",
      "Path": "C:/ApacheCounterUtility/example.txt",
      "Hotkeys": {
        "Increase": "Ctrl+Alt+F5",
        "Decrease": "Ctrl+Alt+F6",
        "Reset": "Ctrl+Alt+F11"
      }
    }
  ]
}
```
This has two counters, one who's path is relative to the executable, and the other has an explicit path on the C Drive as an example.

### Available Key Modifiers
The current version does not distinguish between Left or Right modifiers.  Write Modifiers as the following:
```
Ctrl+
Alt+
Shift+
Win+
```

### Using Keys
Currently functions using Uppercase chacters, Numbers, and F1 though F22 keys.

Todo:
Fix to enable numpad and sytax keys such as punctuations and brackets

### Quirks
Hotkeys set seem to take priority over anything else.  For example: Setting a hotkey to `Ctrl+C` or `Ctrl+V` will prevent you from copy/pastng while the program is active.  Attempt to use unique hotkeys.