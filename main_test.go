package main

import "fmt"
import "net/http"
import "net/http/httptest"
import "encoding/json"
import "testing"
import "io/ioutil"

import "api/card-deck-api/router"
import "api/card-deck-api/models"


// setup for reusabled deck creations

func __test_create_deck (
	t *testing.T,
	ts *httptest.Server,
	query_params string,
	expected_card_count int,
	) {
	// Make a request to /create-deck?cards=AS,KD,AC,2C,KH

	create_deck_resp, err := http.Post(fmt.Sprintf("%s/create-deck%s", ts.URL, query_params), "application/json", nil)

	if err != nil {
	  t.Fatalf("Expected no error, got %v", err)
	}

	if create_deck_resp.StatusCode != 200 {
	  t.Fatalf("Expected status code 200, got %v", create_deck_resp.StatusCode)
	}

	body, err := ioutil.ReadAll(create_deck_resp.Body)
	var deck *models.Deck
	deck_err := json.Unmarshal(body, &deck)

	if deck_err != nil {
	  panic(deck_err)
	}

	// Expected 5 five cards in deck

	if deck.Remaining != expected_card_count {
		t.Fatalf("Expected %d cards in deck, got %v", expected_card_count, deck.Remaining)
	}
}

// test `/create-deck` endpoint

func TestCreateDeck (t *testing.T) {
  // The SetupRouter method, that we previously refactored
  // is injected into a test server

  ts := httptest.NewServer(router.SetupRouter())

  // Shut down the server and block until all requests have gone through

  defer ts.Close()

  __test_create_deck(t, ts, "", 52)
  __test_create_deck(t, ts, "?cards=AS,KD,AC,2C,KH", 5)
}

// test `/open-deck` endpoint

func TestOpenDeck (t *testing.T) {
  // The SetupRouter method, that we previously refactored
  // is injected into a test server

  ts := httptest.NewServer(router.SetupRouter())

  // Shut down the server and block until all requests have gone through

  defer ts.Close()

  // Make a request to /create-deck

  create_deck_resp, err := http.Post(fmt.Sprintf("%s/create-deck", ts.URL), "application/json", nil)

  if err != nil {
    t.Fatalf("Expected no error, got %v", err)
  }

  if create_deck_resp.StatusCode != 200 {
    t.Fatalf("Expected status code 200, got %v", create_deck_resp.StatusCode)
  }

  body, err := ioutil.ReadAll(create_deck_resp.Body)
  var deck *models.Deck
  deck_err := json.Unmarshal(body, &deck)

  if deck_err != nil {
    panic(deck_err)
  }

  deck_id := deck.Deck_id

  open_deck_response, od_err := http.Get(fmt.Sprintf("%s/open-deck/%s", ts.URL, deck_id))

  if od_err != nil {
    t.Fatalf("Expected no error, got %v", od_err)
  }

  body, err = ioutil.ReadAll(open_deck_response.Body)
  deck_err = json.Unmarshal(body, &deck)

  if deck_err != nil {
    panic(deck_err)
  }

  // Expected 52 cards in deck

  if deck.Remaining != 52 {
  	t.Fatalf("Expected 52 cards in deck, got %v", deck.Remaining)
  }
}

// test card draw counts

func __test_card_draw_count (
	t *testing.T,
	ts *httptest.Server,
	deck_id string,
	draw_cards_count int,
	) {
	draw_card_response, od_err := http.Get(fmt.Sprintf("%s/draw-card?deck_id=%s&count=%d", ts.URL, deck_id, draw_cards_count))

	if od_err != nil {
	  t.Fatalf("Expected no error, got %v", od_err)
	}

	draw_card_body, draw_card_err := ioutil.ReadAll(draw_card_response.Body)

	if draw_card_err != nil {
	  panic(draw_card_err)
	}

	var drawn_cards *models.ProxyDeck
	card_draw_err := json.Unmarshal(draw_card_body, &drawn_cards)

	if card_draw_err != nil {
	  panic(card_draw_err)
	}

	// Expected `draw_cards_count` cards in deck

	if len(drawn_cards.Cards) != draw_cards_count {
		t.Fatalf("Expected %d drawn_cards in deck, got %v", draw_cards_count, len(drawn_cards.Cards))
	}
}

// test if card count requested exceeds cards in deck

func __test_card_draw_count_exceeded (
	t *testing.T,
	ts *httptest.Server,
	deck_id string,
	draw_cards_count int,
	) {
	draw_card_response, od_err := http.Get(fmt.Sprintf("%s/draw-card?deck_id=%s&count=%d", ts.URL, deck_id, draw_cards_count))

	if od_err != nil {
	  t.Fatalf("Expected no error, got %v", od_err)
	}

	if draw_card_response.StatusCode != 400 {
		t.Fatalf("Expected 403 error -> bad request, got %v", draw_card_response.StatusCode)
	}
}

// test for non-exisiting decks

func __test_card_draw_404 (
	t *testing.T,
	ts *httptest.Server,
	) {
	draw_card_response, od_err := http.Get(fmt.Sprintf("%s/draw-card?deck_id=%s&count=%d", ts.URL, "random-uuid", 10))

	if od_err != nil {
	  t.Fatalf("Expected no error, got %v", od_err)
	}

	if draw_card_response.StatusCode != 404 {
		t.Fatalf("Expected 404 error -> not found, got %v", draw_card_response.StatusCode)
	}
}

// test for empty parameters

func __test_card_draw_params (
	t *testing.T,
	ts *httptest.Server,
	) {
	draw_card_response, od_err := http.Get(fmt.Sprintf("%s/draw-card", ts.URL))

	if od_err != nil {
	  t.Fatalf("Expected no error, got %v", od_err)
	}

	if draw_card_response.StatusCode != 400 {
		t.Fatalf("Expected 400 error -> bad request, got %v", draw_card_response.StatusCode)
	}
}

// test `/draw-card` endpoint

func TestDrawCard (t *testing.T) {
  // The SetupRouter method, that we previously refactored
  // is injected into a test server

  ts := httptest.NewServer(router.SetupRouter())

  // Shut down the server and block until all requests have gone through

  defer ts.Close()

  // Make a request to /create-deck

  create_deck_resp, err := http.Post(fmt.Sprintf("%s/create-deck", ts.URL), "application/json", nil)

  if err != nil {
    t.Fatalf("Expected no error, got %v", err)
  }

  if create_deck_resp.StatusCode != 200 {
    t.Fatalf("Expected status code 200, got %v", create_deck_resp.StatusCode)
  }

  body, err := ioutil.ReadAll(create_deck_resp.Body)
  var deck *models.Deck
  deck_err := json.Unmarshal(body, &deck)

  if deck_err != nil {
    panic(deck_err)
  }

  deck_id := deck.Deck_id
  draw_cards_count := 15

	__test_card_draw_params(t, ts)
	__test_card_draw_404(t, ts)
	__test_card_draw_count(t, ts, deck_id, draw_cards_count)
	__test_card_draw_count_exceeded(t, ts, deck_id, 60)
}
