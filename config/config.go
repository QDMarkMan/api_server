package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

// Config ftype
type Config struct {
	Name string
}

// 初始化config
func (config *Config) initConfig() error {
	if config.Name != "" {
		viper.SetConfigFile(config.Name) // 如存在指定的配置文件，使用指定的配置文件
	} else {
		viper.AddConfigPath("conf") // 使用默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")     // 设置配置文件类型
	viper.AutomaticEnv()            // 自动读取env
	viper.SetEnvPrefix("APISERVER") // 设置env前缀
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper 解析配置文件
		return err
	}
	return nil
}

// 监听config  通过 viper + fsnotify 实现
func watchConfig(c *Config) {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}

// 初始化日志
func initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	log.InitWithConfig(&passLagerCfg)
}

// Init config file
func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}
	// 监听配置文件
	watchConfig(&c)
	// 初始化log
	initLog()

	return nil
}
