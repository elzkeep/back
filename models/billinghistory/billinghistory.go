package billinghistory


type Type int

const (
    _ Type  = iota

    TypeAccount
    TypeGiro
    TypeCash
    TypeCard
    TypeCms
    TypeEtc
)

var Types = []string{ "", "이체", "지로", "현금", "카드", "CMS", "기타" }



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


