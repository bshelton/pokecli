# Pokecli

Pokecli is a cli application that can retreive cards utilizing the https://api.pokemontcg.io/v2 API.

Building
==========

```bash
$ make build        # Build for current arch
$ make compile      # Compile for Darwin, Linux and Windows.
```

Examples
===============

Pokecli supports two modes.
- Flex Mode (Allows for flags for cardLimit, sortBy and fields)
- Classic Mode (Searches for 10 cards of Type "Fire" or "Grass" with at least 90 HP and a rarity of "Rare" sorted by id.)


```bash
pokecli --help
Usage of pokecli:
  -classic
        In classic mode the cli returns exactly what was asked for.
  -fields string
        A comma seperated list of fields to include in the response. If empty it will include all fields
  -limit int
        The number of cards to return (default 250)
  -loglevel int
        Debug - 0
        Info - 1
        Warning - 2
        Error -3
        Fatal -4 (default 4)
  -sortBy string
        The field to sort by (default "id")
```

```bash
# Run in classic mode
$ pokecli -classic
{
  "Cards": [
    {
      "ID": "base1-11",
      "Name": "Nidoking",
      "Type": "Grass",
      "HP": "90",
      "Rarity": "Rare Holo"
    },
    {
      "ID": "base1-15",
      "Name": "Venusaur",
      "Type": "Grass",
      "HP": "100",
      "Rarity": "Rare Holo"
    },
    {
      "ID": "base1-4",
      "Name": "Charizard",
      "Type": "Fire",
      "HP": "120",
      "Rarity": "Rare Holo"
    },
    {
      "ID": "base2-23",
      "Name": "Nidoqueen",
      "Type": "Grass",
      "HP": "90",
      "Rarity": "Rare"
    },
    {
      "ID": "base2-7",
      "Name": "Nidoqueen",
      "Type": "Grass",
      "HP": "90",
      "Rarity": "Rare Holo"
    },
    {
      "ID": "base4-11",
      "Name": "Nidoking",
      "Type": "Grass",
      "HP": "90",
      "Rarity": "Rare Holo"
    },
    {
      "ID": "base4-12",
      "Name": "Nidoqueen",
      "Type": "Grass",
      "HP": "90",
      "Rarity": "Rare Holo"
    },
    {
      "ID": "base4-18",
      "Name": "Venusaur",
      "Type": "Grass",
      "HP": "100",
      "Rarity": "Rare Holo"
    },
    {
      "ID": "base4-4",
      "Name": "Charizard",
      "Type": "Fire",
      "HP": "120",
      "Rarity": "Rare Holo"
    },
    {
      "ID": "base6-18",
      "Name": "Venusaur",
      "Type": "Grass",
      "HP": "100",
      "Rarity": "Rare Holo"
    }
  ]
}

# Return Only 1 card in flex mode.
$ pokecli -limit 1 
{
  "Cards": [
    {
      "hp": "90",
      "id": "base1-11",
      "name": "Nidoking",
      "rarity": "Rare Holo",
      "types": [
        "Grass"
      ]
    }
  ]
}
```

```bash
# Find 5 cards sorted by name and only display id,name and hp
$ pokecli -sortBy name -limit 5 -fields id,name,hp
{
  "Cards": [
    {
      "hp": "100",
      "id": "dp2-19",
      "name": "Abomasnow"
    },
    {
      "hp": "130",
      "id": "sm6-4",
      "name": "Abomasnow"
    },
    {
      "hp": "140",
      "id": "swsh10tg-TG01",
      "name": "Abomasnow"
    },
    {
      "hp": "140",
      "id": "swsh6-10",
      "name": "Abomasnow"
    },
    {
      "hp": "140",
      "id": "swsh2-13",
      "name": "Abomasnow"
    }
  ]
}

```


An ApiKey can also be used by setting the environment variable `POKE_API_KEY`
Logs regarding the pokecli will be logged to 

Requirements
===============

In order to develop and contribute to pokecli, you must install the
following items

* Go (1.19.5)


Contributing
=================
In order to contribute to pokecli, each merge request must be submitted on a
feature branch. The `main` branch should not be used to commmit.

https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/creating-and-deleting-branches-within-your-repository

The branch should use a short and descriptive name.