package poketypes

type Pokemon struct {
	Id                     int                  `json:"id"`
	Name                   string               `json:"name"`
	BaseExperience         int                  `json:"base_experience"`
	Height                 int                  `json:"height"`
	IsDefault              bool                 `json:"is_default"`
	Order                  int                  `json:"order"`
	Weight                 int                  `json:"weight"`
	Abilities              []PokemonAbility     `json:"abilities"`
	Forms                  []NamedAPIResource   `json:"forms"`
	GameIndices            []VersionGameIndex   `json:"game_indices"`
	HeldItems              []PokemonHeldItem    `json:"held_items"`
	LocationAreaEncounters string               `json:"location_area_encounters"`
	Moves                  []PokemonMove        `json:"moves"`
	PastTypes              []PokemonTypePast    `json:"past_types"`
	PastAbilities          []PokemonAbilityPast `json:"past_abilities"`
	PastStats              []PokemonStatPast    `json:"past_stats"`
	Sprites                PokemonSprites       `json:"sprites"`
	Cries                  PokemonCries         `json:"cries"`
	Species                NamedAPIResource     `json:"species"`
	Stats                  []PokemonStat        `json:"stats"`
	Types                  []PokemonType        `json:"types"`
}

type PokemonAbility struct {
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
	Ability  NamedAPIResource `json:"ability"`
}

type PokemonType struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type PokemonFormType struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type PokemonTypePast struct {
	Generation NamedAPIResource `json:"generation"`
	Types      []PokemonType    `json:"types"`
}

type PokemonAbilityPast struct {
	Generation NamedAPIResource `json:"generation"`
	Abilities  []PokemonAbility `json:"abilities"`
}

type PokemonStatPast struct {
	Generation NamedAPIResource `json:"generation"`
	Stats      []PokemonStat    `json:"stats"`
}

type PokemonHeldItem struct {
	Item           NamedAPIResource         `json:"item"`
	VersionDetails []PokemonHeldItemVersion `json:"version_details"`
}

type PokemonHeldItemVersion struct {
	Version NamedAPIResource `json:"version"`
	Rarity  int              `json:"rarity"`
}

type PokemonMove struct {
	Move                NamedAPIResource     `json:"move"`
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonMoveVersion struct {
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	VersionGroup    NamedAPIResource `json:"version_group"`
	LevelLearnedAt  int              `json:"level_learned_at"`
	Order           int              `json:"order"`
}

type PokemonStat struct {
	Stat     NamedAPIResource `json:"stat"`
	Effort   int              `json:"effort"`
	BaseStat int              `json:"base_stat"`
}

type PokemonSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
}

type PokemonCries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type VersionGameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}
type LocationAreaSet struct {
	Count    int                `json:"Count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type EncounterVersionDetails struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}
type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource          `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}
type Name struct {
	Name     string           `json:"name"`
	Language NamedAPIResource `json:"language"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type Encounter struct {
	MinLevel        int                `json:"min_level"`
	MaxLevel        int                `json:"max_level"`
	ConditionValues []NamedAPIResource `json:"condition_values"`
	Chance          int                `json:"chance"`
	Method          NamedAPIResource   `json:"method"`
}
type VersionEncounterDetail struct {
	Version          NamedAPIResource `json:"version"`
	MaxChance        int              `json:"max_chance"`
	EncounterDetails []Encounter      `json:"encounter_details"`
}
type PokemonEncounter struct {
	Pokemon        NamedAPIResource         `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}
type LocationArea struct {
	Id                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedAPIResource      `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}
