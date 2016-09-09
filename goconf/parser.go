package goconf

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync/atomic"
	"time"

	"github.com/CardInfoLink/log"
	"gopkg.in/yaml.v2"
)

// BizCfg 全局的业务对象
var BizCfg = &BizConfig{}

const (
	layout = "2006-01-02 15:04:05"
)

func init() {
	getWorkDir()

	var err error
	BizCfg, err = ReadBizFile()
	if err != nil {
		log.Errorf("Read biz configuration error %s", err)
		os.Exit(4)
	}

	log.Infof("bizcfg is %+v", BizCfg)

	// TODO 定时任务，每到零点就清零
	go crontab()
}

// 定时任务
func crontab() {
	// 计算到下一次零点还差多久
	midnight, _ := time.ParseInLocation(layout, time.Now().AddDate(0, 0, 1).Format("2006-01-02")+" 00:00:00", time.Local)
	subTime := midnight.Sub(time.Now())
	log.Infof("Process method SeqReset after %s", subTime)

	time.AfterFunc(subTime, func() {
		// 到点先执行一下序列重置
		seqReset()

		// 24小时候再执行序列重置
		tick := time.Tick(time.Hour * 24)
		for {
			select {
			case <-tick:
				seqReset()
			}
		}
	})
	// 主线程阻塞
	select {}
}

func seqReset() {
	log.Infof("Before processing mothod SeqReset(), seq is %d", BizCfg.Seq)
	BizCfg.Seq = 0
	err := WriteBizConfigToFile()
	if err != nil {
		log.Errorf("Write bizCfg error: %s", err)
	}

}

// ReadBizFile 读取业务配置
func ReadBizFile() (cfg *BizConfig, err error) {
	filePath := fmt.Sprintf("%s/config/biz.yaml", WorkDir)

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Errorf("Read config file error %s", err)
		return nil, err
	}

	log.Debugf("content is %s", string(content))

	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		log.Errorf("Unmarshal yaml file error %s", err)
		return nil, err
	}

	return cfg, err
}

// WriteBizConfigToFile 将配置数据保存到配置文件中
func WriteBizConfigToFile() (err error) {
	filePath := fmt.Sprintf("%s/config/biz.yaml", WorkDir)

	log.Debugf("bizCfg is %+v", BizCfg)
	data, err := yaml.Marshal(&BizCfg)
	if err != nil {
		log.Errorf("Marshal BizCfg error %s", err)
		return err
	}

	err = WriteDataToFile(filePath, data)
	if err != nil {
		log.Errorf("Write data to file error: %s", err)
		return err
	}

	return nil
}

// WriteDataToFile 将数据保存到文件中
func WriteDataToFile(filePath string, data []byte) (err error) {
	perm := os.ModePerm
	err = ioutil.WriteFile(filePath, data, perm)
	return err
}

// SeqTick 序列加一
func SeqTick() int32 {
	atomic.AddInt32(&(BizCfg.Seq), 1)

	go func(cfg *BizConfig) {
		err := WriteBizConfigToFile()
		if err != nil {
			log.Errorf("Write bizCfg error: %s", err)
		}
	}(BizCfg)
	return BizCfg.Seq
}
