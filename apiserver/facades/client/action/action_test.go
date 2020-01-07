// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package action_test

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"
	"time"

	jc "github.com/juju/testing/checkers"
	"github.com/kr/pretty"
	gc "gopkg.in/check.v1"
	"gopkg.in/juju/names.v3"

	"github.com/juju/juju/apiserver/common"
	commontesting "github.com/juju/juju/apiserver/common/testing"
	"github.com/juju/juju/apiserver/facades/client/action"
	"github.com/juju/juju/apiserver/params"
	apiservertesting "github.com/juju/juju/apiserver/testing"
	"github.com/juju/juju/core/actions"
	jujutesting "github.com/juju/juju/juju/testing"
	"github.com/juju/juju/state"
	statetesting "github.com/juju/juju/state/testing"
	coretesting "github.com/juju/juju/testing"
	"github.com/juju/juju/testing/factory"
)

func TestAll(t *testing.T) {
	coretesting.MgoTestPackage(t)
}

type actionSuite struct {
	jujutesting.JujuConnSuite
	commontesting.BlockHelper

	action     *action.ActionAPI
	authorizer apiservertesting.FakeAuthorizer
	resources  *common.Resources

	charm         *state.Charm
	machine0      *state.Machine
	machine1      *state.Machine
	dummy         *state.Application
	wordpress     *state.Application
	mysql         *state.Application
	wordpressUnit *state.Unit
	mysqlUnit     *state.Unit
}

var _ = gc.Suite(&actionSuite{})

func (s *actionSuite) SetUpTest(c *gc.C) {
	s.JujuConnSuite.SetUpTest(c)
	s.BlockHelper = commontesting.NewBlockHelper(s.APIState)
	s.AddCleanup(func(*gc.C) { s.BlockHelper.Close() })

	s.authorizer = apiservertesting.FakeAuthorizer{
		Tag: s.AdminUserTag(c),
	}
	s.resources = common.NewResources()
	s.AddCleanup(func(_ *gc.C) { s.resources.StopAll() })

	var err error
	s.action, err = action.NewActionAPI(s.State, s.resources, s.authorizer)
	c.Assert(err, jc.ErrorIsNil)

	s.charm = s.Factory.MakeCharm(c, &factory.CharmParams{
		Name: "wordpress",
	})

	s.dummy = s.Factory.MakeApplication(c, &factory.ApplicationParams{
		Name: "dummy",
		Charm: s.Factory.MakeCharm(c, &factory.CharmParams{
			Name: "dummy",
		}),
	})
	s.wordpress = s.Factory.MakeApplication(c, &factory.ApplicationParams{
		Name:  "wordpress",
		Charm: s.charm,
	})
	s.machine0 = s.Factory.MakeMachine(c, &factory.MachineParams{
		Series: "quantal",
		Jobs:   []state.MachineJob{state.JobHostUnits, state.JobManageModel},
	})
	s.wordpressUnit = s.Factory.MakeUnit(c, &factory.UnitParams{
		Application: s.wordpress,
		Machine:     s.machine0,
	})

	mysqlCharm := s.Factory.MakeCharm(c, &factory.CharmParams{
		Name: "mysql",
	})
	s.mysql = s.Factory.MakeApplication(c, &factory.ApplicationParams{
		Name:  "mysql",
		Charm: mysqlCharm,
	})
	s.machine1 = s.Factory.MakeMachine(c, &factory.MachineParams{
		Series: "quantal",
		Jobs:   []state.MachineJob{state.JobHostUnits},
	})
	s.mysqlUnit = s.Factory.MakeUnit(c, &factory.UnitParams{
		Application: s.mysql,
		Machine:     s.machine1,
	})
}

func (s *actionSuite) TestActions(c *gc.C) {
	arg := params.Actions{
		Actions: []params.Action{
			{Receiver: s.wordpressUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{}},
			{Receiver: s.mysqlUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{}},
			{Receiver: s.wordpressUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{"foo": 1, "bar": "please"}},
			{Receiver: s.mysqlUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{"baz": true}},
		}}

	r, err := s.action.Enqueue(arg)
	c.Assert(err, gc.Equals, nil)
	c.Assert(r.Results, gc.HasLen, len(arg.Actions))

	entities := make([]params.Entity, len(r.Results))
	for i, result := range r.Results {
		entities[i] = params.Entity{Tag: result.Action.Tag}
	}

	actions, err := s.action.Actions(params.Entities{Entities: entities})
	c.Assert(err, gc.Equals, nil)

	c.Assert(len(actions.Results), gc.Equals, len(entities))
	for i, got := range actions.Results {
		c.Logf("check index %d (%s: %s)", i, entities[i].Tag, arg.Actions[i].Name)
		c.Assert(got.Error, gc.Equals, (*params.Error)(nil))
		c.Assert(got.Action, gc.Not(gc.Equals), (*params.Action)(nil))
		c.Assert(got.Action.Tag, gc.Equals, entities[i].Tag)
		c.Assert(got.Action.Name, gc.Equals, arg.Actions[i].Name)
		c.Assert(got.Action.Receiver, gc.Equals, arg.Actions[i].Receiver)
		c.Assert(got.Action.Parameters, gc.DeepEquals, arg.Actions[i].Parameters)
		c.Assert(got.Status, gc.Equals, params.ActionPending)
		c.Assert(got.Message, gc.Equals, "")
		c.Assert(got.Output, gc.DeepEquals, map[string]interface{}{})
	}
}

