package main

import (
	"encoding/csv"
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"os"
	"strings"
)

func main() {
	// 1. 读取 CSV 文件
	file, err := os.Open("/Users/xsky/Desktop/working/ocpf/test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	longest := 0
	for _, row := range records {
		for _, cell := range row {
			if len([]rune(cell)) > longest {
				longest = len([]rune(cell))
			}
		}
	}

	const (
		maxCharactersPerCell = 128
		maxRows              = 500
	)

	if longest > maxCharactersPerCell {
		longest = maxCharactersPerCell
	}

	// 2. 设置图片参数
	var (
		firstCellWidth        = longest * 12 // 单元格宽度
		nonFirstCellWidth     = 450          // 非第一列单元格宽度
		firstRowCellHeight    = 120          // 单元格高度
		nonFirstRowCellHeight = 60           // 非第一行单元格高度
		fontSize              = float64(20)  // 字体大小
	)

	// 计算图片尺寸
	rows := len(records)
	if rows > maxRows {
		rows = maxRows
	}
	cols := len(records[0])
	imgWidth := firstCellWidth + nonFirstCellWidth*(cols-1)
	imgHeight := 1*firstRowCellHeight + (rows-1)*nonFirstRowCellHeight

	// 创建画布
	dc := gg.NewContext(imgWidth, imgHeight)

	// 设置背景颜色
	dc.SetColor(color.White)
	dc.Clear()

	// 设置字体（确保字体文件路径正确）
	if err := dc.LoadFontFace("/Library/Fonts/Arial Unicode.ttf", fontSize); err != nil {
		panic(err)
	}
	// 3. 绘制表格
	dc.SetColor(color.Black)
	recordCount := 0
	for i, row := range records {
		if recordCount >= maxRows {
			break
		}
		for j, cell := range row {
			var x float64
			if j == 0 || j == 1 {
				x = float64(j * firstCellWidth)
			} else {
				x = float64(firstCellWidth) + float64((j-1)*nonFirstCellWidth)
			}

			var y float64
			if i > 0 {
				y = float64(firstRowCellHeight + (i-1)*nonFirstRowCellHeight)
			}
			cellHeight := firstRowCellHeight
			if i > 0 {
				cellHeight = nonFirstRowCellHeight
			}
			// 绘制单元格边框
			if j == 0 {
				dc.DrawRectangle(x, y, float64(firstCellWidth), float64(cellHeight))
			} else {
				dc.DrawRectangle(x, y, float64(nonFirstCellWidth), float64(cellHeight))
			}
			dc.Stroke()
			if i == 0 {
				maxWidth := 0
				if j == 0 {
					maxWidth = firstCellWidth
				} else {
					maxWidth = nonFirstCellWidth
				}
				// 获取文本行高（用于垂直居中计算）
				_, textHeight := dc.MeasureString("A") // 测量单个字符高度
				// 分割多行文本
				lines := strings.Split(cell, "\n")
				totalTextHeight := textHeight * float64(len(lines)) // 计算总文本高度

				// 计算起始Y坐标（实现垂直居中）
				startY := y + (float64(cellHeight)-totalTextHeight)/2 + textHeight*0.3
				for _, line := range lines { // 按换行符分割
					dc.DrawStringWrapped(line, x+8, startY, 0, 0, float64(maxWidth), -1.5, gg.AlignLeft)
					startY += textHeight * 1.5 // 移动到下一行
				}
			} else {
				// 绘制文本
				if len(cell) > maxCharactersPerCell {
					cell = fmt.Sprintf("%s...", cell[:maxCharactersPerCell])
				}
				dc.DrawString(cell, x+8, y+float64(nonFirstRowCellHeight)/2+10)
			}
		}
		recordCount++
	}

	// 4. 保存图片
	if err := dc.SavePNG("output.png"); err != nil {
		panic(err)
	}
}
