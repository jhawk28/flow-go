package fvm

import (
	"fmt"

	"github.com/opentracing/opentracing-go"
	traceLog "github.com/opentracing/opentracing-go/log"
	"github.com/rs/zerolog"

	"github.com/onflow/cadence"
	"github.com/onflow/cadence/runtime"

	"github.com/onflow/flow-go/fvm/blueprints"
	basicMeter "github.com/onflow/flow-go/fvm/meter/basic"
	weightedMeter "github.com/onflow/flow-go/fvm/meter/weighted"
	"github.com/onflow/flow-go/fvm/programs"
	"github.com/onflow/flow-go/fvm/state"
	"github.com/onflow/flow-go/fvm/utils"
	"github.com/onflow/flow-go/module/trace"
)

type TransactionSetExecutionWeights struct {
	logger zerolog.Logger
}

func NewTransactionSetExecutionWeights(logger zerolog.Logger) *TransactionSetExecutionWeights {
	return &TransactionSetExecutionWeights{
		logger: logger,
	}
}

func (c *TransactionSetExecutionWeights) Process(
	vm *VirtualMachine,
	ctx *Context,
	proc *TransactionProcedure,
	sth *state.StateHolder,
	programs *programs.Programs,
) error {
	txIDStr := proc.ID.String()
	var span opentracing.Span
	if ctx.Tracer != nil && proc.TraceSpan != nil {
		span = ctx.Tracer.StartSpanFromParent(proc.TraceSpan, trace.FVMSetExecutionWeights)
		span.LogFields(
			traceLog.String("transaction_id", txIDStr),
		)
		defer span.Finish()
	}

	l := c.logger.With().
		Str("txHash", txIDStr).
		Logger()
	env := NewTransactionEnvironment(*ctx, vm, sth, programs, proc.Transaction, proc.TxIndex, span)

	return setExecutionWeights(env, sth, l)
}

var _ TransactionProcessor = &TransactionSetExecutionWeights{}

type ScriptSetExecutionWeights struct {
	logger zerolog.Logger
}

func NewScriptSetExecutionWeights(logger zerolog.Logger) *ScriptSetExecutionWeights {
	return &ScriptSetExecutionWeights{
		logger: logger,
	}
}

func (c *ScriptSetExecutionWeights) Process(vm *VirtualMachine, ctx Context, _ *ScriptProcedure, sth *state.StateHolder, programs *programs.Programs) error {
	env := NewScriptEnvironment(ctx, vm, sth, programs)

	return setExecutionWeights(env, sth, c.logger)
}

var _ ScriptProcessor = &ScriptSetExecutionWeights{}

func setExecutionWeights(env Environment, sth *state.StateHolder, l zerolog.Logger) error {
	// do not meter getting execution weights
	sth.DisableAllLimitEnforcements()
	defer sth.EnableAllLimitEnforcements()

	// Get the meter to set the weights for
	m := sth.State().Meter()

	// Get the computation weights
	computationWeights, err := GetExecutionEffortWeights(env)

	if err != nil {
		// This could be a reason to error the transaction in the future, but for now we just log it
		l.Info().
			Err(err).
			Msg("failed to get execution effort weights")

		sth.State().SetMeter(basicMeter.NewMeter(
			m.TotalComputationLimit(),
			m.TotalMemoryLimit()))
		return nil
	}

	// Get the memory weights
	memoryWeights, err := GetExecutionMemoryWeights(env)

	if err != nil {
		// This could be a reason to error the transaction in the future, but for now we just log it
		l.Info().Err(err).
			Msg("failed to get execution memory weights")

		sth.State().SetMeter(basicMeter.NewMeter(
			m.TotalComputationLimit(),
			m.TotalMemoryLimit()))
		return nil
	}

	sth.State().SetMeter(weightedMeter.NewMeter(
		m.TotalComputationLimit(),
		m.TotalMemoryLimit(),
		computationWeights,
		memoryWeights))

	return nil
}

// GetExecutionEffortWeights reads stored execution effort weights from the service account
func GetExecutionEffortWeights(env Environment) (map[uint]uint64, error) {
	service := runtime.Address(env.Context().Chain.ServiceAddress())

	value, err := env.VM().Runtime.ReadStored(
		service,
		cadence.Path{
			Domain:     blueprints.TransactionFeesExecutionEffortWeightsPathDomain,
			Identifier: blueprints.TransactionFeesExecutionEffortWeightsPathIdentifier,
		},
		runtime.Context{Interface: env},
	)

	if err != nil {
		return map[uint]uint64{}, err
	}
	weights, ok := utils.CadenceValueToUintUintMap(value)
	if !ok {
		return map[uint]uint64{}, fmt.Errorf("could not decode stored execution effort weights")
	}

	return weights, nil
}

// GetExecutionMemoryWeights reads stored execution memory weights from the service account
func GetExecutionMemoryWeights(_ Environment) (map[uint]uint64, error) {
	// TODO: implement when memory metering is ready
	return map[uint]uint64{}, nil
}