func (s *actionSuite) TestFindActionTagsByPrefix(c *gc.C) {
	// NOTE: full testing with multiple matches has been moved to state package.
	arg := params.Actions{Actions: []params.Action{{Receiver: s.wordpressUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{}}}}
	r, err := s.action.Enqueue(arg)
	c.Assert(err, gc.Equals, nil)
	c.Assert(r.Results, gc.HasLen, len(arg.Actions))

	actionTag, err := names.ParseActionTag(r.Results[0].Action.Tag)
	c.Assert(err, gc.Equals, nil)
	prefix := actionTag.Id()
	tags, err := s.action.FindActionTagsByPrefix(params.FindTags{Prefixes: []string{prefix}})
	c.Assert(err, gc.Equals, nil)

	entities, ok := tags.Matches[prefix]
	c.Assert(ok, gc.Equals, true)
	c.Assert(len(entities), gc.Equals, 1)
	c.Assert(entities[0].Tag, gc.Equals, actionTag.String())
}

func (s *actionSuite) TestFindActionsByName(c *gc.C) {
	machine := s.JujuConnSuite.Factory.MakeMachine(c, &factory.MachineParams{
		Series: "quantal",
		Jobs:   []state.MachineJob{state.JobHostUnits},
	})
	dummyUnit := s.JujuConnSuite.Factory.MakeUnit(c, &factory.UnitParams{
		Application: s.dummy,
		Machine:     machine,
	})
	// NOTE: full testing with multiple matches has been moved to state package.
	arg := params.Actions{Actions: []params.Action{
		{Receiver: s.wordpressUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{}},
		{Receiver: dummyUnit.Tag().String(), Name: "snapshot", Parameters: map[string]interface{}{"outfile": "lol"}},
		{Receiver: s.wordpressUnit.Tag().String(), Name: "juju-run", Parameters: map[string]interface{}{"command": "boo", "timeout": 5}},
		{Receiver: s.mysqlUnit.Tag().String(), Name: "juju-run", Parameters: map[string]interface{}{"command": "boo", "timeout": 5}},
	}}
	r, err := s.action.Enqueue(arg)
	c.Assert(err, gc.Equals, nil)
	c.Assert(r.Results, gc.HasLen, len(arg.Actions))

	actionNames := []string{"snapshot", "juju-run"}
	actions, err := s.action.FindActionsByNames(params.FindActionsByNames{ActionNames: actionNames})
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(len(actions.Actions), gc.Equals, 2)
	for i, actions := range actions.Actions {
		for _, act := range actions.Actions {
			c.Assert(act.Action.Name, gc.Equals, actionNames[i])
			c.Assert(act.Action.Name, gc.Matches, actions.Name)
		}
	}
}

