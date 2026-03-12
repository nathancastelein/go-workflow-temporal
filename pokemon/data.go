package pokemon

// TaskQueue is the Temporal task queue name used across all exercises.
const TaskQueue = "pokemon"

// AllPokemon is the hardcoded list of available Pokemon.
var AllPokemon = []Pokemon{
	{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35},
	{Name: "Charmander", Type: "Fire", HP: 39, MaxHP: 39},
	{Name: "Bulbasaur", Type: "Grass", HP: 45, MaxHP: 45},
	{Name: "Squirtle", Type: "Water", HP: 44, MaxHP: 44},
	{Name: "Jigglypuff", Type: "Normal", HP: 115, MaxHP: 115},
	{Name: "Eevee", Type: "Normal", HP: 55, MaxHP: 55},
	{Name: "Snorlax", Type: "Normal", HP: 160, MaxHP: 160},
	{Name: "Geodude", Type: "Rock", HP: 40, MaxHP: 40},
	{Name: "Machop", Type: "Fighting", HP: 70, MaxHP: 70},
	{Name: "Gastly", Type: "Ghost", HP: 30, MaxHP: 30},
}

// EvolutionMap maps a Pokemon name to its evolved form.
var EvolutionMap = map[string]Pokemon{
	"Charmander": {Name: "Charmeleon", Type: "Fire", HP: 58, MaxHP: 58},
	"Bulbasaur":  {Name: "Ivysaur", Type: "Grass", HP: 60, MaxHP: 60},
	"Squirtle":   {Name: "Wartortle", Type: "Water", HP: 59, MaxHP: 59},
	"Pikachu":    {Name: "Raichu", Type: "Electric", HP: 60, MaxHP: 60},
	"Eevee":      {Name: "Vaporeon", Type: "Water", HP: 130, MaxHP: 130},
	"Machop":     {Name: "Machoke", Type: "Fighting", HP: 80, MaxHP: 80},
	"Gastly":     {Name: "Haunter", Type: "Ghost", HP: 45, MaxHP: 45},
}

// TrainerTeams maps trainer names to their signature Pokemon.
var TrainerTeams = map[string]Pokemon{
	"Ash":      {Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35},
	"Misty":    {Name: "Squirtle", Type: "Water", HP: 44, MaxHP: 44},
	"Brock":    {Name: "Geodude", Type: "Rock", HP: 40, MaxHP: 40},
	"Gary":     {Name: "Eevee", Type: "Normal", HP: 55, MaxHP: 55},
	"Jessie":   {Name: "Jigglypuff", Type: "Normal", HP: 115, MaxHP: 115},
	"James":    {Name: "Gastly", Type: "Ghost", HP: 30, MaxHP: 30},
	"Sabrina":  {Name: "Machop", Type: "Fighting", HP: 70, MaxHP: 70},
	"Giovanni": {Name: "Snorlax", Type: "Normal", HP: 160, MaxHP: 160},
	"Lance":    {Name: "Charmander", Type: "Fire", HP: 39, MaxHP: 39},
	"Cynthia":  {Name: "Bulbasaur", Type: "Grass", HP: 45, MaxHP: 45},
}

// RegionPokemon maps region names to Pokemon found in that region.
var RegionPokemon = map[string][]Pokemon{
	"Kanto": {
		{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35},
		{Name: "Charmander", Type: "Fire", HP: 39, MaxHP: 39},
		{Name: "Squirtle", Type: "Water", HP: 44, MaxHP: 44},
		{Name: "Bulbasaur", Type: "Grass", HP: 45, MaxHP: 45},
	},
	"Johto": {
		{Name: "Jigglypuff", Type: "Normal", HP: 115, MaxHP: 115},
		{Name: "Eevee", Type: "Normal", HP: 55, MaxHP: 55},
		{Name: "Geodude", Type: "Rock", HP: 40, MaxHP: 40},
	},
	"Hoenn": {
		{Name: "Snorlax", Type: "Normal", HP: 160, MaxHP: 160},
		{Name: "Machop", Type: "Fighting", HP: 70, MaxHP: 70},
		{Name: "Gastly", Type: "Ghost", HP: 30, MaxHP: 30},
	},
}
