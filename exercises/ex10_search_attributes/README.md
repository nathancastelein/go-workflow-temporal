# Exercise 10 - Capture Expedition (Search Attributes)

In this exercise, you will implement a workflow that uses **Search Attributes** to tag workflow executions with custom metadata for filtering and observability.

## Concepts

### Search Attributes

Search Attributes are custom key-value metadata attached to workflow executions. They allow you to filter and search workflows in the Temporal UI and CLI.

#### Defining typed keys

```go
var (
    TrainerNameKey = temporal.NewSearchAttributeKeyKeyword("TrainerName")
    CaptureSuccessKey = temporal.NewSearchAttributeKeyBool("CaptureSuccess")
)
```

Available types: `Keyword`, `Int64`, `Bool`, `Float64`, `Time`, `KeywordList`

#### Upserting search attributes in a workflow

```go
err := workflow.UpsertTypedSearchAttributes(ctx,
    TrainerNameKey.ValueSet("Ash"),
    CaptureSuccessKey.ValueSet(true),
)
```

#### Filtering workflows via CLI

```bash
temporal workflow list --query 'TrainerName="Ash" AND CaptureSuccess=true'
```

## Your Task

### activities.go

Implement two activities:

- `EncounterInRegionActivity(region)` — look up the region in `pokemon.RegionPokemon`, return a random Pokemon from that region
- `AttemptCaptureActivity(pokemon)` — return `true` if the Pokemon's HP is less than 100

### workflow.go

Implement `CaptureExpeditionWorkflow` that:
1. Upserts `TrainerName` and `Region` search attributes at start
2. Calls `EncounterInRegionActivity` to encounter a Pokemon
3. Upserts `PokemonType` search attribute with the encountered Pokemon's type
4. Calls `AttemptCaptureActivity` with the encountered Pokemon
5. Upserts `CaptureSuccess` search attribute with the result
6. Returns an `ExpeditionResult` with all fields

## Validate

Run from the project root:

```bash
go test -v ./exercises/ex10_search_attributes/...
```

## How to Run

1. Start the Temporal dev server with custom search attributes:

```bash
temporal server start-dev \
  --search-attribute TrainerName=Keyword \
  --search-attribute PokemonType=Keyword \
  --search-attribute Region=Keyword \
  --search-attribute CaptureSuccess=Bool
```

2. Start the worker:

```bash
go run ./exercises/ex10_search_attributes/worker/
```

3. Start expedition workflows:

```bash
temporal workflow start --task-queue pokemon --type CaptureExpeditionWorkflow \
  --input '"Ash"' --input '"Kanto"' --workflow-id "expedition-ash-1"

temporal workflow start --task-queue pokemon --type CaptureExpeditionWorkflow \
  --input '"Misty"' --input '"Johto"' --workflow-id "expedition-misty-1"
```

4. Filter workflows by search attributes:

```bash
temporal workflow list --query 'TrainerName="Ash"'
temporal workflow list --query 'Region="Kanto"'
temporal workflow list --query 'CaptureSuccess=true'
```
