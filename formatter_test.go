package formatter

import (
	"bytes"
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestConvertLogLevel(t *testing.T) {
	log.SetFormatter(&GKELogFormatter{
		TimestampFormat: "XYZ",
	})
	a := assert.New(t)
	buf := new(bytes.Buffer)
	// log.SetOutput(buf)
	log.SetOutput(buf)
	log.SetLevel(log.DebugLevel)
	log.Debug("HOGE")

	out := buf.String()
	a.JSONEq(`{
		"severity": 100,
		"message": "HOGE",
		"time": "XYZ"
	}`, out)
	fmt.Println(out)
}
