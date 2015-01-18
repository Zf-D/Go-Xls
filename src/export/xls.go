package export

/**
 * Created by Zf_D on 2015-01-16
 */
import (
    "os"
    "errors"
    "time"
    "strings"
    "strconv"
)

func Exec(storeId string, fileName string, header []string, records [][]string) (string, error) {
    //校验列数
    if len(records) == 0 {
        return "", errors.New("错误，数据为空")
    }

    //项目下文件目录
    pjDir, _ := os.Getwd()
    dir := pjDir + "\\www\\downloadResource\\export\\" + storeId

    //检查并生成目录
    err := os.MkdirAll(dir, os.ModePerm)
    if err != nil {
        return "", err
    }

    //文件导出路径
    timeStr := time.Now().Format("2006-01-02 15-04-05")
    filePath := dir +"\\" + fileName + " " + timeStr + ".xls"

    //导出文件
    err = export(filePath, header, records)
    if err != nil {
        return "", err
    }

    filePath = strings.Split(filePath, `www`)[1]
    filePath = strings.Replace(filePath, `\`, `/`, -1)
    return filePath, nil
}

func export(fileName string, header []string, records [][]string) (err error) {
    //创建文件
    f, err := os.Create(fileName)
    if err != nil {
        return
    }
    defer f.Close()

    //写入UTF-8
    f.WriteString("\xEF\xBB\xBF")

    //行数，列数
    cNum := len(header)
    rNum := len(records) + 1

    //写入开始
    f.WriteString(BeginStr)
    f.WriteString(`
    <Worksheet ss:Name="Sheet1">
    <Table ss:ExpandedColumnCount="`+strconv.Itoa(cNum)+`" ss:ExpandedRowCount="`+strconv.Itoa(rNum)+`" x:FullColumns="1" x:FullRows="1"
    ss:DefaultColumnWidth="70" ss:DefaultRowHeight="15">
    `)

    //写入表头
    f.WriteString(`<Row>`)
    for _, h := range header {
        f.WriteString(`<Cell ss:StyleID="s1"><Data ss:Type="String">`+h+`</Data></Cell>`)
    }
    f.WriteString(`</Row>`)

    //写入内容
    for i := range(records) {
        f.WriteString(`<Row>`)
        for _, c := range records[i] {
            f.WriteString(`<Cell><Data ss:Type="String">`+c+`</Data></Cell>`)
        }
        f.WriteString(`</Row>`)
    }

    //写入结束
    f.WriteString(EndStr)

    return
}
