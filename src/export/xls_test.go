package export

/**
 * Created by Zf_D on 2015-01-13
 */
import (
	"testing"
//	. "github.com/smartystreets/goconvey/convey"
	"os"
	"time"
	"fmt"
)

func Benchmark_export(b *testing.B) {
	var err error
	//每行10列
	header := []string {"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	records := [][]string {}
	for i:=0; i<1000; i++ {
		tRecord := []string {
			"这是一列","这是一列","这是一列","这是一列","这是一列",
			"这是一列","这是一列","这是一列","这是一列","这是一列",
		}
		records = append(records, tRecord)
	}

	b.N = 50
	var totalTime float64
	for i := 0; i < b.N; i++ {
		timeNow := time.Now()
		err = export("test_bench.xls", header, records)
		if err != nil {
			b.Log(err)
			break
		}
		timeDiff := time.Now().Sub(timeNow).Seconds()
		totalTime += timeDiff
		fmt.Println("第", i+1, "次", "->", timeDiff, "秒")
//		deleteSrc(createSrc()+"\\test_bench.xls")
	}
	fmt.Println("平均->", totalTime/float64(b.N), "秒")
}

//func Test_export(t *testing.T) {
//	Convey("测试", t, func() {
//		var err error
//		var filePath string
//		var header []string
//		var records [][]string
//
//		Convey("测试 export", func() {
//			err = export("不存在路径\\" + "xmlTestFile.xls",  header, records)
//			So(err, ShouldNotEqual, nil)
//
//			header = []string {"A", "B", "C"}
//			records = [][]string {
//				{"1", "2", "3"},
//				{"4", "5"},
//			}
//			err = export("xmlTestFile.xls",  header, records)
//			defer deleteSrc(createSrc()+"\\xmlTestFile.xls")
//			So(err, ShouldEqual, nil)
//		})
//
//		Convey("测试 Exec", func() {
//			path := createSrc() + "\\www\\"
//			defer deleteSrc(path)
//
//			header = []string {"A", "B", "C"}
//			records = [][]string {}
//			filePath, err = Exec("10000", "测试报表", header, records)
//			So(err, ShouldNotEqual, nil)
//
//			header = []string {"A", "B", "C"}
//			records = [][]string {{"1", "2", "3"}}
//			filePath, err = Exec("错误的路径<>", "测试报表", header, records)
//			So(err, ShouldNotEqual, nil)
//
//			header = []string {"A", "B", "C"}
//			records = [][]string {{"1", "2", "3"}}
//			filePath, err = Exec("10000", "错误\\测试报表", header, records)
//			So(err, ShouldNotEqual, nil)
//
//			header = []string {"A", "B", "C"}
//			records = [][]string {{"1", "2", "3"}}
//			filePath, err = Exec("10000", "测试报表", header, records)
//			So(err, ShouldEqual, nil)
//		})
//
//	})
//}

func createSrc() (path string) {
	path, _ = os.Getwd()
	os.MkdirAll(path, os.ModePerm)
	return
}

func deleteSrc(path string) {
	os.RemoveAll(path)
}
