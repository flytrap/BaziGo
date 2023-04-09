package bazi

// 甲己之年丙作首；乙庚之岁戊为头；
// 丙辛必定寻庚上，丁壬壬位顺水流；
// 若问戊癸何处起，甲寅之上好寻求。

var WuHuDunMap = map[string]string{
	"甲": "丙寅",
	"己": "丙寅",
	"乙": "戊寅",
	"庚": "戊寅",
	"丙": "庚寅",
	"辛": "庚寅",
	"丁": "壬寅",
	"壬": "壬寅",
	"戊": "甲寅",
	"癸": "甲寅",
}

func GetYearFirstMonthGanZhi(yGan int) int {
	return (yGan/2)*12 + 2
}

func GetYearFirstMonthGanZhiString(yGan string) string {
	m, ok := WuHuDunMap[yGan]
	if ok {
		return m
	}
	return ""
}
