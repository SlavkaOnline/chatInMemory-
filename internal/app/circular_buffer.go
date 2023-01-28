package app

import "golang.org/x/exp/constraints"

type CircleBuffer[TEntity any] struct {
	position int
	len      int
	capacity int
	values   []TEntity
}

func NewCircleBuffer[TEntity any](capacity int) *CircleBuffer[TEntity] {
	return &CircleBuffer[TEntity]{
		position: 0,
		len:      0,
		capacity: capacity,
		values:   make([]TEntity, capacity),
	}
}
func (buffer *CircleBuffer[TEntity]) Push(item TEntity) {
	if buffer.len < buffer.capacity {
		buffer.values = append(buffer.values, item)
		buffer.len = min(buffer.capacity, buffer.len+1)
		buffer.position++
	} else {
		if buffer.position == buffer.capacity {
			buffer.position = 0
		}

		buffer.values[buffer.position] = item
		buffer.position++
	}
}

func (buffer *CircleBuffer[TEntity]) Iter() *CircleBufferIter[TEntity] {
	return &CircleBufferIter[TEntity]{
		position: buffer.position,
		end:      buffer.len,
		values:   buffer.values,
		circled:  false,
	}
}

type CircleBufferIter[TEntity any] struct {
	index    int
	position int
	values   []TEntity
	end      int
	circled  bool
}

func (iter *CircleBufferIter[TEntity]) Next() (TEntity, bool) {
	if iter.index == iter.end && !iter.circled {
		iter.circled = true
		iter.index = 1
		return iter.values[0], true
	}

	if iter.index == iter.position && iter.circled {
		return nil, false
	}

	current := iter.values[iter.index]
	iter.index++
	return current, true
}

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
