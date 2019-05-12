package main

import (
	"bufio"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"os"
	"time"
)

func createExcel() {
	f := excelize.NewFile()
	//创建一个新的Excel
	index := f.NewSheet("Sheet2")
	//向一个单元格进行赋值
	f.SetCellValue("Sheet2", "A2", "Hello World.")
	f.SetCellValue("Sheet1", "B2", 100)
	//设置指定sheet为活动sheet
	f.SetActiveSheet(index)
	err := f.SaveAs("E:/go/GoStudy_001/files/test.xlsx")
	if err != nil {
		fmt.Printf("保存excel失败，err=%v", err)
	}
}

//保存excel
func saveExcel(f *excelize.File, path string) error {
	err := f.SaveAs(path)
	return err
}

func readExcel() {
	f, err := excelize.OpenFile("E:/go/GoStudy_001/files/周报_name_yyyy-MM-dd.xlsx")
	if err != nil {
		fmt.Printf("打开Excel失败，err=%v", err)
		return
	}

	//获取一个指定表的指定单元格
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Printf("获取指定单元格失败，err=%v")
		return
	}
	fmt.Println(cell)
	//获取sheet中的rows
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Printf("获取行失败，err=%v")
		return
	}
	for rIndex, row := range rows {
		for cIndex, colCell := range row {
			fmt.Printf("(%v,%v)=%v\n", rIndex, cIndex, colCell)
		}
	}
}

func writeWeeklyReport() {
	//打开Excel
	f, err := excelize.OpenFile("E:/go/GoStudy_001/files/周报_name_yyyy-MM-dd.xlsx")
	if err != nil {
		fmt.Printf("打开Excel失败，err=%v", err)
		return
	}
	fmt.Println("已经打开Excel模板...")
	fmt.Printf("当前时间为:%v", time.Now())
	excelName := ""
	for {
		for {
			fmt.Println("姓名：")
			name := ""
			fmt.Scanln(&name)

			excelName = fmt.Sprintf("周报_%v_%v",
				name, time.Now().Format("2006_01_02"))
			fmt.Printf("根据您的输入，将开启名为《%v》的Excel文件\n", excelName)
			fmt.Println("确认请输入y,输入其他任意字符返回上一步")
			confirm := ""
			fmt.Scanln(&confirm)
			if confirm == "y" {
				break
			}
		}
		fmt.Printf("现在开始编辑%v.xlsx\n", excelName)
		//编辑表头
		f.SetCellValue("Sheet1", "B1", fmt.Sprintf("周报(%v)",
			time.Now().Format("2006_01_02")))

		//读取内容文件
		fmt.Println("现在开始读入内容文件...")
		file,err:=os.OpenFile("E:/go/GoStudy_001/files/content.txt",os.O_RDWR|os.O_APPEND,0666)
		if err != nil {
			fmt.Printf("打开文件失败，err=%v",err)
		}
		defer file.Close()

		reader:=bufio.NewReader(file)

		content:=""
		for{
			str,err:=reader.ReadString('\n')
			if err==io.EOF{
				break
			}
			content+=str
		}

		writer:=bufio.NewWriter(file)
		writer.WriteString(content)
		writer.Flush()

		f.SetCellValue("Sheet1","D4",content)
		fmt.Println("内容读取完毕...")

		fmt.Println("现在进行保存...")
		err = f.SaveAs(fmt.Sprintf("E:/go/GoStudy_001/files/%v.xlsx",excelName))
		if err != nil {
			fmt.Printf("保存excel失败，err=%v", err)
		}
		fmt.Println("保存完毕,程序将在3秒后自动退出...")
		time.Sleep(3*time.Second)
		break
	}
}

func main() {
	//createExcel()
	//readExcel()
	writeWeeklyReport()
}
