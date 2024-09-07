# timeconv
Time String converter to Go time.Duration


StrToDuration(val string, defvalue time.Duration) (time.Duration, error) 

Convert string to time.Duration

-------------------------------
| string | * time.Millisecond | 
-------------------------------
| 10s    | 10000              |
| 1m     | 60000              |
| 250ms  | 250                |
| 125    | 125                |
-------------------------------

defvalue is returned if the string is blank (defvalue is in milliseconds)

