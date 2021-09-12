package controllers

import "net/http"
import "strconv"

import "github.com/gin-gonic/gin"

import "api/card-deck-api/models"
import "api/card-deck-api/utils"


// store all card decks

var decks = []*models.Deck {}


// view to create deck of cards
// card codes (AS, KH ...) can be passed (optional) thought the `cards` query to create specific cards in deck

func CreateDeck (c *gin.Context) {
	// get cards query parameter

	cards_query, query_exists := c.GetQuery("cards")

	// create instances for card suits

	var clubs, diamonds, hearts, spades []models.Card

	// decide if query should be used or not
	// check if the `cards` parameter exists and if it has card codes (AS, KH ...)

	var use_query bool = query_exists && len(cards_query) > 0

	if (use_query) {
		// using `cards` to created decks

		clubs = utils.CreateSuit("clubs", cards_query)
		diamonds = append(clubs, utils.CreateSuit("diamonds", cards_query)...)
		hearts = append(diamonds, utils.CreateSuit("hearts", cards_query)...)
		spades = append(hearts, utils.CreateSuit("spades", cards_query)...)
	} else {
		// create deck with generic pattern of 52 cards

		clubs = utils.CreateSuit("clubs")
		diamonds = append(clubs, utils.CreateSuit("diamonds")...)
		hearts = append(diamonds, utils.CreateSuit("hearts")...)
		spades = append(hearts, utils.CreateSuit("spades")...)
	}

	// collect all card suits

	cards := spades

	// create instance of a new `*models.Deck`
	// - `Deck_id` -> string -> unique UUID of each deck
	// - `Cards` -> [] Cards -> all cards available in deck
	// - `shuffled` -> bool -> if the deck has been shuffled or not
	// - `Remaining` -> int -> length of cards available in deck

	// generate new UUID for `deck_id`

	var new_deck_id string = utils.GenerateUniqueUUID()

	// deck to be added to `decks`

	var __new_deck models.Deck = models.Deck {
		Deck_id: new_deck_id,
		Cards: cards,
		Shuffled: false,
		Remaining: len(cards),
	}

	// add `__new_deck` to `decks`

	decks = append(decks, &__new_deck)

	// `decks` to be returned in response

	var new_deck models.Deck = models.Deck {
		Deck_id: new_deck_id,
		Shuffled: false,
		Remaining: len(cards),
	}

	// return `new_deck` as reponse to endpoint

	c.IndentedJSON(http.StatusOK, new_deck)
}


// view to fetch / open a specified deck using the `deck_id` -> UUID

func OpenDeck (c *gin.Context) {
	// get deck_id

	deck_id := c.Param("deck_id")

	// linear search for deck using `deck_id`

	for _, deck := range decks {
    if deck.Deck_id == deck_id {
      c.IndentedJSON(http.StatusOK, deck)
      return
    }
  }

  // raise `StatusNotFound` -> `not_found` if deck_id is not found

  c.IndentedJSON(http.StatusNotFound, gin.H {"not_found": "*models.Deck not found"})
}


// draw cards from deck using;
// -`deck_id` (required)
// -`count` (required)

func DrawCard (c *gin.Context) {
	// get query parameters -> `deck_id` & `count`

	deck_id, deck_id_query_exists := c.GetQuery("deck_id")
	__count, count_query_exists := c.GetQuery("count")
	count, _ := strconv.Atoi(__count)

	// check if the required parameters are passed

	is_valid := deck_id_query_exists && count_query_exists

	// raise `StatusBadRequest` -> `bad_request` if required query parameters are not passed

	if (!is_valid) {
		c.IndentedJSON(http.StatusBadRequest, gin.H {"bad_request": "Deck ID and Card count is required"})
		return
	}

	// store cards with length of query parameter value of `count`

	var cards = []models.Card {}

	// linear search for deck using `deck_id`

	for _, deck := range decks {
    if deck.Deck_id == deck_id {

    	// check if `count` exceeds the length of cards in deck
    	// - raise `StatusBadRequest` -> 'bad_request' if count limit is exceeded

    	var count_exceeds_deck_cards bool = count > len(deck.Cards)
    	if (count_exceeds_deck_cards) {
    		c.IndentedJSON(http.StatusBadRequest, gin.H {"bad_request": "Card count exceeded"})
    		return
    	}

    	// get cards of length of query parameter value of `count`

    	for card_index := 0; card_index < count; card_index++ {

    		// update (card -> *models.Card) to (cards []*models.Card)

    		cards = append(cards, deck.Cards[card_index])
    	}

    	// create instance of `*models.ProxyDeck` model using array -> `cards`

    	cards_proxy := models.ProxyDeck {
    		Cards: cards,
    	}

    	// return `cards` as `cards_proxy` model

    	c.IndentedJSON(http.StatusOK, cards_proxy)
    	return
    }
  }

  // raise `StatusNotFound` -> `not_found` if deck does not exist

  c.IndentedJSON(http.StatusNotFound, gin.H {"not_found": "*models.Deck not found"})
}
