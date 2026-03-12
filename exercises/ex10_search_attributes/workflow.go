package ex10_search_attributes

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// Search attribute keys — these are provided for you
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
	// TODO: Create activity options with StartToCloseTimeout of 10 seconds
	//   ao := workflow.ActivityOptions{StartToCloseTimeout: 10 * time.Second}
	//   ctx = workflow.WithActivityOptions(ctx, ao)

	// TODO: Upsert initial search attributes (TrainerName and Region)
	//   workflow.UpsertTypedSearchAttributes(ctx,
	//       TrainerNameKey.ValueSet(trainerName),
	//       RegionKey.ValueSet(region),
	//   )

	// TODO: Call EncounterInRegionActivity with the region to encounter a Pokemon
	//   var encountered pokemon.Pokemon
	//   workflow.ExecuteActivity(ctx, EncounterInRegionActivity, region).Get(ctx, &encountered)

	// TODO: Upsert the PokemonType search attribute with the encountered Pokemon's type
	//   workflow.UpsertTypedSearchAttributes(ctx, PokemonTypeKey.ValueSet(encountered.Type))

	// TODO: Call AttemptCaptureActivity with the encountered Pokemon
	//   var success bool
	//   workflow.ExecuteActivity(ctx, AttemptCaptureActivity, encountered).Get(ctx, &success)

	// TODO: Upsert the CaptureSuccess search attribute
	//   workflow.UpsertTypedSearchAttributes(ctx, CaptureSuccessKey.ValueSet(success))

	// TODO: Return an ExpeditionResult with all fields

	_ = time.Second // remove when implementing
	_ = TrainerNameKey
	_ = PokemonTypeKey
	_ = RegionKey
	_ = CaptureSuccessKey
	return pokemon.ExpeditionResult{}, nil
}