func (s *actionSuite) TestEnqueue(c *gc.C) {
	// Ensure wordpress unit is the leader.
	claimer, err := s.LeaseManager.Claimer("application-leadership", s.State.ModelUUID())
	c.Assert(err, jc.ErrorIsNil)
	err = claimer.Claim("wordpress", "wordpress/0", time.Minute)
	c.Assert(err, jc.ErrorIsNil)

	// Make sure no Actions already exist on wordpress Unit.
	actions, err := s.wordpressUnit.Actions()
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(actions, gc.HasLen, 0)

	// Make sure no Actions already exist on mysql Unit.
	actions, err = s.mysqlUnit.Actions()
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(actions, gc.HasLen, 0)

	// Add Actions.
	expectedName := "fakeaction"
	expectedParameters := map[string]interface{}{"kan jy nie": "verstaand"}
	arg := params.Actions{
		Actions: []params.Action{
			// No receiver.
			{Name: "fakeaction"},
			// Good.
			{Receiver: s.wordpressUnit.Tag().String(), Name: expectedName, Parameters: expectedParameters},
			// Application tag instead of Unit tag.
			{Receiver: s.wordpress.Tag().String(), Name: "fakeaction"},
			// Missing name.
			{Receiver: s.mysqlUnit.Tag().String(), Parameters: expectedParameters},
			// Good (leader syntax).
			{Receiver: "wordpress/leader", Name: expectedName, Parameters: expectedParameters},
		},
	}

	// blocking changes should have no effect
	s.BlockAllChanges(c, "Enqueue")

	res, err := s.action.Enqueue(arg)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(res.Results, gc.HasLen, 5)

	emptyActionTag := names.ActionTag{}
	c.Assert(res.Results[0].Error, gc.DeepEquals,
		&params.Error{Message: fmt.Sprintf("%s not valid", arg.Actions[0].Receiver), Code: ""})
	c.Assert(res.Results[0].Action, gc.IsNil)

	c.Assert(res.Results[1].Error, gc.IsNil)
	c.Assert(res.Results[1].Action, gc.NotNil)
	c.Assert(res.Results[1].Action.Receiver, gc.Equals, s.wordpressUnit.Tag().String())
	c.Assert(res.Results[1].Action.Tag, gc.Not(gc.Equals), emptyActionTag)

	errorString := fmt.Sprintf("action receiver interface on entity %s not implemented", arg.Actions[2].Receiver)
	c.Assert(res.Results[2].Error, gc.DeepEquals, &params.Error{Message: errorString, Code: "not implemented"})
	c.Assert(res.Results[2].Action, gc.IsNil)

	c.Assert(res.Results[3].Error, gc.ErrorMatches, "no action name given")
	c.Assert(res.Results[3].Action, gc.IsNil)

	c.Assert(res.Results[4].Error, gc.IsNil)
	c.Assert(res.Results[4].Action, gc.NotNil)
	c.Assert(res.Results[4].Action.Receiver, gc.Equals, s.wordpressUnit.Tag().String())
	c.Assert(res.Results[4].Action.Tag, gc.Not(gc.Equals), emptyActionTag)

	// Make sure that 2 actions were enqueued for the wordpress Unit.
	actions, err = s.wordpressUnit.Actions()
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(actions, gc.HasLen, 2)
	for _, act := range actions {
		c.Assert(act.Name(), gc.Equals, expectedName)
		c.Assert(act.Parameters(), gc.DeepEquals, expectedParameters)
		c.Assert(act.Receiver(), gc.Equals, s.wordpressUnit.Name())
	}

	// Make sure an Action was not enqueued for the mysql Unit.
	actions, err = s.mysqlUnit.Actions()
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(actions, gc.HasLen, 0)
}

type testCaseAction struct {
	Name       string
	Parameters map[string]interface{}
	Execute    bool
}

type receiverGroup struct {
	ExpectedError *params.Error
	Receiver      names.Tag
	Actions       []testCaseAction
}

type testCase struct {
	Groups []receiverGroup
}

var testCases = []testCase{{
	Groups: []receiverGroup{
		{
			ExpectedError: &params.Error{Message: "id not found", Code: "not found"},
			Receiver:      names.NewApplicationTag("wordpress"),
			Actions:       []testCaseAction{},
		}, {
			Receiver: names.NewUnitTag("wordpress/0"),
			Actions: []testCaseAction{
				{"fakeaction", map[string]interface{}{}, false},
				{"fakeaction", map[string]interface{}{"asdf": 3}, true},
				{"fakeaction", map[string]interface{}{"qwer": "ty"}, false},
			},
		}, {
			Receiver: names.NewUnitTag("mysql/0"),
			Actions: []testCaseAction{
				{"fakeaction", map[string]interface{}{"zxcv": false}, false},
				{"fakeaction", map[string]interface{}{}, true},
			},
		},
	},
}}

func (s *actionSuite) TestListAll(c *gc.C) {
	for _, testCase := range testCases {
		// set up query args
		arg := params.Entities{Entities: make([]params.Entity, len(testCase.Groups))}

		// prepare state, and set up expectations.
		expected := params.ActionsByReceivers{Actions: make([]params.ActionsByReceiver, len(testCase.Groups))}
		for i, group := range testCase.Groups {
			arg.Entities[i] = params.Entity{Tag: group.Receiver.String()}

			cur := &expected.Actions[i]
			cur.Error = group.ExpectedError

			// short circuit and bail if the ActionReceiver isn't a Unit.
			if _, ok := group.Receiver.(names.UnitTag); !ok {
				continue
			}

			cur.Receiver = group.Receiver.String()
			cur.Actions = make([]params.ActionResult, len(group.Actions))

			// get Unit (ActionReceiver) for this Pair in the test case.
			unit, err := s.State.Unit(group.Receiver.Id())
			c.Assert(err, jc.ErrorIsNil)
			assertReadyToTest(c, unit)

			// add each action from the test case.
			for j, act := range group.Actions {
				// add action.
				added, err := unit.AddAction(act.Name, act.Parameters)
				c.Assert(err, jc.ErrorIsNil)

				// make expectation
				exp := &cur.Actions[j]
				exp.Action = &params.Action{
					Tag:        added.ActionTag().String(),
					Name:       act.Name,
					Parameters: act.Parameters,
				}
				exp.Status = params.ActionPending
				exp.Output = map[string]interface{}{}

				if act.Execute {
					added, err = added.Begin()
					c.Assert(err, jc.ErrorIsNil)
					err = added.Log("hello")
					c.Assert(err, jc.ErrorIsNil)
					status := state.ActionCompleted
					output := map[string]interface{}{
						"output":         "blah, blah, blah",
						"Stdout":         "out",
						"StdoutEncoding": "utf-8",
						"Stderr":         "err",
						"StderrEncoding": "utf-8",
						"Code":           "1",
					}
					message := "success"

					fa, err := added.Finish(state.ActionResults{Status: status, Results: output, Message: message})
					c.Assert(err, jc.ErrorIsNil)
					c.Assert(fa.Status(), gc.Equals, state.ActionCompleted)

					exp.Status = string(status)
					exp.Message = message
					exp.Output = map[string]interface{}{
						"output":          "blah, blah, blah",
						"stdout":          "out",
						"stdout-encoding": "utf-8",
						"stderr":          "err",
						"stderr-encoding": "utf-8",
						"return-code":     1,
					}
					exp.Log = []params.ActionMessage{{Message: "hello"}}
				}
			}
		}

		// validate assumptions.
		actionList, err := s.action.ListAll(arg)
		c.Assert(err, jc.ErrorIsNil)
		assertSame(c, actionList, expected)
	}
}

