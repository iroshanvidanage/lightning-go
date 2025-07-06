package defaultargs

import "reflect"

type Parameters struct {
	WriteToFile     bool   `default:false`
	DurationSeconds int    `default:"default-durationSeconds"`
	RoleArn         string `default:"default-roleArn"`
}

func main() {
	defaultargs()
}

func defaultargs(prm Parameters) {

	if prm.WriteToFile == 
}
