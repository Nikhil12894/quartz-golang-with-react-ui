package scheduling

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/abecciu/cronexpr"
	"github.com/lnquy/cron"
)

//Describe is to get human redable text of cron expration
func Describe(exp string) (string, error) {
	exprDesc, err := cron.NewDescriptor(
		cron.Use24HourTimeFormat(true),
		cron.DayOfWeekStartsAtOne(false),
		cron.Verbose(true),
		cron.SetLogger(log.New(os.Stdout, "cron: ", 0)),
		cron.SetLocales(cron.Locale_en),
	)
	if err != nil {
		return "", fmt.Errorf("failed to create CRON expression descriptor: %s", err)
	}

	desc, err := exprDesc.ToDescription(exp, cron.Locale_en)
	if err != nil {
		return "", fmt.Errorf("failed to create CRON expression descriptor: %s", err)
	}
	return desc, nil
}

//NextNScheduledTime is used to get next N schedule for a expration
func NextNScheduledTime(exp string, n uint) []string {
	// get the current time
	now := time.Now()
	// 1. Define two cronJob
	expr1 := cronexpr.MustParse(exp) // parse cron expression will be successful
	times := expr1.NextN(now, n)
	lenthData := int(n)
	response := make([]string, lenthData)
	for i := 0; i < lenthData; i++ {
		response[i] = times[i].Format(time.ANSIC)
	}
	return response

}