func (s *actionSuite) TestListPending(c *gc.C) {
	for _, testCase := range testCases {
		// set up query args
		arg := params.Entities{Entities: make([]params.Entity, len(testCase.Groups))}

		// prepare state, and set up expectations.
		expected := params.ActionsByReceivers{Actions: make([]params.ActionsByReceiver, len(testCase.Groups))}
		for i, group := range testCase.Groups {
			arg.Entities[i] = params.Entity{Tag: group.Receiver.String()}

			cur := &expected.Actions[i]
			cur.Error = group.ExpectedError

			// short circuit and bail if the ActionReceiver isn't a Unit.
			if _, ok := group.Receiver.(names.UnitTag); !ok {
				continue
			}

			cur.Receiver = group.Receiver.String()
			cur.Actions = []params.ActionResult{}

			// get Unit (ActionReceiver) for this Pair in the test case.
			unit, err := s.State.Unit(group.Receiver.Id())
			c.Assert(err, jc.ErrorIsNil)
			assertReadyToTest(c, unit)

			// add each action from the test case.
			for _, act := range group.Actions {
				// add action.
				added, err := unit.AddAction(act.Name, act.Parameters)
				c.Assert(err, jc.ErrorIsNil)

				if act.Execute {
					status := state.ActionCompleted
					output := map[string]interface{}{"output": "blah, blah, blah"}
					message := "success"

					fa, err := added.Finish(state.ActionResults{Status: status, Results: output, Message: message})
					c.Assert(err, jc.ErrorIsNil)
					c.Assert(fa.Status(), gc.Equals, state.ActionCompleted)
				} else {
					// add expectation
					exp := params.ActionResult{
						Action: &params.Action{
							Tag:        added.ActionTag().String(),
							Name:       act.Name,
							Parameters: act.Parameters,
						},
						Status: params.ActionPending,
						Output: map[string]interface{}{},
					}
					cur.Actions = append(cur.Actions, exp)
				}
			}
		}

		// validate assumptions.
		actionList, err := s.action.ListPending(arg)
		c.Assert(err, jc.ErrorIsNil)
		assertSame(c, actionList, expected)
	}
}

func (s *actionSuite) TestListRunning(c *gc.C) {
	for _, testCase := range testCases {
		// set up query args
		arg := params.Entities{Entities: make([]params.Entity, len(testCase.Groups))}

		// prepare state, and set up expectations.
		expected := params.ActionsByReceivers{Actions: make([]params.ActionsByReceiver, len(testCase.Groups))}
		for i, group := range testCase.Groups {
			arg.Entities[i] = params.Entity{Tag: group.Receiver.String()}

			cur := &expected.Actions[i]
			cur.Error = group.ExpectedError

			// short circuit and bail if the ActionReceiver isn't a Unit.
			if _, ok := group.Receiver.(names.UnitTag); !ok {
				continue
			}

			cur.Receiver = group.Receiver.String()
			cur.Actions = []params.ActionResult{}

			// get Unit (ActionReceiver) for this Pair in the test case.
			unit, err := s.State.Unit(group.Receiver.Id())
			c.Assert(err, jc.ErrorIsNil)
			assertReadyToTest(c, unit)

			// add each action from the test case.
			for _, act := range group.Actions {
				// add action.
				added, err := unit.AddAction(act.Name, act.Parameters)
				c.Assert(err, jc.ErrorIsNil)

				if act.Execute {
					started, err := added.Begin()
					c.Assert(err, jc.ErrorIsNil)
					c.Assert(started.Status(), gc.Equals, state.ActionRunning)

					// add expectation
					exp := params.ActionResult{
						Action: &params.Action{
							Tag:        added.ActionTag().String(),
							Name:       act.Name,
							Parameters: act.Parameters,
						},
						Status: params.ActionRunning,
						Output: map[string]interface{}{},
					}
					cur.Actions = append(cur.Actions, exp)
				}
			}
		}

		// validate assumptions.
		actionList, err := s.action.ListRunning(arg)
		c.Assert(err, jc.ErrorIsNil)
		assertSame(c, actionList, expected)
	}
}

