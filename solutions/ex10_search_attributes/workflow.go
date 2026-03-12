package ex10_search_attributes

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

var (
	TrainerNameKey    = temporal.NewSearchAttributeKeyKeyword("TrainerName")
	PokemonTypeKey    = temporal.NewSearchAttributeKeyKeyword("PokemonType")
	RegionKey         = temporal.NewSearchAttributeKeyKeyword("Region")
	CaptureSuccessKey = temporal.NewSearchAttributeKeyBool("CaptureSuccess")
)

// CaptureExpeditionWorkflow runs a capture expedition for a trainer in a given region.
// It sets search attributes at start, encounters a Pokemon, upserts its type,
// attempts capture, and upserts the success status.
func CaptureExpeditionWorkflow(ctx workflow.Context, trainerName string, region string) (pokemon.ExpeditionResult, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Set initial search attributes
	err := workflow.UpsertTypedSearchAttributes(ctx,
		TrainerNameKey.ValueSet(trainerName),
		RegionKey.ValueSet(region),
	)
	if err != nil {
		return pokemon.ExpeditionResult{}, err
	}

	// Encounter a Pokemon in the region
	var encountered pokemon.Pokemon
	err = workflow.ExecuteActivity(ctx, EncounterInRegionActivity, region).Get(ctx, &encountered)
	if err != nil {
		return pokemon.ExpeditionResult{}, err
	}

	// Upsert Pokemon type
	err = workflow.UpsertTypedSearchAttributes(ctx,
		PokemonTypeKey.ValueSet(encountered.Type),
	)
	if err != nil {
		return pokemon.ExpeditionResult{}, err
	}

	// Attempt capture
	var success bool
	err = workflow.ExecuteActivity(ctx, AttemptCaptureActivity, encountered).Get(ctx, &success)
	if err != nil {
		return pokemon.ExpeditionResult{}, err
	}

	// Upsert capture result
	err = workflow.UpsertTypedSearchAttributes(ctx,
		CaptureSuccessKey.ValueSet(success),
	)
	if err != nil {
		return pokemon.ExpeditionResult{}, err
	}

	return pokemon.ExpeditionResult{
		TrainerName: trainerName,
		Pokemon:     encountered,
		Region:      region,
		Success:     success,
	}, nil
}
