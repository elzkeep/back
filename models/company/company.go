package company


type Type int

const (
    _ Type  = iota

    TypeWork
    TypeBuilding
)

var Types = []string{ "", "점검", "건물" }



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


