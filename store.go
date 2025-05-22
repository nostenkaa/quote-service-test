
package main

var quotes []Quote
var nextID = 1

func AddQuote(author, text string) Quote {
	q := Quote{
		ID:     nextID,
		Author: author,
		Quote:  text,
	}
	quotes = append(quotes, q)
	nextID++
	SaveQuotes()
	return q
}

func GetAllQuotes() []Quote {
	return quotes
}

func GetQuotesByAuthor(author string) []Quote {
	var filtered []Quote
	for _, q := range quotes {
		if q.Author == author {
			filtered = append(filtered, q)
		}
	}
	return filtered
}

func DeleteQuoteByID(id int) bool {
	for i, q := range quotes {
		if q.ID == id {
			quotes = append(quotes[:i], quotes[i+1:]...)
			SaveQuotes()
			return true
		}
	}
	return false
}
