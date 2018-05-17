package formatter

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	defaultTimestampFormat = time.RFC3339Nano
)

type logSeverity = int16

const (
	severityDefault   logSeverity = 0
	severityDebug                 = 100
	severityInfo                  = 200
	severityNotice                = 300
	severityWarning               = 400
	severityError                 = 500
	severityCRITICAL              = 600
	severityAlert                 = 700
	severityEmergency             = 800
)

var levelMap = map[logrus.Level]logSeverity{
	logrus.DebugLevel: severityDebug,
	logrus.InfoLevel:  severityInfo,
	logrus.WarnLevel:  severityWarning,
	logrus.ErrorLevel: severityError,
	logrus.PanicLevel: severityEmergency,
}

// GKELogFormatter formats messages for use to output to stdout on a container on GKE.
type GKELogFormatter struct {
	TimestampFormat string
}

// Format formats log so Stackdriver Logging can construe.
func (f *GKELogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	prefixFieldClashes(data)

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}
	data["time"] = entry.Time.Format(timestampFormat)
	data["message"] = entry.Message
	data["severity"] = convertLogLevel(entry.Level)

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields: %v", err)
	}
	return append(serialized, '\n'), nil
}

func convertLogLevel(level logrus.Level) logSeverity {
	severity, found := levelMap[level]
	if !found {
		return severityDefault
	}
	return severity
}

func prefixFieldClashes(data logrus.Fields) {
	if t, ok := data["time"]; ok {
		data["fields.time"] = t
	}

	if m, ok := data["msg"]; ok {
		data["fields.msg"] = m
	}

	if l, ok := data["level"]; ok {
		data["fields.level"] = l
	}
}
