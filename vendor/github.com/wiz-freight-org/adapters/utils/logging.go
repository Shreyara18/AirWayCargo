package utils

import (
	"net"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
	LevelFatal = "FATAL"

	LoggingRequestPathKey     = "request_path"
	LoggingRequestBodyKey     = "request_body"
	LoggingRequestMethodKey   = "request_method"
	LoggingStatusCodeKey      = "status_code"
	LoggingResponseTimeKey    = "response_time"
	LoggingResponseTimeInSKey = "response_time_in_s"
	LoggingResponseKey        = "response"
	LoggingClientIPKey        = "client_ip"
	LoggingUserAgentKey       = "user_agent"
)

type LoggerConfig struct {
	refID   string
	appName string
	level   zapcore.Level
}

type Logger struct {
	*zap.Logger
}

func NewLoggerConfig(appName string) *LoggerConfig {
	return &LoggerConfig{refID: uuid.NewString(), appName: appName, level: zap.DebugLevel}
}

func NewLogger(config *LoggerConfig) *Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Level.SetLevel(config.level)
	logger, _ := loggerConfig.Build()
	logger = logger.With(zap.String("ref_id", config.refID), zap.String("app_name", config.appName))
	return &Logger{Logger: logger}
}

func (c *LoggerConfig) SetLevel(level string) {
	switch level {
	case LevelDebug:
		c.level = zap.DebugLevel
	case LevelInfo:
		c.level = zap.InfoLevel
	case LevelWarn:
		c.level = zap.WarnLevel
	case LevelError:
		c.level = zap.ErrorLevel
	case LevelFatal:
		c.level = zap.FatalLevel
	default:
		c.level = zap.InfoLevel
	}
}

func (c *LoggerConfig) SetAppName(appName string) {
	c.appName = appName
}

func (c *LoggerConfig) SetReference(refID string) {
	c.refID = refID
}

func (c *LoggerConfig) GetRefernce() string {
	return c.refID
}

func (c *LoggerConfig) GetAppName() string {
	return c.appName
}

func (c *LoggerConfig) GetLevel() zapcore.Level {
	return c.level
}

func UserIP(req *http.Request) string {
	ip, _, err := net.SplitHostPort(strings.TrimSpace(req.RemoteAddr))
	if err != nil {
		return ""
	}

	remoteIP := net.ParseIP(ip)
	if remoteIP == nil {
		return ""
	}

	ipHeader := req.Header.Get("X-Forwarded-For")
	ip, valid := validateHeader(ipHeader)
	if valid {
		return ip
	}

	ipHeader = req.Header.Get("X-Real-Ip")
	ip, valid = validateHeader(ipHeader)
	if valid {
		return ip
	}

	return remoteIP.String()
}

func validateHeader(header string) (clientIP string, valid bool) {
	if header == "" {
		return "", false
	}
	items := strings.Split(header, ",")
	for i, ipStr := range items {
		ipStr = strings.TrimSpace(ipStr)
		ip := net.ParseIP(ipStr)
		if ip == nil {
			return "", false
		}
		
		if i == 0 {
			clientIP = ipStr
			valid = true
		}
	}
	return
}
