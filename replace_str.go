package main

import (
    "strings"
    "fmt"
)

func main() {
    s1 := `---
    title: k3s 中 selflink 问题排查解决
    description: k8s、k3s 中动态创建pv、pvc时，Provisioner出现 selflink 错误解决方案。
    ---
    
    ## 问题描述
    
    在部署 statefulset 类型的工作负载时，动态创建 PV/PVC 是一种比较常用的配置方式，动态创建 PV/PVC 的方法基本如下：
    
    - 1、创建自己的 StorageClass 备用。
    - 2、创建 statefulset ，在 yaml 文件的 volumeClaimTemplates 块，添加 StorageClass 的名字。`

    new := `---
    title: k3s 中 selflink 问题排查解决
    description: k8s、k3s 中动态创建pv、pvc时，Provisioner出现 selflink 错误解决方案。
    toc: true
    authors: deanwu
    tags : [Kubernetes, k3s]
    categories: [Kubernetes, ]
    date: '2022-02-14'
    lastmod: '2022-02-14'
    draft: false
    url: /posts/2022-02-14-k8s-issue-selflink.html 
---`

    old := `---
    title: k3s 中 selflink 问题排查解决
    description: k8s、k3s 中动态创建pv、pvc时，Provisioner出现 selflink 错误解决方案。
    ---`


    s := strings.Replace(s1, old, new, -1)
    fmt.Println(s)
}
