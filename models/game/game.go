package game


type Status int

const (
    _ Status  = iota

    StatusReady
    StatusFaction
    StatusNormal
    StatusEnd
)

var Statuss = []string{ "", "준비중", "종족선택중", "게임중", "게임종료" }



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


