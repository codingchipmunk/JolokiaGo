package events

import "github.com/codingchipmunk/jolokiago/java"

type Client interface {
	ClientID() string
	JolokiaBaseURL() string
	FullURL() string
	SubscribeToMBean(bean java.MBean) error
//	GetSubscribedMBeans() []java.MBean
//	UnsubscribeFromMBean(bean java.MBean) (beanRemoved bool, err error)
}

type Converter interface{
	WriteToChannel(target chan<- EventData)
	Run()
	Stop()
}