func (s *actionSuite) TestListCompleted(c *gc.C) {
	for _, testCase := range testCases {
		// set up query args
		arg := params.Entities{Entities: make([]params.Entity, len(testCase.Groups))}

		// prepare state, and set up expectations.
		expected := params.ActionsByReceivers{Actions: make([]params.ActionsByReceiver, len(testCase.Groups))}
		for i, group := range testCase.Groups {
			arg.Entities[i] = params.Entity{Tag: group.Receiver.String()}

			cur := &expected.Actions[i]
			cur.Error = group.ExpectedError

			// short circuit and bail if the ActionReceiver isn't a Unit.
			if _, ok := group.Receiver.(names.UnitTag); !ok {
				continue
			}

			cur.Receiver = group.Receiver.String()
			cur.Actions = []params.ActionResult{}

			// get Unit (ActionReceiver) for this Pair in the test case.
			unit, err := s.State.Unit(group.Receiver.Id())
			c.Assert(err, jc.ErrorIsNil)
			assertReadyToTest(c, unit)

			// add each action from the test case.
			for _, act := range group.Actions {
				// add action.
				added, err := unit.AddAction(act.Name, act.Parameters)
				c.Assert(err, jc.ErrorIsNil)

				if act.Execute {
					status := state.ActionCompleted
					output := map[string]interface{}{
						"output":         "blah, blah, blah",
						"Stdout":         "out",
						"StdoutEncoding": "utf-8",
						"Stderr":         "err",
						"StderrEncoding": "utf-8",
						"Code":           "1",
					}
					message := "success"

					_, err = added.Finish(state.ActionResults{Status: status, Results: output, Message: message})
					c.Assert(err, jc.ErrorIsNil)

					// add expectation
					exp := params.ActionResult{
						Action: &params.Action{
							Tag:        added.ActionTag().String(),
							Name:       act.Name,
							Parameters: act.Parameters,
						},
						Status:  string(status),
						Message: message,
						Output: map[string]interface{}{
							"output":          "blah, blah, blah",
							"stdout":          "out",
							"stdout-encoding": "utf-8",
							"stderr":          "err",
							"stderr-encoding": "utf-8",
							"return-code":     1,
						},
					}
					cur.Actions = append(cur.Actions, exp)
				}
			}
		}

		// validate assumptions.
		actionList, err := s.action.ListCompleted(arg)
		c.Assert(err, jc.ErrorIsNil)
		assertSame(c, actionList, expected)
	}
}

func (s *actionSuite) TestCancel(c *gc.C) {
	// Make sure no Actions already exist on wordpress Unit.
	actions, err := s.wordpressUnit.Actions()
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(actions, gc.HasLen, 0)

	// Make sure no Actions already exist on mysql Unit.
	actions, err = s.mysqlUnit.Actions()
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(actions, gc.HasLen, 0)

	// Add Actions.
	tests := params.Actions{
		Actions: []params.Action{{
			Receiver: s.wordpressUnit.Tag().String(),
			Name:     "fakeaction",
		}, {
			Receiver: s.wordpressUnit.Tag().String(),
			Name:     "fakeaction",
		}, {
			Receiver: s.mysqlUnit.Tag().String(),
			Name:     "fakeaction",
		}, {
			Receiver: s.mysqlUnit.Tag().String(),
			Name:     "fakeaction",
		}},
	}

	results, err := s.action.Enqueue(tests)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results.Results, gc.HasLen, 4)
	for _, res := range results.Results {
		c.Assert(res.Error, gc.IsNil)
	}

	// blocking changes should have no effect
	s.BlockAllChanges(c, "Cancel")

	// Cancel Some.
	arg := params.Entities{
		Entities: []params.Entity{
			// "wp-two"
			{Tag: results.Results[1].Action.Tag},
			// "my-one"
			{Tag: results.Results[2].Action.Tag},
		}}
	results, err = s.action.Cancel(arg)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results.Results, gc.HasLen, 2)

	// Assert the Actions are all in the expected state.
	tags := params.Entities{Entities: []params.Entity{
		{Tag: s.wordpressUnit.Tag().String()},
		{Tag: s.mysqlUnit.Tag().String()},
	}}
	obtained, err := s.action.ListAll(tags)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(obtained.Actions, gc.HasLen, 2)

	wpActions := obtained.Actions[0].Actions
	c.Assert(wpActions, gc.HasLen, 2)
	c.Assert(wpActions[0].Action.Name, gc.Equals, "fakeaction")
	c.Assert(wpActions[0].Status, gc.Equals, params.ActionPending)
	c.Assert(wpActions[1].Action.Name, gc.Equals, "fakeaction")
	c.Assert(wpActions[1].Status, gc.Equals, params.ActionCancelled)

	myActions := obtained.Actions[1].Actions
	c.Assert(myActions, gc.HasLen, 2)
	c.Assert(myActions[0].Action.Name, gc.Equals, "fakeaction")
	c.Assert(myActions[0].Status, gc.Equals, params.ActionPending)
	c.Assert(myActions[1].Action.Name, gc.Equals, "fakeaction")
	c.Assert(myActions[1].Status, gc.Equals, params.ActionCancelled)
}

