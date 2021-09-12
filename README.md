# card-deck-api

Card Deck for card games like Poker and Blackjack. This API would be able to handle the deck and cards to be used in any
game like these. There are 3 API endpoints to manage cards

- Create Deck -> /create-deck
- Open Deck -> /open-deck/:deck-id
- Draw Card -> /draw-card?deck_id=<deck_id>&count=<count>

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
