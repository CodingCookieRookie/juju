// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package jujuc

import (
	"github.com/juju/errors"
	"github.com/juju/gnuflag"

	jujucmd "github.com/juju/juju/cmd"
	"github.com/juju/juju/core/status"
	"github.com/juju/juju/internal/cmd"
)

// StatusSetCommand implements the status-set command.
type StatusSetCommand struct {
	cmd.CommandBase
	ctx         Context
	status      string
	message     string
	application bool
}

// NewStatusSetCommand makes a jujuc status-set command.
func NewStatusSetCommand(ctx Context) (cmd.Command, error) {
	return &StatusSetCommand{ctx: ctx}, nil
}

func (c *StatusSetCommand) Info() *cmd.Info {
	doc := `
Sets the workload status of the charm. Message is optional.
The "last updated" attribute of the status is set, even if the
status and message are the same as what's already set.

Further details:
status-set changes what is displayed in juju status.
status-set allows charms to describe their current status.
This places the responsibility on the charm to know its status,
and set it accordingly using the status-set hook tool.
Changes made via status-set are applied without waiting for a
hook execution to end and are not rolled back if a hook
execution fails.

The leader unit is responsible for setting the overall status
of the application by using the --application option.

This hook tool takes 2 arguments. The first is the status code
and the second is a message to report to the user.

Valid status codes are:
    maintenance (the unit is not currently providing a service,
	  but expects to be soon, e.g. when first installing)
    blocked (the unit cannot continue without user input)
    waiting (the unit itself is not in error and requires no
	  intervention, but it is not currently in service as it
	  depends on some external factor, e.g. an application to
	  which it is related is not running)
    active (This unit believes it is correctly offering all
	  the services it is primarily installed to provide)

For more extensive explanations of these status codes, please see
the status reference page https://juju.is/docs/juju/status.

The second argument is a user-facing message, which will be displayed
to any users viewing the status, and will also be visible in the status
history. This can contain any useful information.

In the case of a blocked status though the status message should tell
the user explicitly how to unblock the unit insofar as possible, as this
is primary way of indicating any action to be taken (and may be surfaced
by other tools using Juju, e.g. the Juju GUI).

A unit in the active state with should not generally expect anyone to
look at its status message, and often it is better not to set one at
all. In the event of a degradation of service, this is a good place to
surface an explanation for the degradation (load, hardware failure
or other issue).

A unit in error state will have a message that is set by Juju and not
the charm because the error state represents a crash in a charm hook
- an unmanaged and uninterpretable situation. Juju will set the message
to be a reflection of the hook which crashed.
For example “Crashed installing the software” for an install hook crash
, or “Crash establishing database link” for a crash in a relationship hook.
`
	examples := `
Set the unit’s status

    # Set the unit's workload status to "maintenance".
    # This implies a short downtime that should self-resolve.
    status-set maintenance "installing software"
    status-set maintenance "formatting storage space, time left: 120s"

    # Set the unit's workload status to "waiting"
    # The workload is awaiting something else in the model to become active
    status-set waiting "waiting for database"

    # Set the unit workload's status to "active"
    # The workload is installed and running. Any messages should be informational.
    status-set active
    status-set active "Storage 95% full"

    # Set the unit's workload status to "blocked"
    # This implies human intervention is required to unblock the unit.
    # Messages should describe what is needed to resolve the problem.
    status-set blocked "Add a database relation"
    status-set blocked "Storage full"

Set the application’s status:

    # From a unit, update its status
    status-set maintenance "Upgrading to 4.1.1"

    # From the leader, update the application's status line
    status-set --application maintenance "Application upgrade underway"

Non-leader units which attempt to use --application will receive an error:

    $ status-set --application maintenance "I'm not the leader."
    error: this unit is not the leader
`
	return jujucmd.Info(&cmd.Info{
		Name:     "status-set",
		Args:     "<maintenance | blocked | waiting | active> [message]",
		Purpose:  "Set status information.",
		Doc:      doc,
		Examples: examples,
	})
}

var validStatus = []status.Status{
	status.Maintenance,
	status.Blocked,
	status.Waiting,
	status.Active,
}

func (c *StatusSetCommand) SetFlags(f *gnuflag.FlagSet) {
	f.BoolVar(&c.application, "application", false, "set this status for the application to which the unit belongs if the unit is the leader")
}

func (c *StatusSetCommand) Init(args []string) error {
	if len(args) < 1 {
		return errors.Errorf("invalid args, require <status> [message]")
	}
	valid := false
	for _, s := range validStatus {
		if string(s) == args[0] {
			valid = true
			break
		}
	}
	if !valid {
		return errors.Errorf("invalid status %q, expected one of %v", args[0], validStatus)
	}
	c.status = args[0]
	if len(args) > 1 {
		c.message = args[1]
		return cmd.CheckEmpty(args[2:])
	}
	return nil
}

func (c *StatusSetCommand) Run(ctx *cmd.Context) error {
	statusInfo := StatusInfo{
		Status: c.status,
		Info:   c.message,
	}
	if c.application {
		return c.ctx.SetApplicationStatus(ctx, statusInfo)
	}
	return c.ctx.SetUnitStatus(ctx, statusInfo)

}
