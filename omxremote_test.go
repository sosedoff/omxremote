package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fileToTitle(t *testing.T) {
	examples := []string{
		"Movie Name (2014) [1080p].mp4",
		"Movie Name (2009) [1080p] [HSBS] [3d].mp4",
		"Movie.Name.2011.480p.BRRip.XviD.AC3-AsA.mp4",
		"Movie Name 2007 BRRip 720p x264 AAC - PRiSTiNE [P2PDL].mp4",
		"Movie Name.2011.limited.720p.BRRip.H264.AAC-MAJESTiC.mp4",
		"Movie.Name.2010.1080p.BrRip.x264.YIFY.mp4",
		"Movie.Name.S05E01.HDTV.x264-LOL.mp4",
	}

	for _, val := range examples {
		assert.Equal(t, fileToTitle(val), "Movie Name")
	}
}
