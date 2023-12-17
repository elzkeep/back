package report


type Status int

const (
    _ Status  = iota

    StatusNewer
    StatusIng
    StatusCheck
    StatusComplete
)

var Statuss = []string{ "", "신규", "점검중", "점검완료", "작성완료" }



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


