package cron

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

var (
	c = cron.New()
)

func Init() {
	logrus.Info("init cron")
	c.Start()
}

func GetCron() *cron.Cron {
	return c
}

func StopCron() {
	c.Stop()
}
