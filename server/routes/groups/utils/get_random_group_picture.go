package utils

import (
	"fmt"
	"math/rand"
)

// WHEN YOU CREATE A GROUP, A RANDOMLY CHOSEN PICTURE WILL BE ADDED AS THUMBNAIL & WALLPAPER

func GetRandomDefaulGroupPicture() string {
	// WE GOT THREE DIFFERENT POSSIBILITIES
	return fmt.Sprintf("/images/group/default%d.jpg", rand.Intn(4))
}
