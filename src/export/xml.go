package export

/**
 * Created by Zf_D on 2015-01-16
 */

/**
    需要更改：
    ExpandedColumnCount 列数
    ExpandedRowCount 行数
    可选：
    Styles 自定义样式
 */
const (
    BeginStr = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<?mso-application progid="Excel.Sheet"?>
<Workbook xmlns="urn:schemas-microsoft-com:office:spreadsheet" xmlns:o="urn:schemas-microsoft-com:office:office"
          xmlns:x="urn:schemas-microsoft-com:office:excel" xmlns:ss="urn:schemas-microsoft-com:office:spreadsheet"
          xmlns:html="http://www.w3.org/TR/REC-html40" xmlns:dt="uuid:C2F41010-65B3-11d1-A29F-00AA00C14882">
    <DocumentProperties xmlns="urn:schemas-microsoft-com:office:office">
        <Author>Zf_D</Author>
        <Created>2015-01-16T07:54:55Z</Created>
        <LastSaved>2015-01-16T07:55:33Z</LastSaved>
    </DocumentProperties>
    <CustomDocumentProperties xmlns="urn:schemas-microsoft-com:office:office">
        <KSOProductBuildVer dt:dt="string">2052-9.1.0.4885</KSOProductBuildVer>
    </CustomDocumentProperties>
    <ExcelWorkbook xmlns="urn:schemas-microsoft-com:office:excel">
        <WindowWidth>21510</WindowWidth>
        <WindowHeight>10350</WindowHeight>
        <ProtectStructure>False</ProtectStructure>
        <ProtectWindows>False</ProtectWindows>
    </ExcelWorkbook>
    <Styles>
        <Style ss:ID="Default" ss:Name="Normal">
            <Alignment/>
            <Borders/>
            <Font ss:FontName="宋体" x:CharSet="134" ss:Size="12"/>
            <Interior/>
            <NumberFormat/>
            <Protection/>
        </Style>
        <Style ss:ID="s1">
            <Font ss:FontName="宋体" x:CharSet="134" ss:Size="12" ss:Bold="1"/>
        </Style>
    </Styles>
    `

    EndStr = `
</Table>
    <WorksheetOptions xmlns="urn:schemas-microsoft-com:office:excel">
        <PageSetup>
            <Header x:Margin="0.511111111111111"/>
            <Footer x:Margin="0.511111111111111"/>
        </PageSetup>
    <Selected/>
    <TopRowVisible>0</TopRowVisible>
    <LeftColumnVisible>0</LeftColumnVisible>
    <PageBreakZoom>100</PageBreakZoom>
    <Panes>
        <Pane>
            <Number>1</Number>
            <ActiveRow>1</ActiveRow>
            <ActiveCol>1</ActiveCol>
            <RangeSelection>A1A1</RangeSelection>
        </Pane>
    </Panes>
    <ProtectObjects>False</ProtectObjects>
    <ProtectScenarios>False</ProtectScenarios>
    </WorksheetOptions>
    </Worksheet>
</Workbook>
    `
)
