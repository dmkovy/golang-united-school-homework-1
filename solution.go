package solution

import "github.com/kyokomi/emoji/v2"

func GetMessage() string {
	hello := emoji.Sprint("Hello :world_map:")
	return hello
}
