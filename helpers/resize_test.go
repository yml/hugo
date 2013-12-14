package helpers

import (
	"log"
	"testing"
)

func TestThumbnailUrl(t *testing.T) {
	log.Println("Running my TestThumbnailUrl")
	// TODO (yml) be smarter. Provide a test image or create one on the fly
	ThumbnailUrl("/home/yml/Dropbox/Devs/golang/workspace_hugo/baracuda.jpg", "200", "100")
}
