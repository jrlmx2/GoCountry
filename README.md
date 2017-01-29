[![rcard](https://goreportcard.com/badge/github.com/jrlmx2/GoCountry)](https://goreportcard.com/report/github.com/jrlmx2/GoCountry)

GoCountry is a simple packge for dealing with countries.

# Country From Text

Given the text "Example full countries: Benin, azerbaijan. Example three letter country code: BOL. Example two letter country code: BR Example number: 408"

```go
import "github.com/jrlmx2/GoCountry"
options := &gocountry.Options{
  Full:      true, // search full country name
  CodeTwo:   true, // search for 2 character country code
  CodeThree: true, // search for 3 character country code
  Number:    true, //search for country number. (MUST BE 0 Padded)
}

countries := gocountry.Search(options, "Example full countries: Benin, azerbaijan. Example three letter country code: BOL. Example two letter country code: BR Example number: 408")
for _, country := range countries {
  fmt.Printf("Found country in text: %s, %s, %s, %s\n", country.Full(), country.CodeTwo(), country.CodeThree(), country.Number())
}
```

# Country From Text, only full country names

Given the text "Example full countries: Benin, azerbaijan. Example three letter country code: BOL. Example two letter country code: BR Example number: 408"

```go
import "github.com/jrlmx2/GoCountry"
countries := gocountry.Search(&gocountry.Options{Full:true}, "Example full countries: Benin, azerbaijan. Example three letter country code: BOL. Example two letter country code: BR Example number: 408")
for _, country := range countries {
  fmt.Printf("Found country in text: %s, %s, %s, %s\n", country.Full(), country.CodeTwo(), country.CodeThree(), country.Number())
}
```

# Lookup country number code

```go
import "github.com/jrlmx2/GoCountry"

country := FindByNumber(4)
if country != nil {
  fmt.Printf("Found country by number: %s, %s, %s, %s\n", country.Full(), country.CodeTwo(), country.CodeThree(), country.Number())
}

```

# How to get

```
go get github.com/jrlmx2/GoCountry
```

# Contribution Welcomed !
