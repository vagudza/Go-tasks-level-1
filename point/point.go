package point

/*
вспомогательный пакет для задчи 24
*/

type Point struct {
	// неэкспортируемые идентификаторы
	x, y float64
}

func (p *Point) Set(x, y float64) {
	p.x = x
	p.y = y
}

// доступ к неэксп. полям через Getters/Setters
func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}
