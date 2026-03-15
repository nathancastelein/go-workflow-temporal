package ex05_testing

import (
	"testing"
)

func TestFetchPokemonActivity_KnownPokemon(t *testing.T) {
	// TODO: Create a test activity environment using testsuite.WorkflowTestSuite{}
	// TODO: Register and execute FetchPokemonActivity with a known Pokemon name (e.g. "Pikachu")
	// TODO: Assert no error and the returned Pokemon has the correct name, type, and HP
	t.Fatal("implement this test")
}

func TestFetchPokemonActivity_UnknownPokemon(t *testing.T) {
	// TODO: Create a test activity environment using testsuite.WorkflowTestSuite{}
	// TODO: Register and execute FetchPokemonActivity with an unknown Pokemon name (e.g. "MissingNo")
	// TODO: Assert an error is returned
	t.Fatal("implement this test")
}
