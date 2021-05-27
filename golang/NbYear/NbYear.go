package kata
// import "math"

func NbYear(p0 int, percent float64, aug int, p int) (year int) {
  // your code
  for p0 < p { 
    p0 = p0 + int(float64(p0) * percent / 100) + aug
    year  += 1
  }
  return 
}