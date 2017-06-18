package subtitles

import (
	log "github.com/Sirupsen/logrus"

	"github.com/kennygrant/sanitize"
)

// filterHTML removes all html tags from captions
func (subtitle *Subtitle) filterHTML() *Subtitle {
	for _, cap := range subtitle.Captions {
		for i, line := range cap.Text {
			clean := sanitize.HTML(line)
			if clean != cap.Text[i] {
				log.Printf("[html] %s -> %s\n", cap.Text[i], clean)
				cap.Text[i] = clean // XXX works?!
			}
		}
	}
	return subtitle
}