package log

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func NewKafakLoger() *KafakLoger {
	log := KafakLoger{}
	return &log
}

type KafakLoger struct {
	DataCollector     sarama.SyncProducer
	AccessLogProducer sarama.AsyncProducer
}

func (l *KafakLoger) Debugf(format string, args ...interface{}) {

}
func (l *KafakLoger) Infof(format string, args ...interface{}) {

}
func (l *KafakLoger) Printf(format string, args ...interface{}) {

}
func (l *KafakLoger) Warnf(format string, args ...interface{}) {

}
func (l *KafakLoger) Warningf(format string, args ...interface{}) {

}
func (l *KafakLoger) Errorf(format string, args ...interface{}) {

}
func (l *KafakLoger) Fatalf(format string, args ...interface{}) {

}
func (l *KafakLoger) Panicf(format string, args ...interface{}) {

}
func (l *KafakLoger) Debug(args ...interface{}) {
	fmt.Println(args)
}
func (l *KafakLoger) Info(args ...interface{}) {
	fmt.Println(args)
}
func (l *KafakLoger) Print(args ...interface{}) {
	fmt.Println(args)
}
func (l *KafakLoger) Warn(args ...interface{}) {
	fmt.Println(args)
}
func (l *KafakLoger) Warning(args ...interface{}) {
	fmt.Println(args)
}
func (l *KafakLoger) Error(args ...interface{}) {
	fmt.Println(args)
}
func (l *KafakLoger) Fatal(args ...interface{}) {
	fmt.Println(args)
}
func (l *KafakLoger) Panic(args ...interface{}) {
	fmt.Println(args)
}
