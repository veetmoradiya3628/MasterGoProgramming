package main

import "fmt"

type ConfigItem struct {
	Key   string
	Value any
	IsSet bool
}

/*
	%v - the default formatting
	%+v - with field name
	%#v - value
	%T - type
	%s - string
	%d - int
	%.nf - float with n precision
	%t - boolean
	%q - double quote string
	%% - escape character
*/
func (c ConfigItem) String() string {
	return fmt.Sprintf("Key : %s, Value: %s, IsSet : %t", c.Key, c.Value, c.IsSet)
}

func main() {
	appName := "EnvParser"
	version := 1.2
	port := 8080
	isEnabled := true

	status := fmt.Sprintf("Application: %s (Version: %.1f) running on port %d. Enabled: %t", appName, version, port, isEnabled)
	fmt.Println(status)

	item1 := ConfigItem{"appName", appName, true}
	fmt.Println(item1)
	item2 := ConfigItem{"version", version, true}
	fmt.Println(item2)
	item3 := ConfigItem{"port", port, true}
	fmt.Println(item3)
}
