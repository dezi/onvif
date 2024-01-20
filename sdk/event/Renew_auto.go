// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package event

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/event"
	"github.com/use-go/onvif/sdk"
)

// Call_Subscribe forwards the call to dev.CallMethod() then parses the payload of the reply as a SubscribeResponse.
func Call_Renew(ctx context.Context, dev *onvif.Device, request event.Renew) (event.RenewResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RenewResponse event.RenewResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RenewResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "Renew")
		return reply.Body.RenewResponse, errors.Annotate(err, "reply")
	}
}
