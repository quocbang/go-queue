package queue

type Queue struct {
	ProductTypes []string
}

func BuildQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(productType string) {
	q.ProductTypes = append(q.ProductTypes, productType)
}

func (q *Queue) Pop() string {
	if q.Empty() {
		return "product types are empty, please push more"
	}
	popedElement := q.ProductTypes[0]
	q.ProductTypes = q.ProductTypes[1:]
	return popedElement
}

func (q *Queue) Empty() bool {
	return len(q.ProductTypes) == 0
}
