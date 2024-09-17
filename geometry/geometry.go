package geometry

type Rectangle struct { //ширина и высота прямоугольника
    Width  float64 
    Height float64
}

func (r Rectangle) Area() float64 { //определение площади
    return r.Width * r.Height
}

func (r Rectangle) Scale(factor float64) (float64, float64) { //Scale изменяет размеры прямоугольника
    r.Width *= factor
    r.Height *= factor
    return r.Width, r.Height
}