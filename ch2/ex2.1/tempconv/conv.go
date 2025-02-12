package tempconv

// CToF преобразует температуру по Цельсию в температуру по Фаренгейту.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC преобразует температуру по Фаренгейту в температуру по Цельсию
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToF Kelvin to Fahrenheit
// Formula: (0K − 273.15) × 9/5 + 32 = -459.7°F
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit((k-273.15)*9/5 + 32)
}

// KToC Kelvin to Celsius
// Formula: 0K − 273.15 = -273.1°C
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
