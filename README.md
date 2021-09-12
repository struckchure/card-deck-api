# Card Deck API

Card Deck for card games like Poker and Blackjack. This API would be able to handle the deck and cards to be used in any
game like these. 

## To run the server

```
go run .
```
The server is started on port `8080` by default

## Running automated tests
```
go test -v
```
This would run tests on all endpoints listed below.


## API endpoints to manage cards

- Create Deck -> `/create-deck`
- Open Deck -> `/open-deck/<deck_id>`
- Draw Card -> `/draw-card?deck_id=<deck_id>&count=<count>`

## Create Deck
> It would create the standard 52-card deck of French playing cards, It includes
all thirteen ranks in each of the four suits: clubs ( ♣ ), diamonds ( ♦ ), hearts (♥)
and spades ( ♠ ).

Send a `POST` request to `/create-deck`
a new deck would be created with a return response like this

```json
{
  "deck_id": "104caed9-5de3-46af-5703-aba2407b51b6",
  "shuffled": false,
  "remaining": 52
}
```

You can also add queries to specify what kind of cards to create
Send a `POST` request to `/create-deck?cards=AS,KD,AC,2C,KH`

```json
{
  "deck_id": "104caed9-5de3-46af-5703-aba2407b51b6",
  "shuffled": false,
  "remaining": 5
}
```

## Open Deck
> It would return a given deck by its UUID. If the deck was not passed over or is
invalid, response will return an error. This method will "open the deck", meaning that
it will list all cards by the order it was created.

Send a `GET` request to `/open-deck/:deck_id` -> `/open-deck/104caed9-5de3-46af-5703-aba2407b51b6`

```json
{
  "deck_id": "104caed9-5de3-46af-5703-aba2407b51b6",
  "shuffled": false,
  "remaining": 52,
  "cards": [
    {
      "value": "KING",
      "suit": "CLUBS",
      "code": "KC"
    },
    {
      "value": "QUEEN",
      "suit": "CLUBS",
      "code": "QC"
    },
    {
      "value": "JOKER",
      "suit": "CLUBS",
      "code": "JC"
    },
    {
      "value": "10",
      "suit": "CLUBS",
      "code": "10C"
    },
    {
      "value": "9",
      "suit": "CLUBS",
      "code": "9C"
    },
    {
      ...
      ...
    },
    {
      "value": "3",
      "suit": "SPADES",
      "code": "3S"
    },
    {
      "value": "2",
      "suit": "SPADES",
      "code": "2S"
    },
    {
      "value": "ACE",
      "suit": "SPADES",
      "code": "AS"
    }
  ]
}
```

## Draw Card
> To draw a card(s) of a given Deck. If the deck was not passed over or
invalid it should return an error. A count parameter needs to be provided to
define how many cards to draw from the deck.

Send a `GET` request to `/draw-card` and pass query parameters `?deck_id=104caed9-5de3-46af-5703-aba2407b51b6&count=2`
Response should look something like this, the cards are returned in the order they were created.
  
```json
{
  "cards": [
      {
        "value": "KING",
        "suit": "CLUBS",
        "code": "KC"
      },
      {
        "value": "QUEEN",
        "suit": "CLUBS",
        "code": "QC"
      }
  ]
}
```
