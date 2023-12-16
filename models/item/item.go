package item


type Type int

const (
    _ Type  = iota

    TypeText
    TypeSelect
    TypeStatus
)

var Types = []string{ "", "Text", "Select", "Status" }

type Status int

const (
    _ Status  = iota

    StatusGood
    StatusWarning
    StatusDanger
    StatusNotuse
)

var Statuss = []string{ "", "적합", "부적합", "요주의", "해당없음" }



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

func GetStatus(value Status) string {
    i := int(value)
    if i <= 0 || i >= len(Statuss) {
        return ""
    }
     
    return Statuss[i]
}

func ConvertStatus(value []int) []Status {
     items := make([]Status, 0)

     for item := range value {
         items = append(items, Status(item))
     }
     
     return items
}


