package main

type Valider interface { // what a stupid name
	Valid() error
}
