package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/aobco/log"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"time"
	"xsky.com/ocpf/pkg/k8s"
)

func main() {
	// 64.86
	clusterInfo := base64.StdEncoding.EncodeToString([]byte("apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJmRENDQVNHZ0F3SUJBZ0lCQURBS0JnZ3Foa2pPUFFRREFqQWtNU0l3SUFZRFZRUUREQmx5YTJVeUxYTmwKY25abGNpMWpZVUF4TnpRMU9EQTJNVGd6TUNBWERUSTFNRFF5T0RBeU1EazBNMW9ZRHpJeE1qVXdOREEwTURJdwpPVFF6V2pBa01TSXdJQVlEVlFRRERCbHlhMlV5TFhObGNuWmxjaTFqWVVBeE56UTFPREEyTVRnek1Ga3dFd1lICktvWkl6ajBDQVFZSUtvWkl6ajBEQVFjRFFnQUVJU282Qzg2WUFwalpnTFpRZ1pVOCtXRWE3b2o0NWwvaklTOGcKNFpleFhvNjJmTlJKR09nL0NCbzhxRDJlS1pONjZVTmtQaHBMZm1WeEc2Z0ZOcnpubUtOQ01FQXdEZ1lEVlIwUApBUUgvQkFRREFnS2tNQThHQTFVZEV3RUIvd1FGTUFNQkFmOHdIUVlEVlIwT0JCWUVGR0FmSG51ZHExUHB1cVNLCmNrS0x2cGVnZ25odk1Bb0dDQ3FHU000OUJBTUNBMGtBTUVZQ0lRRHFYU0duZjZqRlJnYjFWWE1KS3FMRjgvekIKQXVEcm52Z1pmWHcwZklZaTJ3SWhBTnRuc1lFaFQ3R3drYlhpS1U1REVVdXFTWWJjNytLSld3enN3eGZmWjNSRwotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    server: https://10.16.64.86:6443\n  name: default\ncontexts:\n- context:\n    cluster: default\n    user: default\n  name: default\ncurrent-context: default\nkind: Config\npreferences: {}\nusers:\n- name: default\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJsRENDQVRxZ0F3SUJBZ0lJWWtWM2daNjN3Tmd3Q2dZSUtvWkl6ajBFQXdJd0pERWlNQ0FHQTFVRUF3d1oKY210bE1pMWpiR2xsYm5RdFkyRkFNVGMwTlRnd05qRTRNekFnRncweU5UQTBNamd3TWpBNU5ETmFHQTh5TVRJMQpNRFF3TkRBeU1EazBNMW93TURFWE1CVUdBMVVFQ2hNT2MzbHpkR1Z0T20xaGMzUmxjbk14RlRBVEJnTlZCQU1UCkRITjVjM1JsYlRwaFpHMXBiakJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUFCTGlQTDU1WjFvZFMKZFdtMkh3b1ZMSGdIRzNOemtndHljcmFaL25zMHBIVFMwSEtoRExGRnVKU2FFa0tBbmVCQXUvZEFTaVZGWlR4TQpDeVZwbTNLWFdEbWpTREJHTUE0R0ExVWREd0VCL3dRRUF3SUZvREFUQmdOVkhTVUVEREFLQmdnckJnRUZCUWNECkFqQWZCZ05WSFNNRUdEQVdnQlNyUjJzSDJ3UDlueVJuUCtidWpiZERaVElNVVRBS0JnZ3Foa2pPUFFRREFnTkkKQURCRkFpRUFqbUpBSWdYZ3RJRld5ZGIrTEJMMkhtdGFmRlJhanc3VVp4OUptNHExMkM0Q0lHL3Z6MTFuQ1dNbAp2VVZYckVTd09sVjZUY1M0aEpxeFQveFFscnlZRlROQwotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tQkVHSU4gQ0VSVElGSUNBVEUtLS0tLQpNSUlCZkRDQ0FTR2dBd0lCQWdJQkFEQUtCZ2dxaGtqT1BRUURBakFrTVNJd0lBWURWUVFEREJseWEyVXlMV05zCmFXVnVkQzFqWVVBeE56UTFPREEyTVRnek1DQVhEVEkxTURReU9EQXlNRGswTTFvWUR6SXhNalV3TkRBME1ESXcKT1RReldqQWtNU0l3SUFZRFZRUUREQmx5YTJVeUxXTnNhV1Z1ZEMxallVQXhOelExT0RBMk1UZ3pNRmt3RXdZSApLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFaHI4ZEV3eURPN05BZEVRTFNTMHJHWEcxSGRUSFcwVklEM2lEClgyWTJ5M0RVR29pWVBUamRZOWZld0Q1UHg0SVFXT2s0YkV6R01OZnR6cEJ1RVZhYS82TkNNRUF3RGdZRFZSMFAKQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCL3dRRk1BTUJBZjh3SFFZRFZSME9CQllFRkt0SGF3ZmJBLzJmSkdjLwo1dTZOdDBObE1neFJNQW9HQ0NxR1NNNDlCQU1DQTBrQU1FWUNJUURkeTJlbnIwRzgzQ0x5dFYveC8yWFM1Y2FyCko5WGF0blVBVkN3UlpFNWZRQUloQUtTWmpsNldXdTBNNGtxNWlBWUl1ekRneHBSdXpoN0YzcmZSd0RTa0p3MWEKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=\n    client-key-data: LS0tLS1CRUdJTiBFQyBQUklWQVRFIEtFWS0tLS0tCk1IY0NBUUVFSUZpMHdndjJ3RlBGQTl4UFZtUWl5VU1HdlRyaHhpaHZEY1ZSd1lrb2JTTjRvQW9HQ0NxR1NNNDkKQXdFSG9VUURRZ0FFdUk4dm5sbldoMUoxYWJZZkNoVXNlQWNiYzNPU0MzSnl0cG4rZXpTa2ROTFFjcUVNc1VXNApsSm9TUW9DZDRFQzc5MEJLSlVWbFBFd0xKV21iY3BkWU9RPT0KLS0tLS1FTkQgRUMgUFJJVkFURSBLRVktLS0tLQo="))
	// 113.14
	//clusterInfo := base64.StdEncoding.EncodeToString([]byte("apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJlekNDQVNHZ0F3SUJBZ0lCQURBS0JnZ3Foa2pPUFFRREFqQWtNU0l3SUFZRFZRUUREQmx5YTJVeUxYTmwKY25abGNpMWpZVUF4TnpNMk5qazFNRFUyTUNBWERUSTFNREV4TWpFMU1UY3pObG9ZRHpJeE1qUXhNakU1TVRVeApOek0yV2pBa01TSXdJQVlEVlFRRERCbHlhMlV5TFhObGNuWmxjaTFqWVVBeE56TTJOamsxTURVMk1Ga3dFd1lICktvWkl6ajBDQVFZSUtvWkl6ajBEQVFjRFFnQUVaVGM4amlRUGVVckZiekdMbkNsL1grTjhRd2orbzVRc1hQVkQKd2JEUlBoVmJ6WnJ3WXRYOW8vanloTFRoVG1Qd1I5cjBkRGVMc3hncnpySHJBRzdvd0tOQ01FQXdEZ1lEVlIwUApBUUgvQkFRREFnS2tNQThHQTFVZEV3RUIvd1FGTUFNQkFmOHdIUVlEVlIwT0JCWUVGT1BoditPQ2hja3d3a3k0CmpheC9WbGZTZ0EweE1Bb0dDQ3FHU000OUJBTUNBMGdBTUVVQ0lIdlZYbGxjbUZ3UUNxbkNWTWNoQU1ySk1peEwKVkRSVUNWM0xwdUx1MVBwbkFpRUFvcHk0a3pqOWluT0tCaExNQW1wWVdlY0JFcDVHbVI0aVByRnplaEVnWGhRPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    server: https://10.16.113.14:6443\n  name: default\ncontexts:\n- context:\n    cluster: default\n    user: default\n  name: default\ncurrent-context: default\nkind: Config\npreferences: {}\nusers:\n- name: default\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJsRENDQVRxZ0F3SUJBZ0lJVU5GUTdiNTNtT3d3Q2dZSUtvWkl6ajBFQXdJd0pERWlNQ0FHQTFVRUF3d1oKY210bE1pMWpiR2xsYm5RdFkyRkFNVGN6TmpZNU5UQTFOakFnRncweU5UQXhNVEl4TlRFM016WmFHQTh5TVRJMApNVEl4T1RFMU1UY3pObG93TURFWE1CVUdBMVVFQ2hNT2MzbHpkR1Z0T20xaGMzUmxjbk14RlRBVEJnTlZCQU1UCkRITjVjM1JsYlRwaFpHMXBiakJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUFCUEZhWG4zOE5oeXEKSjRwMWxBZlVzWWpYOVIyQS9ZbXA0dDg1YUpBT0k2UVJNVUV1SWtvT0xZUXV6K1lKZmxncFdjaW9RR0FvcXNZdAo2RjBHQkM5RWhEdWpTREJHTUE0R0ExVWREd0VCL3dRRUF3SUZvREFUQmdOVkhTVUVEREFLQmdnckJnRUZCUWNECkFqQWZCZ05WSFNNRUdEQVdnQlROYzdrNEl1MFNOWjFwWEc2YUZCT0xlaXFGVkRBS0JnZ3Foa2pPUFFRREFnTkkKQURCRkFpQkhkanNGVitNaWhnQWkzWWFvMHhRcDdIbllXaGxTVVFzQnp1YkhRc3hOaVFJaEFJb1V3cXNVMFFaVwp5dUhQMGlodC80SVNaVnM2MU5SSWJWRm1vVTFJMDdHdQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tQkVHSU4gQ0VSVElGSUNBVEUtLS0tLQpNSUlCZkRDQ0FTR2dBd0lCQWdJQkFEQUtCZ2dxaGtqT1BRUURBakFrTVNJd0lBWURWUVFEREJseWEyVXlMV05zCmFXVnVkQzFqWVVBeE56TTJOamsxTURVMk1DQVhEVEkxTURFeE1qRTFNVGN6TmxvWUR6SXhNalF4TWpFNU1UVXgKTnpNMldqQWtNU0l3SUFZRFZRUUREQmx5YTJVeUxXTnNhV1Z1ZEMxallVQXhOek0yTmprMU1EVTJNRmt3RXdZSApLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFMXhMc0QwVmNnS2x5c2N2cGIzWmFaaDZzYnJRamFHUExVdm1FClhhV3pUbjB5SFk1aFplbkpXbTIxTUFOL0Z1dzQ0MnhpVjFlV2J4bjhZVGE4eDVhb3lhTkNNRUF3RGdZRFZSMFAKQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCL3dRRk1BTUJBZjh3SFFZRFZSME9CQllFRk0xenVUZ2k3UkkxbldsYwpicG9VRTR0NktvVlVNQW9HQ0NxR1NNNDlCQU1DQTBrQU1FWUNJUURiS2ExNUROVWtiT1JkUTdYaTFyQlVtRkhFCktwd29uUjZQMVB5K2dSMkQ2Z0loQU05WG1zSFR3TDBhdFo1b2FxcGJlUklwMUdqTFAxSkJyQ2p3OHpRMHNlR2QKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=\n    client-key-data: LS0tLS1CRUdJTiBFQyBQUklWQVRFIEtFWS0tLS0tCk1IY0NBUUVFSUpiTExNa290MlJwRmdXZUYva01VaDcxS1FDVlNTWnFVbEdYTnRkWEU3NExvQW9HQ0NxR1NNNDkKQXdFSG9VUURRZ0FFOFZwZWZmdzJIS29uaW5XVUI5U3hpTmYxSFlEOWlhbmkzemxva0E0anBCRXhRUzRpU2c0dApoQzdQNWdsK1dDbFp5S2hBWUNpcXhpM29YUVlFTDBTRU93PT0KLS0tLS1FTkQgRUMgUFJJVkFURSBLRVktLS0tLQo="))
	k8sCfg, err := base64.StdEncoding.DecodeString(clusterInfo)
	if err != nil {
		log.Fatalf("decode error, %v", err)
	}
	k8s.Init(k8sCfg)
	podList, err := k8s.K8sClient.ListDeploymentPods("ocpf", "portal")
	if err != nil {
		log.Fatalf("list pods error, %v", err)
	}
	podIps := make([]string, 0)
	for _, pod := range podList.Items {
		if pod.Status.PodIP != "" {
			podIps = append(podIps, pod.Status.PodIP)
		}
	}
	log.Infof("pod ips: %v", podIps)
	watchPodIPs("ocpf", k8s.K8sClient.ClientSet)
}

func watchPodIPs(namespace string, clientset *kubernetes.Clientset) {
	// 创建Watcher
	watcher, err := clientset.CoreV1().Pods(namespace).Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("\nStarting to watch Pod IP changes in namespace %s...\n", namespace)

	// 处理事件
	for event := range watcher.ResultChan() {
		pod, ok := event.Object.(*corev1.Pod)
		if !ok {
			continue
		}

		switch event.Type {
		case "ADDED":
			if pod.Status.PodIP != "" {
				fmt.Printf("[%s] Pod Added: %s, IP: %s\n",
					time.Now().Format(time.RFC3339), pod.Name, pod.Status.PodIP)
			}
		case "MODIFIED":
			// 需要比较旧状态和新状态来确定IP是否变化
			// 在实际应用中，你可能需要维护一个本地缓存来比较
			fmt.Printf("[%s] Pod Modified: %s, IP: %s\n",
				time.Now().Format(time.RFC3339), pod.Name, pod.Status.PodIP)
		case "DELETED":
			fmt.Printf("[%s] Pod Deleted: %s\n",
				time.Now().Format(time.RFC3339), pod.Name)
		}
	}
}
