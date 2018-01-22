### Simple message queue 

[![Build Status](https://travis-ci.org/chlins/me.svg?branch=master)](https://travis-ci.org/chlins/me)

#### Features 
* Simple and easy using.
* Provide message mode pub/sub via topic.

#### Usage:
`./me`

regInfo:
```
{
  "role"  // producer/consumer
  "topic" // set topic name 
}
```
cases: 

  producer:

    nc 127.0.0.1 8001 and send Reg info {"role":"p","topic":"test"}

  consumer:

    nc 127.0.0.1 8001 and send Reg info {"role":"c","topic":"test"}

![](assets/producer.png)
![](assets/consumer.png)