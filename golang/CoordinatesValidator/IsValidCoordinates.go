package kata
import(
  "fmt"
  "regexp"
)

func IsValidCoordinates(coordinates string) bool {
  matched, err := regexp.MatchString(`^[-+]?([0-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`, coordinates)
  fmt.Println(matched)
  fmt.Println(err)
  return matched
}