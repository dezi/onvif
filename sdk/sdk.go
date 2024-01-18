package sdk

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/juju/errors"
	"github.com/liesa-care/project.go.liesa.main/goodies/smart/xmlpretty"
	"github.com/rs/zerolog"
)

var debug = false

var (
	// LoggerContext is the builder of a zerolog.Logger that is exposed to the application so that
	// options at the CLI might alter the formatting and the output of the logs.
	LoggerContext = zerolog.
			New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			With().Timestamp()

	// Logger is a zerolog logger, that can be safely used from any part of the application.
	// It gathers the format and the output.
	Logger = LoggerContext.Logger()
)

func SetDebug(dgb bool) {
	debug = dgb
}

func ReadAndParse(ctx context.Context, httpReply *http.Response, reply interface{}, tag string) error {

	if debug {
		Logger.Debug().
			Str("msg", httpReply.Status).
			Int("status", httpReply.StatusCode).
			Str("action", tag).
			Msg("RPC")
	}

	if httpReply.StatusCode == 400 {
		debug = false
		return errors.New("400 bad request")
	}

	// TODO(jfsmig): extract the deadline from ctx.Deadline() and apply it on the reply reading
	b, err := ioutil.ReadAll(httpReply.Body)
	if err != nil {
		return errors.Annotate(err, "read")
	}

	httpReply.Body.Close()

	if debug {
		pretty := xmlpretty.FormatXMLDezi(string(b))
		fmt.Printf(">>>>>>>>>>>>>>>>\n%s\n----------------\n", pretty)
		debug = false
	}

	err = xml.Unmarshal(b, reply)
	return errors.Annotate(err, "decode")
}
