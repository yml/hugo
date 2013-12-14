// Copyright © 2013 Steve Francia <spf@spf13.com>.
//
// Licensed under the Simple Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://opensource.org/licenses/Simple-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helpers

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	_ "image"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path"
	"strconv"
)

const (
	thumbnailDir = "thumbs"
	thumbnailExt = "jpeg"
)

func CreateThumbnailDir(thumbDir string) error {
	if _, err := os.Stat(thumbDir); err != nil {
		if os.IsNotExist(err) {
			if os.MkdirAll(thumbDir, 0700) != nil {
				log.Println("[Error] An error occured while creating the thumbDir: ", err)
			}
		}
	}
	return nil
}

func thumbnailPath(file, width, height string) string {
	dir, base := path.Split(file)
	thumbnailName := fmt.Sprintf("%s_%sx%s.%s", base, width, height, thumbnailExt)
	return path.Join(dir, thumbnailDir, thumbnailName)
}

func CreateThumbnail(filePath, thumbPath, width, height string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("[Error] error while opening the Image: ", err)
		return err
	}
	defer file.Close()

	// decode jpeg into image.Image
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("[Error] An error while decoding the Image: ", err)
		return err
	}

	w, err := strconv.Atoi(width)
	if err != nil {
		log.Println("[Error] An error occured while converting the width to int: ", err)
		return err
	}
	h, err := strconv.Atoi(height)
	if err != nil {
		log.Println("[Error] An error occured while converting the height to int: ", err)
		return err
	}

	if err := CreateThumbnailDir(path.Dir(thumbPath)); err != nil {
		log.Println("[Error] An error occured while creating thumbPath: ", err)
	}

	// resize to width and height using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	out, err := os.Create(thumbPath)
	if err != nil {
		log.Println("[Error] An error occured while saving the thumb", err)
		return err
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
	return nil
}

func ThumbnailUrl(file, width, height string) string {
	log.Println("[Debug] running ThumbnailUrl with args: ", file, width, height)
	thumbPath := thumbnailPath(file, width, height)
	err := CreateThumbnail(file, thumbPath, width, height)
	if err != nil {
		log.Println("[Debug] An error occured while trying to Create the thumbnail")
	}
	return thumbPath
}
