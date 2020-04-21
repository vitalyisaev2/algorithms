// Package matching contains algorithm implementations from book
// "EXACT STRING MATCHING ALGORITHMS" by Christian Charras - Thierry Lecroq
// http://www-igm.univ-mlv.fr/~lecroq/string/index.html
package matching

type matcher func(pattern, text string) bool
