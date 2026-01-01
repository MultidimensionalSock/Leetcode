import
(
    "fmt"
    "strings"
    "strconv"
    "unicode"
)

func myAtoi(s string) int {
    strar := []rune(s) 
    charindex := 0  
    for i := 0; i < len(strar); i++{
        if !unicode.IsSpace(strar[i]){
            charindex = i;
            break
        }  
        fmt.Println(string(strar))
    }
    strar = strar[charindex:] 
    fmt.Println(string(strar))

    if len(strar) == 0{
        return 0;
    } 

    if len(strar) > 1 && !unicode.IsDigit(strar[0]) && !unicode.IsDigit(strar[1]){
        fmt.Println("two non digits")
        return 0
    } 
    s = string(strar)
    var negative bool = false
    fmt.Println("string after strar" + s)
    if s[0] == '-' {
        s = strings.Replace(s, "-", "", 1)
        negative = true
    } else if s[0] == '+' {
        s = strings.Replace(s, "+", "", 1)
    }
    fmt.Println(s)
    for i, c := range s {
        fmt.Println(string(c))
        fmt.Println(unicode.IsDigit(c))
        if !unicode.IsDigit(c){
            s = s[:i]
            break
        }
    }
    fmt.Println(s)
    val, _ := strconv.Atoi(s);
    
    if negative{
        val = -val
        if val < -2147483648{
            val = -2147483648
        }
    }
    if (val > 2147483647) {
        val = 2147483647 
    }
    return val;
}