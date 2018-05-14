package lunarcal

import (
	"fmt"
	"strconv"
	"time"
)

//MinYear 最小年
var MinYear = 1900

//MaxYear 最大年
var MaxYear = 2050

//DateFormat 日期格式
var DateFormat = "2006-01-02" //DATE_FORMAT

//StartDate 阴历算法的起始日期
var StartDate = "1900-01-30" //START_DATE
//LunarInfo 数据转换LUNAR_INFO
var LunarInfo = []int{
	0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2,
	0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977,
	0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970,
	0x06566, 0x0d4a0, 0x0ea50, 0x06e95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950,
	0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4, 0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557,
	0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5d0, 0x14573, 0x052d0, 0x0a9a8, 0x0e950, 0x06aa0,
	0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0,
	0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b5a0, 0x195a6,
	0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570,
	0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50, 0x06b58, 0x055c0, 0x0ab60, 0x096d5, 0x092e0,
	0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5,
	0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9, 0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930,
	0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530,
	0x05aa0, 0x076a3, 0x096d0, 0x04bd7, 0x04ad0, 0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45,
	0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0,
	0x14b63}

//ToLunarDate 把阳历日期转换成阴历日期,返回Time
func ToLunarDate(date string) time.Time {
	lunarYear, lunarMonth, lunarDay, _, _ := calculateLunar(date)
	lm := time.January
	switch lunarMonth {
	case 1:
		lm = time.January
	case 2:
		lm = time.February
	case 3:
		lm = time.March
	case 4:
		lm = time.April
	case 5:
		lm = time.May
	case 6:
		lm = time.June
	case 7:
		lm = time.July
	case 8:
		lm = time.August
	case 9:
		lm = time.September
	case 10:
		lm = time.October
	case 11:
		lm = time.November
	case 12:
		lm = time.December
		//	default:
		//		return error
	}
	//the_time := time.Date(2014, 1, 7, 5, 50, 4, 0, time.Local)
	theTime := time.Date(lunarYear, lm, lunarDay, 0, 0, 0, 0, time.Local)
	return theTime
}

//ToLunarStr 把日期转换为中国阴历的日期,返回date
func ToLunarStr(date string) string {
	lunarYear, lunarMonth, lunarDay, leapMonth, leapMonthFlag := calculateLunar(date)
	result := strconv.Itoa(lunarYear) + "年"
	if leapMonthFlag && (lunarMonth == leapMonth) {
		result += "闰"
	}
	if lunarMonth < 10 {
		result += "0" + strconv.Itoa(lunarMonth) + "月"
	} else {
		result += strconv.Itoa(lunarMonth) + "月"
	}
	if lunarDay < 10 {
		result += "0" + strconv.Itoa(lunarDay) + "日"
	} else {
		result += strconv.Itoa(lunarDay) + "日"
	}
	return result
}
func calculateLunar(date string) (lunarYear, lunarMonth, lunarDay, leapMonth int, leapMonthFlag bool) {
	loc, _ := time.LoadLocation("Local")
	i := 0
	temp := 0
	leapMonthFlag = false
	isLeapYear := false

	myDate, err := time.ParseInLocation(DateFormat, date, loc)
	if err != nil {
		fmt.Println(err.Error())
	}
	startDate, err := time.ParseInLocation(DateFormat, StartDate, loc)
	if err != nil {
		fmt.Println(err.Error())
	}

	offset := daysBwteen(myDate, startDate)
	for i = MinYear; i < MaxYear; i++ {
		temp = getYearDays(i) //求当年农历年天数
		if offset-temp < 1 {
			break
		} else {
			offset -= temp
		}
	}
	lunarYear = i

	leapMonth = getLeapMonth(lunarYear) //计算该年闰哪个月

	//设定当年是否有闰月
	if leapMonth > 0 {
		isLeapYear = true
	} else {
		isLeapYear = false
	}

	for i = 1; i <= 12; i++ {
		if i == leapMonth+1 && isLeapYear {
			temp = getLeapMonthDays(lunarYear)
			isLeapYear = false
			leapMonthFlag = true
			i--
		} else {
			temp = getMonthDays(lunarYear, uint(i))
		}
		offset -= temp
		if offset <= 0 {
			break
		}
	}
	offset += temp
	lunarMonth = i
	lunarDay = offset
	return
}
func getYearDays(year int) int {
	sum := 29 * 12
	for i := 0x8000; i >= 0x8; i >>= 1 {
		if (LunarInfo[year-1900] & 0xfff0 & i) != 0 {
			sum++
		}
	}
	return sum + getLeapMonthDays(year)
}

//	计算阴历年闰哪个月 1-12 , 没闰传回 0
func getLeapMonth(year int) int {
	return (int)(LunarInfo[year-1900] & 0xf)
}

//	计算阴历年闰月多少天
func getLeapMonthDays(year int) int {
	if getLeapMonth(year) != 0 {
		if (LunarInfo[year-1900] & 0xf0000) == 0 {
			return 29
		}
		return 30
	}
	return 0

}

// 计算该月总天数
func getMonthDays(lunarYeay int, month uint) int {
	if (month > 31) || (month < 0) {
		fmt.Println("error month")
	}
	// 0X0FFFF[0000 {1111 1111 1111} 1111]中间12位代表12个月，1为大月，0为小月
	bit := 1 << (16 - month)
	if ((LunarInfo[lunarYeay-1900] & 0x0FFFF) & bit) == 0 {
		return 29
	}
	return 30
}

// 计算差的天数
func daysBwteen(myDate time.Time, startDate time.Time) int {
	subValue := float64(myDate.Unix()-startDate.Unix())/86400.0 + 0.5
	return int(subValue)
}
