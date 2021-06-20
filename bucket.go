package gostat

import "fmt"

// bucket keep track of segmented data
type bucket struct {
	// used later for statistics
	sum, sum2 float64
	// center and width of bucket. Bucket contains from c - w/2 to c+ w/2, both ends INCLUDED.
	// There is always at leaset a data point on each end, unless the bucket is empty.
	c, w float64
	// nb of data points in bucket
	n int
}

func (b bucket) Header() string {
	return "[ From\t\tCenter\t\tTO\t  ]  =>\t   n\tdensity\t\twidth\t\tsurface(n*w)\t"
}

func (b *bucket) String() string {
	return fmt.Sprintf("[%10.3f\t%10.3f\t%10.3f]  =>  %d\t%10.3f\t%10.3f\t%10.1f", b.low(), b.c, b.high(), b.n, b.density(), b.w, float64(b.n)*b.w)
}

func (b *bucket) add(data float64) {
	b.sum += data
	b.sum2 += data * data
	b.n += 1
}

// low is lower limit
func (b bucket) low() float64 {
	return b.c - b.w/2.
}

// high is higher limit
func (b bucket) high() float64 {
	return b.c + b.w/2.
}

func (b bucket) density() float64 {
	if b.w == 0. {
		return 0.
	}
	return float64(b.n) / b.w

}

// test if data can fit in this bucket ?
func (b bucket) contains(d float64) bool {
	return (d >= b.low() && d <= b.high())
}

// =======================================

// buckets are sorted based on their center.
type buckets []bucket

func (bb buckets) Len() int           { return len(bb) }
func (bb buckets) Swap(i, j int)      { bb[i], bb[j] = bb[j], bb[i] }
func (bb buckets) Less(i, j int) bool { return bb[i].c < bb[j].c }

// eval evaluate the cost of merging i with i+1.
// Cost should be minimum.
// TODO allow for different eval startégies ...
func (bb buckets) eval(i int) float64 {
	w1, w2 := bb[i].w, bb[i+1].w
	w := bb[i+1].high() - bb[i].low()

	n1, n2 := float64(bb[i].n), float64(bb[i+1].n)
	n := n1 + n2

	return n*n*w*w - n1*n1*w1*w1 - n2*n2*w2*w2
}

// merge bucket i with i+1.
func (bb buckets) merge(i int) buckets {
	w := bb[i+1].high() - bb[i].low()
	c := (bb[i+1].high() + bb[i].low()) / 2.
	bb[i].c = c
	bb[i].w = w
	bb[i].sum += bb[i+1].sum
	bb[i].sum2 += bb[i+1].sum2
	bb[i].n += bb[i+1].n
	if i+2 < len(bb) {
		return append(bb[:i+1], bb[i+2:]...)
	} else {
		return bb[:i+1]
	}
}
