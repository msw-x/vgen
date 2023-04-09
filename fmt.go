package vgen

func fmtTime(tm time.Time) string {
	return tm.Format("06.01.02.150405")
}

func fmtRelativeTime(tm time.Time) string {
	now := time.Now()
	fmtLst := []string{"06.", "01.", "02.", "15", "04", "05"}
	makeFmtStr := func() {
		nowYear, nowMonth, nowDay := now.Date()
		year, month, day := tm.Date()
		nowHour, nowMin, nowSec := now.Clock()
		hour, min, sec := tm.Clock()
		if year != nowYear {
			return
		}
		fmtLst = fmtLst[1:]
		if month != nowMonth {
			return
		}
		fmtLst = fmtLst[1:]
		if day != nowDay {
			return
		}
		fmtLst = fmtLst[1:]
		if hour != nowHour {
			return
		}
		fmtLst = fmtLst[1:]
		if min != nowMin {
			return
		}
		fmtLst = fmtLst[1:]
		if sec != nowSec {
			return
		}
		fmtLst = fmtLst[1:]
		return
	}
	makeFmtStr()
	fmtStr := strings.Join(fmtLst, "")
	if fmtStr == "" {
		return ""
	}
	return now.Format(fmtStr)
}
