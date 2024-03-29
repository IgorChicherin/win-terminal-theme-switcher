# Windows Terminal Theme Switcher 

![Windows Terminal](https://github.com/IgorChicherin/win-terminal-theme-switcher/blob/master/images/terminal.gif)
![Windows Auto Dark Mode](https://github.com/IgorChicherin/win-terminal-theme-switcher/blob/master/images/auto_dark.gif)

# Plugin for:
> [Windows Auto Dark Mode](https://github.com/AutoDarkMode/Windows-Auto-Night-Mode)
> and
> [Windows Terminal](https://www.microsoft.com/store/productId/9N0DX20HK701?ocid=pdpshare)

# Usage of win-terminal-theme-switcher.exe:
```
 -s, --settings string   Terminal settings path (default C:\\Users\\%USERNAME%\\Local\\Packages\\Microsoft.WindowsTerminal_8wekyb3d8bbwe\\LocalState\\settings.json")
 -t, --theme string      Theme name (default "One Half Dark")

``` 

# [Configuration](https://github.com/AutoDarkMode/Windows-Auto-Night-Mode/wiki/How-to-add-custom-scripts) Auto dark mode 

# Example Config for Windows Auto Dark Mode:
# scripts.yaml
```
Enabled: true
Component:
  Scripts:
  - Name: cmd theme switch
    Command: C:\Users\r00t\win-terminal-theme-switcher.exe
    WorkingDirectory: C:\Users\r00t\AppData\Roaming\AutoDarkMode
    ArgsLight: [-t, "One Half Light"]
    ArgsDark: [-t, "One Half Dark"]
    AllowedSources: [Any]


    
```
