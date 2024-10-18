package execl
import (
    "strings"
    "fmt"
    "errors"
    excelize "github.com/xuri/excelize/v2"
)
type ExeclFile struct {
    File *excelize.File
}
func NewExeclFile() *ExeclFile {
    f := excelize.NewFile()
    return &ExeclFile{
        File: f,
    }
}

//创建sheet工作表
func(e *ExeclFile) CreateSheet(sheetName string) error {
    index,err := e.File.GetSheetIndex(sheetName)
    if err != nil {
        return err
    }
    if index != -1 {
        return errors.New(fmt.Sprintf("%v exist",sheetName))
    }
    _, err = e.File.NewSheet(sheetName)
    if err != nil {
        return err
    }

    return nil
}
//修改sheet名称
func (e *ExeclFile) SetSheetName(oldSheetName,newSheetName string) error {
    err := e.File.SetSheetName(oldSheetName, newSheetName)
    if err != nil {
        return err
    }
    return nil
}

//sheet增加多列(可间隔，可不间隔),并设置表头
func (e *ExeclFile) AddSheetHeaderCols(sheetName string,colinfo map[int]string) error {
    if colinfo == nil {
        return errors.New("col info invalid")
    }
    for index,value := range colinfo {
        err := e.AddSheetHeaderCol(sheetName,value,index)
        if err != nil {
            return err
        }
    }

    return nil
}
//sheet增加单一列，并设置表头
func (e *ExeclFile) AddSheetHeaderCol(sheetName,header string,col int) error {
     chars := generateNames(col)
     err := e.File.InsertCols(sheetName,chars[col],1)
     if err != nil {
         return err
     }
     e.File.SetCellValue(sheetName, fmt.Sprintf("%v%v",chars[col],"1"), header)

     return err
}
//sheet页批量初始化表头
func (e *ExeclFile) SetSheetHeader(sheetName string,headers []string) error {
    if len(headers) <= 0 {
        return errors.New("header is null")
    }
    chars := generateNames(len(headers))
    for index,header := range headers {
        e.File.SetCellValue(sheetName, fmt.Sprintf("%v%v",chars[index],"1"), header)
    }

    return nil
}
//更新固定列数据
func (e *ExeclFile) UpdateFixColData(sheetName string,col int,data map[int]interface{}) error {
    if data == nil {
        return errors.New("data is empty")
    }
    chars := generateNames(col)
    for index,value := range data {
        err := e.File.SetCellValue(sheetName, fmt.Sprintf("%v%v",chars[col],index), value)
        if err != nil {
             return err
        }
    }

    return nil
}
type ExeclData struct {
    Col int
    Row int
    Data interface{}
}
//更新随机位置数据
func (e *ExeclFile) UpdateRandomLocationData(sheetName string,datas []*ExeclData) error {
    if len(datas) <= 0 {
        return errors.New("data invalid")
    }

    for _,v := range datas {
        colindex := generateNames(v.Col)
        err := e.File.SetCellValue(sheetName, fmt.Sprintf("%v%v",colindex[v.Col],v.Row), v.Data)
        if err != nil {
            return err
        }
    }

    return nil
}

//保存文件
func (e *ExeclFile) SaveFile(filepath string) error {
    if err := e.File.SaveAs(filepath); err != nil {
        return err
    }

    return nil
}

//关闭file
func (e *ExeclFile) CloseFile() {
     if err := e.File.Close(); err != nil {
        return
     }
     return
}

//数字列转sheet列索引
func generateNames(n int) []string {
    fmt.Printf("n :%v\n",n)
    var names []string
    for i := 0; i < n; i++ {
        names = append(names, toAlphabeticName(i))
    }
    return names
}
func toAlphabeticName(index int) string {
    var name string
    for ; index >= 0; index = index / 26 - 1 {
        name = string(rune(index%26+'a')) + name
    }
    return strings.ToUpper(name)
}
