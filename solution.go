package solution

import "github.com/kyokomi/emoji/v2"

func GetMessage() string {
	greeting := emoji.Sprint("Hello :world_map:")
	return greeting
}
