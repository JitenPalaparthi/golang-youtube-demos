package rectangle

func (r *Rect) Perimeter() float32 {
	r.p = 2 * (r.L + r.B)
	return r.p
}
