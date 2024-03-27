## Config Loader

This Golang package helps to load and get values from configuration files, most specifically **json** and **yml** types.
For **.env**, **.properties** or **toml** extensions, please checkout for another great packages to solve your context:

* https://github.com/joho/godotenv
* https://github.com/spf13/viper


### Technologies

* Golang 1.21

### Config Loader
Config Loader are the main structure in the package. It will provide the **Parse()** function to read values from config file, returning the map[string]interface{}.

```go
type FileProperties struct {
	Name_          string   `json:"name"`
	Type_          string   `json:"type"`
	PathLocations_ []string `json:"path_locations"`
	FinalPath      string   `json:"final_path"`
}
``` 
| Property          | Description                              |
|-------------------|------------------------------------------|
|Name               | Name of config file                      |
|Type               | Specifies the type of file (yaml or json)|
|PathLocations      | List of paths according to the context (local environment, container, CI/CD etc.)|
|FinalPath          | After parse, this properties will have the definitive config file path location|

### Usage

**example.yaml**
```yaml
AppName: "my-golang-app"
Version: "1.0.0"
LogLevel: 0
```

Basic implementation
```go
func main() {
	cl := configLoader.New().
		Name("app").
		FileType(filehandlers.YML).
		PathLocations(".", "./configs"). // assuming that file location is . or ./configs
		Build()

	configs, err := cl.Parse()
	if err != nil {
		panic(err)
	}
	fmt.Println(configs["AppName"])
	fmt.Println(fp.FinalPath)
}
```

The output:
```bash
$ go run main.go 
my-app
./configs
```

Thank you! Enjoy!
