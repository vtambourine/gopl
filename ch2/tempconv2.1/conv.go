package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KtoC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func CtoK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}