func (s *actionSuite) TestApplicationsCharmsActions(c *gc.C) {
	actionSchemas := map[string]map[string]interface{}{
		"snapshot": {
			"type":        "object",
			"title":       "snapshot",
			"description": "Take a snapshot of the database.",
			"properties": map[string]interface{}{
				"outfile": map[string]interface{}{
					"description": "The file to write out to.",
					"type":        "string",
					"default":     "foo.bz2",
				},
			},
		},
		"fakeaction": {
			"type":        "object",
			"title":       "fakeaction",
			"description": "No description",
			"properties":  map[string]interface{}{},
		},
	}
	tests := []struct {
		applicationNames []string
		expectedResults  params.ApplicationsCharmActionsResults
	}{{
		applicationNames: []string{"dummy"},
		expectedResults: params.ApplicationsCharmActionsResults{
			Results: []params.ApplicationCharmActionsResult{
				{
					ApplicationTag: names.NewApplicationTag("dummy").String(),
					Actions: map[string]params.ActionSpec{
						"snapshot": {
							Description: "Take a snapshot of the database.",
							Params:      actionSchemas["snapshot"],
						},
					},
				},
			},
		},
	}, {
		applicationNames: []string{"wordpress"},
		expectedResults: params.ApplicationsCharmActionsResults{
			Results: []params.ApplicationCharmActionsResult{
				{
					ApplicationTag: names.NewApplicationTag("wordpress").String(),
					Actions: map[string]params.ActionSpec{
						"fakeaction": {
							Description: "No description",
							Params:      actionSchemas["fakeaction"],
						},
					},
				},
			},
		},
	}, {
		applicationNames: []string{"nonsense"},
		expectedResults: params.ApplicationsCharmActionsResults{
			Results: []params.ApplicationCharmActionsResult{
				{
					ApplicationTag: names.NewApplicationTag("nonsense").String(),
					Error: &params.Error{
						Message: `application "nonsense" not found`,
						Code:    "not found",
					},
				},
			},
		},
	}}

	for i, t := range tests {
		c.Logf("test %d: applications: %#v", i, t.applicationNames)

		svcTags := params.Entities{
			Entities: make([]params.Entity, len(t.applicationNames)),
		}

		for j, app := range t.applicationNames {
			svcTag := names.NewApplicationTag(app)
			svcTags.Entities[j] = params.Entity{Tag: svcTag.String()}
		}

		results, err := s.action.ApplicationsCharmsActions(svcTags)
		c.Assert(err, jc.ErrorIsNil)
		c.Check(results.Results, jc.DeepEquals, t.expectedResults.Results)
	}
}

func assertReadyToTest(c *gc.C, receiver state.ActionReceiver) {
	// make sure there are no actions on the receiver already.
	actions, err := receiver.Actions()
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(actions, gc.HasLen, 0)

	// make sure there are no actions pending already.
	actions, err = receiver.PendingActions()
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(actions, gc.HasLen, 0)

	// make sure there are no actions running already.
	actions, err = receiver.RunningActions()
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(actions, gc.HasLen, 0)

	// make sure there are no actions completed already.
	actions, err = receiver.CompletedActions()
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(actions, gc.HasLen, 0)
}

func assertSame(c *gc.C, got, expected params.ActionsByReceivers) {
	c.Assert(got.Actions, gc.HasLen, len(expected.Actions))
	for i, g1 := range got.Actions {
		e1 := expected.Actions[i]
		c.Assert(g1.Error, gc.DeepEquals, e1.Error)
		c.Assert(g1.Receiver, gc.DeepEquals, e1.Receiver)
		for _, a1 := range g1.Actions {
			for _, m := range a1.Log {
				c.Assert(m.Timestamp.IsZero(), jc.IsFalse)
				m.Timestamp = time.Time{}
			}
		}
		c.Assert(toStrings(g1.Actions), jc.SameContents, toStrings(e1.Actions))
	}
}

