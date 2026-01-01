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
    //remove white space at the beginning
    for i := 0; i < len(strar); i++{
        if !unicode.IsSpace(strar[i]){
            charindex = i;
            break
        }   
    }
    strar = strar[charindex:]  

    //return 0 if blank string or if the first two characters are not digits
    if (len(strar) == 0) || (len(strar) > 1 && !unicode.IsDigit(strar[0]) && !unicode.IsDigit(strar[1])){
        return 0;
    }  

    s = string(strar)

    //check if integer is positive or negative
    var negative bool = false 
    if s[0] == '-' {
        s = strings.Replace(s, "-", "", 1)
        negative = true
    } else if s[0] == '+' {
        s = strings.Replace(s, "+", "", 1)
    } 

    //if any characters are in the string, remove everything after that point
    for i, c := range s {  
        if !unicode.IsDigit(c){
            s = s[:i]
            break
        }
    } 
    //convert sanitised string to int
    val, _ := strconv.Atoi(s);
    
    //rounding if number isn't 32 bit
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