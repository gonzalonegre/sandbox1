package planet
type Weather int

const (
	Unknown Weather = iota
	Sequia                   // Alineados con el sol, esto es el mismo angulo.
	TemperaturaOptima        // Alineados sin el sol
	Lluvia                   // Centro triangulo
	LluviaIntensa            // Centro triangulo y perimetro mas largo
)

func (w Weather) String() string {
	switch w {
	case Unknown:
		return "Sin Datos"
	case Sequia:
		return "Sequia"
	case TemperaturaOptima:
		return "Condiciones optimas"
	case Lluvia:
		return "Lluvia"
	case LluviaIntensa:
		return "Lluvia Intensa"
	}
        return ""
}

func WeatherForDay(planet1, planet2, planet3 *Planet, day int) (Weather, float64) {
   weather:= Unknown
   var perimetro float64 = 0;
   p1:= planet1.CurrentPos(day)
   p2:= planet2.CurrentPos(day)
   p3:= planet3.CurrentPos(day)
   p4sol:= Point{0,0}
   if p1.IsInLine(p2,p3) {
   	if p4sol.IsInLine(p2,p3){
   		weather	= Sequia
   	}else{
   		weather	= TemperaturaOptima 
   	}
   }else{
  
   	if PointInsideTriangle(p4sol, p1, p2, p3) {
   		weather = Lluvia
		perimetro = p1.Distance(p2) + p2.Distance(p3) + p3.Distance(p1) 
   	}
   }
   return weather,perimetro
   
}


