# logrus-formatter-gke
A Formatter implementation for logging on GKE.

It aims to support both the specs of [logrus](https://github.com/sirupsen/logrus) and GKE's [structured logging](https://cloud.google.com/logging/docs/structured-logging).

For now, it supports logrus' custom fields and caller.

When caller is set to `true`, `logging.googleapis.com/sourceLocation` will be used for caller in logging.

This is a fork of yanana's [logrus-formatter-gke](https://github.com/yanana/logrus-formatter-gke). Thanks for yanana's greak work. 🎉
