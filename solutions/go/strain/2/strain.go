package strain

type Predicate[T any] func(x T) bool

// Filter iterates over the elements of a collection and returns a new slice
// containing only the elements that satisfy the given predicate function.
//
// If 'exclude' is false, it behaves like 'Keep' (elements where predicate is true are kept).
// If 'exclude' is true, it behaves like 'Discard' (elements where predicate is false are kept).
func Filter[T any](collection []T, predicate Predicate[T], exclude bool) []T {
    var result []T
    for _, elem := range collection {
        predicateCheck := predicate(elem)
        if ((predicateCheck && !exclude) || (!predicateCheck && exclude)) {
            result = append(result, elem)
        }
    }
    return result
}

// Keep iterates over the elements of a collection and returns a new slice
// containing only the elements for which the predicate function returns true.
// It is a shorthand for calling Filter with 'exclude' set to false.
func Keep[T any](collection []T, predicate Predicate[T]) []T {
    return Filter(collection, predicate, false)
}

// Discard iterates over the elements of a collection and returns a new slice
// containing only the elements for which the predicate function returns false.
// It is a shorthand for calling Filter with 'exclude' set to true.
func Discard[T any](collection []T, predicate Predicate[T]) []T {
    return Filter(collection, predicate, true)
}
