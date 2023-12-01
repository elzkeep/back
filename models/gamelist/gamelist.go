package gamelist


type Status int

const (
    _ Status  = iota

    StatusReady
    StatusFaction
    StatusNormal
    StatusEnd
)

var Statuss = []string{ "", "대기중", "종족선택중", "게임중", "게임종료" }

type Illusionists int

const (
    _ Illusionists  = iota

    IllusionistsNormal
    IllusionistsVp2
    IllusionistsVp1
    IllusionistsVp0
    IllusionistsBan
)

var Illusionistss = []string{ "", "너프 없음", "2, 3점으로 너프", "1, 2점으로 너프", "점수안줌", "사용 안함" }



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

func GetIllusionists(value Illusionists) string {
    i := int(value)
    if i <= 0 || i >= len(Illusionistss) {
        return ""
    }
     
    return Illusionistss[i]
}

func ConvertIllusionists(value []int) []Illusionists {
     items := make([]Illusionists, 0)

     for item := range value {
         items = append(items, Illusionists(item))
     }
     
     return items
}


