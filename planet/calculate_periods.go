package planet
import "github.com/gonzalonegre/sandbox1/compare"
import "fmt"
// 	Unknown Weather = iota
// 	Sequia                   // Alineados con el sol, esto es el mismo angulo.
// 	TemperaturaOptima        // Alineados sin el sol
// 	Lluvia                   // Centro triangulo
// 	LluviaIntensa            // Centro triangulo y perimetro mas largo


func CalculatePeriods (totaldays uint) {

  // Aca se pueden hacer varias optimizaciones porque todas las posiciones se repiten cada 360 dias.
  // Es decir puedo calcular para la cantidad de anios enteras y luego el resto
  // Esto es porque al fin y al principio del anio no se repite el weather si empezas en 0, sino 
  // tambien se podria tenes eso en cuenta al multiplicar.
  // o mejor, suponer que son 360 dias por anio
  // solo pueden coincidir los angulos cada 15 grados pasando por el centro
  

   P1 := Planet{ 500, 0, -1}
   P2 := Planet{2000, 0, -3}
   P3 := Planet{1000, 0,  5}

   
   var cantidadPeriodos = make(map[Weather]int)

   var perimetroMasLargo float64 = 0
   var diaPerimetroMasLargo  = 0
   var cantidadDiasPerimetroMasLargo  = 0

   var lastPerimetro float64 = 0
   var previousWeather = Unknown;
   var currentWeather = Unknown;
   for i := 0 ; i <= int(totaldays); i++ {
	// cuanto los cambios a comenzar el periodo
	// hago un ciclo de mas para contar el ultimo solamente.

	if(previousWeather!=currentWeather){
		 cantidadPeriodos[currentWeather] =  cantidadPeriodos[currentWeather] + 1
	 }

	if previousWeather==Lluvia {
		if compare.FloatsEquals(lastPerimetro,perimetroMasLargo) {
			cantidadDiasPerimetroMasLargo++
		}else{
			if lastPerimetro>perimetroMasLargo {
				 perimetroMasLargo = lastPerimetro
				 diaPerimetroMasLargo = i - 1
				 cantidadDiasPerimetroMasLargo = 1
			 }
		}
	 }
	// Calculo el siguiente dia
	previousWeather = currentWeather
	currentWeather, lastPerimetro = WeatherForDay(&P1, &P2, &P3, i)
   }
   
   fmt.Println(cantidadPeriodos)
   fmt.Println("Habra" , cantidadDiasPerimetroMasLargo , "dias de lluvia intensa, por ejemplo el dia", diaPerimetroMasLargo, "Perimetro", perimetroMasLargo, "kms" )
   
}


