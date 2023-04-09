package bazi

// 输入tgA，dzB，顺序为Gan，Zhi顺序, 从0开始，如 癸在未，则输入 （9，7）=8 墓 ；丁在午，输入（3，6）=3 临官
var ChangShengDict = [][]int{
	{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0},
	{6, 5, 4, 3, 2, 1, 0, 11, 10, 9, 8, 7},
	{10, 11, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 11, 10},
	{10, 11, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 11, 10},
	{7, 8, 9, 10, 11, 0, 1, 2, 3, 4, 5, 6},
	{0, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	{4, 5, 6, 7, 8, 9, 10, 11, 0, 1, 2, 3},
	{3, 2, 1, 0, 11, 10, 9, 8, 7, 6, 5, 4},
}
var ChangSheng = [12]string{"长生", "沐浴", "冠带", "临官（禄）", "帝旺", "病", "衰", "死", "墓", "绝", "胎", "养"}

// GetChangShengFromNumber 从数字获得长生名, 0-11
func GetChangShengFromNumber(nValue int) string {
	if nValue > 11 {
		return ""
	}
	return ChangSheng[nValue]
}

// NewShiShen 新建长生宫
func NewChangSheng(nValue int) *TChangSheng {
	nValue %= 10
	pChangSheng := TChangSheng(nValue)
	return &pChangSheng
}

// NewChangShengFromGan 从日干和目标支获取十神
// nDayGan 日干 nZhi 目标支
func NewChangShengFromGan(nDayGan int, nZhi int) *TChangSheng {
	return NewChangSheng(ChangShengDict[nDayGan][nZhi])
}

// ToString 转换成可阅读的字符串
func (m *TChangSheng) ToString() string {
	return m.String()
}

// ToInt 转换成int
func (m *TChangSheng) ToInt() int {
	return m.Value()
}

// TChangSheng 长生宫
type TChangSheng int

// Value 转换成int
func (m *TChangSheng) Value() int {
	return (int)(*m)
}

// String 转换成可阅读的字符串
func (m *TChangSheng) String() string {
	return GetShiShenFromNumber(m.Value())
}
