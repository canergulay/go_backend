package utils

import (
	"backend/global/constants"
	"fmt"
	"math/rand"
)

// WHEN YOU CREATE A GROUP, A RANDOMLY CHOSEN PICTURE WILL BE ADDED AS THUMBNAIL & WALLPAPER

func GetRandomDefaulGroupPicture() string {
	// WE GOT THREE DIFFERENT POSSIBILITIES
	return fmt.Sprint(constants.DefaultImagePath, rand.Intn(4))
}
