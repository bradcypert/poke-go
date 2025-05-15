package v2

type Cries struct {
	Latest string `json:"latest,omitempty"`
	Legacy string `json:"legacy,omitempty"`
}

type ReferenceItem struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type GameIndex struct {
	GameIndex int           `json:"game_index,omitempty"`
	Version   ReferenceItem `json:"version,omitempty"`
}

type Sprites struct {
	BackDefault           string `json:"back_default,omitempty"`
	BackFemale            string `json:"back_female,omitempty"`
	BackShiny             string `json:"back_shiny,omitempty"`
	BackShinyFemale       string `json:"back_shiny_female,omitempty"`
	FrontDefault          string `json:"front_default,omitempty"`
	FrontFemale           string `json:"front_female,omitempty"`
	FrontShiny            string `json:"front_shiny,omitempty"`
	FrontShinyFemale      string `json:"front_shiny_female,omitempty"`
	BackShinyTransparent  string `json:"back_shiny_transparent,omitempty"`
	BackTransparent       string `json:"back_transparent,omitempty"`
	FrontShinyTransparent string `json:"front_shiny_transparent,omitempty"`
	FrontTransparent      string `json:"front_transparent,omitempty"`
}

type ResultSet struct {
	Count    int             `json:"count,omitempty"`
	Next     string          `json:"next,omitempty"`
	Previous string          `json:"previous,omitempty"`
	Results  []ReferenceItem `json:"results,omitempty"`
}

type Pokemon struct {
	Abilities []struct {
		Ability  ReferenceItem `json:"ability,omitempty"`
		IsHidden bool          `json:"is_hidden,omitempty"`
		Slot     int           `json:"slot,omitempty"`
	} `json:"abilities,omitempty"`
	BaseExperience int             `json:"base_experience,omitempty"`
	Cries          Cries           `json:"cries,omitempty"`
	Forms          []ReferenceItem `json:"forms,omitempty"`
	GameIndices    []GameIndex     `json:"game_indices,omitempty"`
	Height         int             `json:"height,omitempty"`
	HeldItems      []struct {
		Item           ReferenceItem `json:"item,omitempty"`
		VersionDetails []struct {
			Rarity  int           `json:"rarity,omitempty"`
			Version ReferenceItem `json:"version,omitempty"`
		} `json:"version_details,omitempty"`
	} `json:"held_items,omitempty"`
	ID                     int    `json:"id,omitempty"`
	IsDefault              bool   `json:"is_default,omitempty"`
	LocationAreaEncounters string `json:"location_area_encounters,omitempty"`
	Moves                  []struct {
		Move                ReferenceItem `json:"move,omitempty"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int           `json:"level_learned_at,omitempty"`
			MoveLearnMethod ReferenceItem `json:"move_learn_method,omitempty"`
			Order           int           `json:"order,omitempty"`
			VersionGroup    ReferenceItem `json:"version_group,omitempty"`
		} `json:"version_group_details,omitempty"`
	} `json:"moves,omitempty"`
	Name          string `json:"name,omitempty"`
	Order         int    `json:"order,omitempty"`
	PastAbilities []struct {
		Abilities []struct {
			Ability  *ReferenceItem `json:"ability,omitempty"`
			IsHidden bool           `json:"is_hidden,omitempty"`
			Slot     int            `json:"slot,omitempty"`
		} `json:"abilities,omitempty"`
		Generation ReferenceItem `json:"generation,omitempty"`
	} `json:"past_abilities,omitempty"`
	PastTypes []struct {
		Generation ReferenceItem `json:"generation,omitempty"`
		Types      []struct {
			Slot int           `json:"slot,omitempty"`
			Type ReferenceItem `json:"type,omitempty"`
		} `json:"types,omitempty"`
	} `json:"past_types,omitempty"`
	Species ReferenceItem `json:"species,omitempty"`
	Sprites struct {
		Sprites
		Other struct {
			DreamWorld      Sprites `json:"dream_world,omitempty"`
			Home            Sprites `json:"home,omitempty"`
			OfficialArtwork Sprites `json:"official-artwork,omitempty"`
			Showdown        Sprites `json:"showdown,omitempty"`
		} `json:"other,omitempty"`
		Versions struct {
			GenerationI struct {
				RedBlue Sprites `json:"red-blue,omitempty"`
				Yellow  Sprites `json:"yellow,omitempty"`
			} `json:"generation-i,omitempty"`
			GenerationIi struct {
				Crystal Sprites `json:"crystal,omitempty"`
				Gold    Sprites `json:"gold,omitempty"`
				Silver  Sprites `json:"silver,omitempty"`
			} `json:"generation-ii,omitempty"`
			GenerationIii struct {
				Emerald          Sprites `json:"emerald,omitempty"`
				FireredLeafgreen Sprites `json:"firered-leafgreen,omitempty"`
				RubySapphire     Sprites `json:"ruby-sapphire,omitempty"`
			} `json:"generation-iii,omitempty"`
			GenerationIv struct {
				DiamondPearl        Sprites `json:"diamond-pearl,omitempty"`
				HeartgoldSoulsilver Sprites `json:"heartgold-soulsilver,omitempty"`
				Platinum            Sprites `json:"platinum,omitempty"`
			} `json:"generation-iv,omitempty"`
			GenerationV struct {
				BlackWhite struct {
					Sprites
					Animated Sprites `json:"animated,omitempty"`
				} `json:"black-white,omitempty"`
			} `json:"generation-v,omitempty"`
			GenerationVi struct {
				OmegarubyAlphasapphire Sprites `json:"omegaruby-alphasapphire,omitempty"`
				XY                     Sprites `json:"x-y,omitempty"`
			} `json:"generation-vi,omitempty"`
			GenerationVii struct {
				Icons             Sprites `json:"icons,omitempty"`
				UltraSunUltraMoon Sprites `json:"ultra-sun-ultra-moon,omitempty"`
			} `json:"generation-vii,omitempty"`
			GenerationViii struct {
				Icons Sprites `json:"icons,omitempty"`
			} `json:"generation-viii,omitempty"`
		} `json:"versions,omitempty"`
	} `json:"sprites,omitempty"`
	Stats []struct {
		BaseStat int           `json:"base_stat,omitempty"`
		Effort   int           `json:"effort,omitempty"`
		Stat     ReferenceItem `json:"stat,omitempty"`
	} `json:"stats,omitempty"`
	Types []struct {
		Slot int           `json:"slot,omitempty"`
		Type ReferenceItem `json:"type,omitempty"`
	} `json:"types,omitempty"`
	Weight int `json:"weight,omitempty"`
}

type Generation struct {
	Abilities  []ReferenceItem `json:"abilities,omitempty"`
	ID         int             `json:"id,omitempty"`
	MainRegion ReferenceItem   `json:"main_region,omitempty"`
	Moves      []ReferenceItem `json:"moves,omitempty"`
	Name       string          `json:"name,omitempty"`
	Names      []struct {
		Language ReferenceItem `json:"language,omitempty"`
		Name     string        `json:"name,omitempty"`
	} `json:"names,omitempty"`
	PokemonSpecies []ReferenceItem `json:"pokemon_species,omitempty"`
	Types          []ReferenceItem `json:"types,omitempty"`
	VersionGroups  []ReferenceItem `json:"version_groups,omitempty"`
}
