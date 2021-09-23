package pine

import "time"

func hour(t *time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
}

func cron() {
	for v := range formatChan {
		cronChan <- &CronData{
			GoalFile: v.Ts.Format("2006-01-02 15:04:05"),
			Foramte:  string(v.Foramte),
		}
	}
}
