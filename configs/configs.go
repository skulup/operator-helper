package configs

import (
	"fmt"
	"github.com/go-logr/logr"
	"github.com/skulup/operator-helper/util"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"log"
	"os"
	"path/filepath"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"strings"
	"sync"
)

var logger logr.Logger
var loggerOnce sync.Once

var envEnableWebHooks = "ENABLE_WEBHOOKS"
var envWebHookCertificateDir = "WEBHOOK_CERTIFICATES_DIR"
var envNamespacesToWatch = "NAMESPACES_TO_WATCH"
var envEnableLeaderElection = "ENABLE_LEADER_ELECTION"
var envLeaderElectionNamespace = "LEADER_ELECTION_NAMESPACE"

// GetLogger get the logger instance to use
func GetLogger(operatorName string, opts ...zap.Opts) logr.Logger {
	loggerOnce.Do(func() {
		if len(opts) == 0 {
			opts = append(opts, zap.UseDevMode(true))
		}
		logger = zap.New(opts...).WithName(operatorName)
		ctrl.SetLogger(logger)
	})
	return logger
}

// GetManagerParams get the manager options to use
func GetManagerParams(scheme *runtime.Scheme, operatorName, domainName string) (*rest.Config, ctrl.Options) {
	options := ctrl.Options{
		Scheme:                  scheme,
		Port:                    9443,
		MetricsBindAddress:      "",
		Logger:                  GetLogger(operatorName),
		LeaderElection:          LeaderElectionEnabled(),
		LeaderElectionNamespace: LeaderElectionNamespace(operatorName),
		LeaderElectionID:        fmt.Sprintf("leader-lock-65403bab.%s.%s", operatorName, domainName),
	}
	namespaces := NamespacesToWatch()
	if len(namespaces) == 0 {
		// watch all namespaces
		options.Namespace = ""
	} else if len(namespaces) == 1 {
		options.Namespace = namespaces[0]
	} else {
		options.NewCache = cache.MultiNamespacedCacheBuilder(namespaces)
	}
	options.CertDir = GetWebHookCertDir()
	return ctrl.GetConfigOrDie(), options
}

// LeaderElectionEnabled checks if leader election is enabled
func LeaderElectionEnabled() bool {
	return strings.TrimSpace(os.Getenv(envEnableLeaderElection)) != "false"
}

// WebHooksEnabled checks if webhook is enabled
func WebHooksEnabled() bool {
	if strings.TrimSpace(os.Getenv(envEnableWebHooks)) != "false" {
		if _, err := os.Stat(GetWebHookCertDir()); !os.IsNotExist(err) {
			return true
		}
		log.Printf("The webhook cert directory does not exists: %s", GetWebHookCertDir())
	}
	return false
}

// LeaderElectionNamespace get the leader election namespace
func LeaderElectionNamespace(operatorName string) string {
	ns := strings.TrimSpace(os.Getenv(envLeaderElectionNamespace))
	if ns != "" {
		return ns
	}
	return operatorName
}

// NamespacesToWatch get the array of namespaces to watch
func NamespacesToWatch() []string {
	val := os.Getenv(envNamespacesToWatch)
	if val == "" {
		return []string{}
	}
	namespaces := strings.Split(val, ",")
	for i, n := range namespaces {
		// cleanup
		namespaces[i] = strings.TrimSpace(n)
	}
	return namespaces
}

// GetWebHookCertDir returns the directory of the webhook certificates
func GetWebHookCertDir() string {
	def := filepath.Join(os.TempDir(), "k8s-webhook-server", "serving-certs")
	return util.ValueOr(envWebHookCertificateDir, def)
}