func toStrings(items []params.ActionResult) []string {
	ret := make([]string, len(items))
	for i, a := range items {
		ret[i] = stringify(a)
	}
	return ret
}

func stringify(r params.ActionResult) string {
	a := r.Action
	if a == nil {
		a = &params.Action{}
	}
	// Convert action output map to ordered result.
	var keys, orderedOut []string
	for k := range r.Output {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		orderedOut = append(orderedOut, fmt.Sprintf("%v=%v", k, r.Output[k]))
	}
	return fmt.Sprintf("%s-%s-%#v-%s-%s-%v", a.Tag, a.Name, a.Parameters, r.Status, r.Message, orderedOut)
}

func (s *actionSuite) toSupportNewActionID(c *gc.C) {
	ver, err := s.Model.AgentVersion()
	c.Assert(err, jc.ErrorIsNil)

	if !state.IsNewActionIDSupported(ver) {
		err := s.State.SetModelAgentVersion(state.MinVersionSupportNewActionID, true)
		c.Assert(err, jc.ErrorIsNil)
	}
}

func (s *actionSuite) TestWatchActionProgress(c *gc.C) {
	s.toSupportNewActionID(c)

	unit, err := s.State.Unit("mysql/0")
	c.Assert(err, jc.ErrorIsNil)
	assertReadyToTest(c, unit)

	added, err := unit.AddAction("fakeaction", nil)
	c.Assert(err, jc.ErrorIsNil)

	w, err := s.action.WatchActionsProgress(
		params.Entities{Entities: []params.Entity{{Tag: "action-1"}}},
	)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(w.Results, gc.HasLen, 1)
	c.Assert(w.Results[0].Error, gc.IsNil)
	c.Assert(w.Results[0].Changes, gc.HasLen, 0)

	// Verify the resource was registered and stop when done
	c.Assert(s.resources.Count(), gc.Equals, 1)
	resource := s.resources.Get("1")
	defer statetesting.AssertStop(c, resource)

	// Check that the Watch has consumed the initial event
	wc := statetesting.NewStringsWatcherC(c, s.State, resource.(state.StringsWatcher))
	wc.AssertNoChange()

	// Log a message and check the watcher result.
	added, err = added.Begin()
	c.Assert(err, jc.ErrorIsNil)
	err = added.Log("hello")
	c.Assert(err, jc.ErrorIsNil)

	a, err := s.Model.Action("1")
	c.Assert(err, jc.ErrorIsNil)
	logged := a.Messages()
	c.Assert(logged, gc.HasLen, 1)
	expected, err := json.Marshal(actions.ActionMessage{
		Message:   logged[0].Message(),
		Timestamp: logged[0].Timestamp(),
	})
	c.Assert(err, jc.ErrorIsNil)

	wc.AssertChange(string(expected))
	wc.AssertNoChange()
}

func (s *actionSuite) setupOperations(c *gc.C) {
	s.toSupportNewActionID(c)

	arg := params.Actions{
		Actions: []params.Action{
			{Receiver: s.wordpressUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{}},
			{Receiver: s.mysqlUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{}},
			{Receiver: s.wordpressUnit.Tag().String(), Name: "fakeaction", Parameters: map[string]interface{}{}},
			{Receiver: s.mysqlUnit.Tag().String(), Name: "anotherfakeaction", Parameters: map[string]interface{}{}},
		}}

	r, err := s.action.Enqueue(arg)
	c.Assert(err, gc.Equals, nil)
	c.Assert(r.Results, gc.HasLen, len(arg.Actions))
	a, err := s.Model.Action("1")
	c.Assert(err, jc.ErrorIsNil)
	_, err = a.Begin()
	c.Assert(err, jc.ErrorIsNil)
	a, err = s.Model.Action("2")
	c.Assert(err, jc.ErrorIsNil)
	_, err = a.Finish(state.ActionResults{})
	c.Assert(err, jc.ErrorIsNil)
}

func (s *actionSuite) TestOperationsStatusFilter(c *gc.C) {
	s.setupOperations(c)
	actions, err := s.action.Operations(params.OperationQueryArgs{
		Status: []string{"running"},
	})
	c.Assert(err, gc.Equals, nil)
	c.Assert(actions.Results, gc.HasLen, 1)
	result := actions.Results[0]
	c.Assert(result.Action, gc.NotNil)
	if result.Enqueued.IsZero() {
		c.Fatal("enqueued time not set")
	}
	if result.Started.IsZero() {
		c.Fatal("started time not set")
	}
	c.Assert(result.Status, gc.Equals, "running")
	c.Assert(result.Action.Name, gc.Equals, "fakeaction")
	c.Assert(result.Action.Receiver, gc.Equals, "unit-wordpress-0")
	c.Assert(result.Action.Tag, gc.Equals, "action-1")
}

