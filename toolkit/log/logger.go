package log

import (
	"bytes"
	"strings"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
	采用自二次封装的zap，比原zap效率高很多.测试的结论为io.Write所造成的写入耗时。
原zap采用每条都用写Write写入，Sync函数同步到文件。新方案采用缓存处理一次写入同时
同步文件。MaxBufferLine 越大效率越高，最大建议8000. 错误日志默认不缓存，出现错误
请进行处理。

注意:如果程序异常退出缓存中的日志将会丢失，需要考虑数据安全则 MaxBufferLine 应为0.
*/

func init() {
	loggerConfig := &LoggerConfig{
		Name:          "log",
		Dir:           "./logs/",
		MaxBufferLine: 0,
	}

	map_variable := make(map[string]LoggerConfig)
	map_variable["log"] = *loggerConfig
	InitLog(map_variable)
}

const (
	_Mb                   = 1
	_Day                  = 1
	_DefaultMaxBufferLine = 0 // 缓存最大行数,默认不缓存
)

type LoggerConfig struct {
	Name          string `yaml:"name"`
	Dir           string `yaml:"dir"`
	MaxBufferLine int    `yaml:"maxBufferLine"`
}

// 日志接口
type MLogger interface {
	Error(...string)
	Info(...string)
	Warn(...string)
}

var (
	DfLogger MLogger
	Loggers  = make(map[string]MLogger)
	mu       sync.Mutex
	encoder  zapcore.Encoder
	lagTime  = time.Now().Format("2006-01-02 15:04:05")

	// 增加简写
	Info = func(arg ...string) {
		DfLogger.Info(arg...)
	}
	Error = func(arg ...string) {
		DfLogger.Info(arg...)
	}
	Warn = func(arg ...string) {
		DfLogger.Info(arg...)
	}
)

type mZap struct {
	zap           *zap.Logger
	out           zapcore.WriteSyncer
	cache         []string
	maxBufferLine int // 缓冲行数,最小为0不缓存
}

type mLogger struct {
	errLog    *mZap
	commonLog *mZap
	logBuffer bytes.Buffer
}

// 打印错误日志
func (m *mLogger) Error(msg ...string) {
	_str := strings.Join(msg, "")
	_str = stringContact(lagTime, " Error ", _str, zapcore.DefaultLineEnding)
	m.writeErrCache(_str)
}

// 打印提示日志
func (m *mLogger) Info(msg ...string) {
	_str := strings.Join(msg, "")
	_str = stringContact(lagTime, " Info ", _str, zapcore.DefaultLineEnding)
	m.writeInfoCache(_str)
}

// 打印告警日志
func (m *mLogger) Warn(msg ...string) {
	_str := strings.Join(msg, "")
	_str = stringContact(lagTime, " Warn ", _str, zapcore.DefaultLineEnding)
	m.writeInfoCache(_str)
}

// 缓存时间
func getLogNowTime() {
	for {
		mu.Lock()
		lagTime = time.Now().Format("2006-01-02 15:04:05")
		mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}

// 写入常规缓存  make([]string,0,2000)
func (m *mLogger) writeInfoCache(data string) {
	mu.Lock()
	m.commonLog.cache = append(m.commonLog.cache, data)
	if len(m.commonLog.cache) > m.commonLog.maxBufferLine {
		m.outFile(m.commonLog)
	}
	mu.Unlock()
}

// 写入error的缓存
func (m *mLogger) writeErrCache(data string) {
	mu.Lock()
	m.errLog.cache = append(m.errLog.cache, data)
	if len(m.errLog.cache) > m.errLog.maxBufferLine {
		m.outFile(m.errLog)
	}
	mu.Unlock()
}

// 输出日志到文件
func (m *mLogger) outFile(_log *mZap) {
	m.logBuffer.WriteString(strings.Join(_log.cache, ""))
	_log.out.Write(m.logBuffer.Bytes())
	m.logBuffer.Reset()
	_log.zap.Sync()
	_log.cache = _log.cache[:0]
	_log.cache = make([]string, 0, _log.maxBufferLine+1)
}

// 创建获取日志
func newLogger(zap *zap.Logger, zapErr *zap.Logger, maxBufferLine int, sync zapcore.WriteSyncer, errSync zapcore.WriteSyncer) MLogger {
	if maxBufferLine < 0 {
		maxBufferLine = _DefaultMaxBufferLine
	}
	return &mLogger{
		commonLog: &mZap{
			zap:           zap,
			out:           sync,
			cache:         make([]string, 0, maxBufferLine+1),
			maxBufferLine: maxBufferLine,
		},
		errLog: &mZap{
			zap:           zapErr,
			out:           errSync,
			cache:         make([]string, 0, 1),
			maxBufferLine: 0,
		},
	}
}

// 初始化日志
func InitLog(logConfigs map[string]LoggerConfig) {
	encoder = getEncoder()
	one := true
	go getLogNowTime()
	for key := range logConfigs {
		logger := logConfigs[key]
		Loggers[key] = createZapLog(logger, 1000*_Mb, 7*_Day)
		if one {
			DfLogger = Loggers[key]
			one = false
		}
	}
}

// 创建zapLog 包含errorLog
func createZapLog(conf LoggerConfig, mb int, day int) MLogger {
	synced := getLogWriter(conf.Dir+conf.Name, mb, day)
	Core := zapcore.NewCore(encoder, synced, zapcore.DebugLevel)
	suLog := zap.New(Core)

	errorSynced := getLogWriter(conf.Dir+conf.Name+`_err`, mb, day)
	errorCore := zapcore.NewCore(encoder, errorSynced, zapcore.DebugLevel)
	errorLog := zap.New(errorCore)
	return newLogger(suLog, errorLog, conf.MaxBufferLine, synced, errorSynced)
}

// 设置编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// log文件基本配置
func getLogWriter(fileName string, maxSize int, day int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName, // "./test.log",
		MaxSize:    maxSize,  // 兆(mb)
		MaxBackups: 5,        // 份数
		MaxAge:     day,      // 天
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func stringContact(item ...string) string {
	var itemBuffer bytes.Buffer
	for _, v := range item {
		itemBuffer.WriteString(v)
	}
	return itemBuffer.String()
}
