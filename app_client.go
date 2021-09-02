import {
    "context"
    "fmt"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/clientcmd"
}

func main() {
    rules := clientcmd.NewDefaultClientConfigLoadingRules()
    kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
    config, err := kubeconfig.ClientConfig()
    if err != nil {
        panic(err)
    }
    clientset := kubernetes.NewForConfigOrDie(config)

    nodeList, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
    if err != nil {
        panic(err)
    }

    for _, n := range nodeList.Items {
        fmt.Println(n.Name)
    }

    newPod := &corev1.Pod {
        ObjectMeta: metav1.ObjectMeta {
            Name: "test-pod",
        },
        Spec: corev1.PodSpec{
            Containers: []corev1.Container{
                {Name: "busy-box", Image: "busy-box:latest", Command: []string{"sleep", "100000"}},
            },
        },
    }

    pod, err := clientset.CoreV1().Pods("default").Create(context.Background(), newPod, metav1.CreateOptions{})
    if err != nil {
        panic(err)
    }
    fmt.Println(pod)
}