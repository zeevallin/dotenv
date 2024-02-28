package loader

import "github.com/zeevallin/dotenv"

func init() {
	f, err := dotenv.Open(".env")
	if err != nil {
		panic(err)
	}
	for _, line := range f.Lines {
		if err := dotenv.Setenv(line.Key, line.Value); err != nil {
			panic(err)
		}
	}
}
