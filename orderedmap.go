package orderedmap

import "sync"

// OrderdMap is a map that maintains the order of keys.
type OrderdMap[TK comparable, TV any] struct {
	Map      map[TK]TV    // Map stores the key-value pairs.
	Keys     []TK         // Keys stores the order of keys.
	withlock bool         // withlock is a flag to enable or disable the mutex.
	mu       sync.RWMutex // Mutex to handle concurrent access.
}

// Element represents a key-value pair.
type OrderdMapElement[TK comparable, TV any] struct {
	Key   TK // Key is the key of the element.
	Value TV // Value is the value of the element.
}

// NewOrderdMap creates a new OrderdMap from a slice of Elements.
// It initializes the OrderdMap, and populates them with the provided elements.
// NewOrderdMap([]Element[TK, TV]{})
func NewOrderdMap[TK comparable, TV any](es []OrderdMapElement[TK, TV], withlock bool) *OrderdMap[TK, TV] {
	o := &OrderdMap[TK, TV]{
		Map:      make(map[TK]TV),
		Keys:     make([]TK, 0),
		withlock: withlock,
	}
	for _, v := range es {
		o.Map[v.Key] = v.Value
		o.Keys = append(o.Keys, v.Key)
	}
	return o
}

// Range returns a function that iterates over the elements in the OrderdMap.
// The provided yield function is called for each element with its index and value.
// If the yield function returns false, the iteration stops.
func (o *OrderdMap[TK, TV]) Range() func(func(TK, TV) bool) {
	if o.withlock {
		o.mu.RLock()
		defer o.mu.RUnlock()
	}
	return func(yield func(TK, TV) bool) {
		for _, mk := range o.Keys {
			if !yield(mk, o.Map[mk]) {
				return
			}
		}
	}
}

// Get returns the value associated with the given key.
func (o *OrderdMap[TK, TV]) Get(key TK) TV {
	if o.withlock {
		o.mu.RLock()
		defer o.mu.RUnlock()
	}
	return o.Map[key]
}

// Len returns the number of elements in the OrderdMap.
func (o *OrderdMap[TK, TV]) Len() int {
	if o.withlock {
		o.mu.RLock()
		defer o.mu.RUnlock()
	}
	return len(o.Keys)
}

// Set adds or updates the value associated with the given key.
// If the key does not exist, it is added to the keys slice.
func (o *OrderdMap[TK, TV]) Set(key TK, value TV) {
	if o.withlock {
		o.mu.RLock()
		defer o.mu.RUnlock()
	}
	if _, ok := o.Map[key]; !ok {
		o.Keys = append(o.Keys, key)
	}
	o.Map[key] = value
}
