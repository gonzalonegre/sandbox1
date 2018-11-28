package planet_test

import "testing"

import "github.com/gonzalonegre/sandbox1/planet";

func TestWeatherForDay(t *testing.T) {
{
	P1 := planet.Planet{3,  90 ,-10 }
	P2 := planet.Planet{4,  90 , 10 }
	P3 := planet.Planet{5, 270 ,  0 }
	weather, _:= planet.WeatherForDay( &P1,&P2,&P3,45)
	if weather!=planet.Lluvia {
		t.Fatalf("WeatherForDay( *P1,*P2,*P3,1)!=Lluvia failed.")
	}
}
{
	P1 := planet.Planet{3, 90 ,-1 }
	P2 := planet.Planet{4, 90 , 1 }
	P3 := planet.Planet{5, 90 , 0 }
	weather, _:= planet.WeatherForDay( &P1,&P2,&P3,45)
	if weather!=planet.Unknown {
		t.Fatalf("WeatherForDay( *P1,*P2,*P3,1)!=Unknown failed.")
	}
}
{
	P1 := planet.Planet{3000, 90 , 1 }
	P2 := planet.Planet{3000, 90 , 1 }
	P3 := planet.Planet{3000, 90 , 1 }
	weather, _:= planet.WeatherForDay( &P1,&P2,&P3,45)
	if weather!=planet.Sequia {
		t.Fatalf("WeatherForDay( *P1,*P2,*P3,1)!=Sequia failed.")
	}
}
{
	P1 := planet.Planet{5000, 0 , 1 }
	P2 := planet.Planet{5000, 0 , -1 }
	P3 := planet.Planet{4330.127 ,0 , 0 }
	weather, _:= planet.WeatherForDay( &P1,&P2,&P3,30)
	if weather!=planet.TemperaturaOptima {
		t.Fatalf("WeatherForDay( &P1,&P2,&P3,1)!=TemperaturaOptima failed.")
	}
}
}