func (s *actionSuite) TestOperationsNameFilter(c *gc.C) {
	s.setupOperations(c)
	actions, err := s.action.Operations(params.OperationQueryArgs{
		FunctionNames: []string{"anotherfakeaction"},
	})
	c.Assert(err, gc.Equals, nil)
	c.Assert(actions.Results, gc.HasLen, 1)
	result := actions.Results[0]
	c.Assert(result.Action, gc.NotNil)
	if result.Enqueued.IsZero() {
		c.Fatal("enqueued time not set")
	}
	c.Assert(result.Status, gc.Equals, "pending")
	c.Assert(result.Action.Name, gc.Equals, "anotherfakeaction")
	c.Assert(result.Action.Receiver, gc.Equals, "unit-mysql-0")
	c.Assert(result.Action.Tag, gc.Equals, "action-4")
}

func (s *actionSuite) TestOperationsAppFilter(c *gc.C) {
	s.setupOperations(c)
	actions, err := s.action.Operations(params.OperationQueryArgs{
		Applications: []string{"wordpress"},
	})
	c.Assert(err, gc.Equals, nil)
	c.Assert(actions.Results, gc.HasLen, 2)
	result0 := actions.Results[0]
	result1 := actions.Results[1]

	c.Assert(result0.Action, gc.NotNil)
	if result0.Enqueued.IsZero() {
		c.Fatal("enqueued time not set")
	}
	c.Assert(result0.Status, gc.Equals, "pending")
	c.Assert(result0.Action.Name, gc.Equals, "fakeaction")
	c.Assert(result0.Action.Receiver, gc.Equals, "unit-wordpress-0")
	c.Assert(result0.Action.Tag, gc.Equals, "action-3")

	c.Assert(result1.Action, gc.NotNil)
	if result1.Enqueued.IsZero() {
		c.Fatal("enqueued time not set")
	}
	if result1.Started.IsZero() {
		c.Fatal("started time not set")
	}
	c.Assert(result1.Status, gc.Equals, "running")
	c.Assert(result1.Action.Name, gc.Equals, "fakeaction")
	c.Assert(result1.Action.Receiver, gc.Equals, "unit-wordpress-0")
	c.Assert(result1.Action.Tag, gc.Equals, "action-1")
}

func (s *actionSuite) TestOperationsUnitFilter(c *gc.C) {
	s.setupOperations(c)
	actions, err := s.action.Operations(params.OperationQueryArgs{
		Units:  []string{"wordpress/0"},
		Status: []string{"pending"},
	})
	c.Assert(err, gc.Equals, nil)
	c.Assert(actions.Results, gc.HasLen, 1)
	result := actions.Results[0]

	c.Assert(result.Action, gc.NotNil)
	if result.Enqueued.IsZero() {
		c.Fatal("enqueued time not set")
	}
	c.Assert(result.Status, gc.Equals, "pending")
	c.Assert(result.Action.Name, gc.Equals, "fakeaction")
	c.Assert(result.Action.Receiver, gc.Equals, "unit-wordpress-0")
	c.Assert(result.Action.Tag, gc.Equals, "action-3")
}

func (s *actionSuite) TestOperationsAppAndUnitFilter(c *gc.C) {
	s.setupOperations(c)
	actions, err := s.action.Operations(params.OperationQueryArgs{
		Applications: []string{"mysql"},
		Units:        []string{"wordpress/0"},
		Status:       []string{"pending"},
	})
	c.Assert(err, gc.Equals, nil)
	c.Assert(actions.Results, gc.HasLen, 2)
	mysqlAction := actions.Results[0]
	wordpressAction := actions.Results[1]
	c.Log(pretty.Sprint(actions.Results))

	c.Assert(mysqlAction.Action, gc.NotNil)
	if mysqlAction.Enqueued.IsZero() {
		c.Fatal("enqueued time not set")
	}
	c.Assert(mysqlAction.Status, gc.Equals, "pending")
	c.Assert(mysqlAction.Action.Name, gc.Equals, "anotherfakeaction")
	c.Assert(mysqlAction.Action.Receiver, gc.Equals, "unit-mysql-0")
	c.Assert(mysqlAction.Action.Tag, gc.Equals, "action-4")

	c.Assert(wordpressAction.Action, gc.NotNil)
	if wordpressAction.Enqueued.IsZero() {
		c.Fatal("enqueued time not set")
	}
	c.Assert(wordpressAction.Status, gc.Equals, "pending")
	c.Assert(wordpressAction.Action.Name, gc.Equals, "fakeaction")
	c.Assert(wordpressAction.Action.Receiver, gc.Equals, "unit-wordpress-0")
	c.Assert(wordpressAction.Action.Tag, gc.Equals, "action-3")

}
