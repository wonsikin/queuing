package goconf

import "github.com/CardInfoLink/log"

// BizConfig 业务配置信息
type BizConfig struct {
	Seq int32 `yaml:"seq"` // 序列码
}

// UnmarshalYAML 自定义的解析YAML的方法
func (c *BizConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var aux struct {
		Seq int32 `yaml:"seq"` // 序列码
	}

	if err := unmarshal(&aux); err != nil {
		log.Errorf("unmarshal error %s", err)
	}

	c.Seq = aux.Seq
	return nil
}

// MarshalYAML 自定义的将对象转换成data的方法
func (c *BizConfig) MarshalYAML() (interface{}, error) {
	var aux struct {
		Seq int32 `yaml:"seq"`
	}

	aux.Seq = c.Seq
	return aux, nil
}
