package billing


type Status int

const (
    _ Status  = iota

    StatusWait
    StatusPart
    StatusComplete
)

var Statuss = []string{ "", "입금대기", "부분입금", "입금완료" }

type Giro int

const (
    _ Giro  = iota

    GiroWait
    GiroComplete
)

var Giros = []string{ "", "미발행", "발행" }



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

func GetGiro(value Giro) string {
    i := int(value)
    if i <= 0 || i >= len(Giros) {
        return ""
    }
     
    return Giros[i]
}

func ConvertGiro(value []int) []Giro {
     items := make([]Giro, 0)

     for item := range value {
         items = append(items, Giro(item))
     }
     
     return items
}


