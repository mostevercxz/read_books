### 7.6
Just add a case "K" in switch
```go
case "K":
f.Celsius = KToC(Kelvin(value))
return nil

type Kelvin float64
func (k Kelvin) String() {return fmt.Sprintf("%gK", k)}
func KToC(k Kelvin) Celsius {return Celsius(float64(k) - 273.15)}
```

### 7.7
Because the String() method of Celsius
The String() method formats the flag's value for use in command-line help messages.
