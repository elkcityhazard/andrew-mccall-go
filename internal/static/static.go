package static

import "embed"

//go:embed css
//go:embed dist
var static embed.FS

func GetStaticDir() embed.FS {
	return static
}
