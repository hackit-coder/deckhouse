package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func recordMetrics() {
	go func() {
		for {
			// Init kubeClient
			config, err := rest.InClusterConfig()
			if err != nil {
				log.Fatalf("failed to init kubeClient config. Error: %v", err)
			}
			kubeClient, err := kubernetes.NewForConfig(config)
			if err != nil {
				log.Fatalf("Failed to init kubeClient. Error: %v", err)
			}
			log.Infof("kubeClient successfully inited")
			// подключение к kube
			// получение списка ns нужной аннотацией
			// построение extended_monitoring_enabled

			opsProcessed.Inc()
			time.Sleep(10 * 60 * time.Second)
		}
	}()
}

var (
	//    540 extended_monitoring_pod_threshold
	//     90 extended_monitoring_pod_enabled
	//     52 extended_monitoring_ingress_threshold
	//     45 extended_monitoring_deployment_threshold
	//     45 extended_monitoring_deployment_enabled
	//     26 extended_monitoring_ingress_enabled
	//     22 extended_monitoring_enabled
	//     18 extended_monitoring_node_threshold
	//     14 extended_monitoring_daemonset_threshold
	//     14 extended_monitoring_daemonset_enabled
	//      3 extended_monitoring_statefulset_threshold
	//      3 extended_monitoring_statefulset_enabled
	//      3 extended_monitoring_node_enabled

	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
	EXTENDED_MONITORING_LABEL_THRESHOLD_PREFIX = "threshold.extended-monitoring.deckhouse.io/"
	EXTENDED_MONITORING_ENABLED_LABEL          = "extended-monitoring.deckhouse.io/enabled"
)

func main() {
	r := prometheus.NewRegistry()
	//r.MustRegister(opsProcessed)
	handler := promhttp.HandlerFor(
		r,
		promhttp.HandlerOpts{
			EnableOpenMetrics: false,
		})
	recordMetrics()
	http.Handle("/metrics", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8081", nil))
}
