package user


type Level int

const (
    _ Level  = iota

    LevelNormal
    LevelManager
    LevelAdmin
    LevelRootadmin
)

var Levels = []string{ "", "일반", "팀장", "관리자", "전체관리자" }

type Status int

const (
    _ Status  = iota

    StatusUse
    StatusNotuser
)

var Statuss = []string{ "", "사용", "사용안함" }



func GetLevel(value Level) string {
    i := int(value)
    if i <= 0 || i >= len(Levels) {
        return ""
    }
     
    return Levels[i]
}

func ConvertLevel(value []int) []Level {
     items := make([]Level, 0)

     for item := range value {
         items = append(items, Level(item))
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


