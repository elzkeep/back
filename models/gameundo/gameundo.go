package gameundo


type Status int

const (
    _ Status  = iota

    StatusWait
    StatusComplete
)

var Statuss = []string{ "", "Wait", "Complete" }



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


