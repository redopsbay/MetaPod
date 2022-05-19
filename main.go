package main

import (
    "net/http"
    "log"
    "github.com/gin-gonic/gin"
    "time"
    "os"
)

type podMeta struct {
	Namespace string `json:"Namespace"`
    PodIp string `json:"PodIp"`
    NodeName string `json:"NodeName"`
    PodName string `json:"PodName"`
    CpuReq string `json:"CpuRequests"`
    CpuLimit string `json:"CpuLimits"`
    MemReq string `json:"MemoryRequests"`
    MemLimit string `json:"MemoryLimit"`
    ServiceAccount string `json:"ServiceAccount"`
    ClusterName string `json:"ClusterName"`
}

func getMetaData(c *gin.Context) {
    curTime := time.Now()
    log.Println(curTime, "===> Get Pod MetaData", http.StatusOK)
    var _Podinfo *podMeta
    _Podinfo = new(podMeta)
    *_Podinfo = getEnv(_Podinfo)
    c.HTML(http.StatusOK, "index.tmpl", _Podinfo)
}

func getEnv(meta *podMeta) (metadata podMeta) {
    meta = new(podMeta)
    (*meta).Namespace = os.Getenv("POD_NAMESPACE")
    (*meta).PodIp = os.Getenv("POD_IP")
    (*meta).NodeName = os.Getenv("NODE_NAME")
    (*meta).PodName = os.Getenv("POD_NAME")
    (*meta).CpuReq = os.Getenv("POD_CPU_REQUEST")
    (*meta).CpuLimit = os.Getenv("POD_CPU_LIMIT")
    (*meta).MemReq = os.Getenv("POD_MEM_REQUEST")
    (*meta).MemLimit = os.Getenv("POD_MEM_LIMIT")
    (*meta).ServiceAccount = os.Getenv("POD_SERVICE_ACCOUNT")
    (*meta).ClusterName = os.Getenv("CLUSTER_NAME")
    metadata = *meta
    return
}

func main() {
    router := gin.Default()
    router.Static("/static", "./templates/")
    router.LoadHTMLGlob("templates/*")
    router.GET("/metapod", getMetaData)
    router.Run("0.0.0.0:8080")
}