package customer


type Type int

const (
    _ Type  = iota

    TypeDirect
    TypeOutsourcing
)

var Types = []string{ "", "직영", "위탁관리" }



func GetType(value Type) string {
    i := int(value)
    if i <= 0 || i >= len(Types) {
        return ""
    }
     
    return Types[i]
}

func ConvertType(value []int) []Type {
     items := make([]Type, 0)

     for item := range value {
         items = append(items, Type(item))
     }
     
     return items
}


