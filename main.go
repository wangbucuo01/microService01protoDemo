package main

import (
	"fmt"
	study "goroutineAllProject/protoDemo/proto"

	"github.com/golang/protobuf/proto"
)

func main() {
	// 首先初始化proto中的消息
	study_info := &study.StudyInfo{}

	study_info.Id = 1
	study_info.Name = "protobuf"
	study_info.Duration = 180

	score := make(map[string]int32)
	score["实战"] = 100
	study_info.Score = score

	// 用字符串的方式打印protobuf消息
	fmt.Printf("字符串输出结果：%v\n", study_info.String())

	// 转成二进制文件
	marshal, _ := proto.Marshal(study_info)
	fmt.Printf("marshal result: %v\n", marshal)

	// 将二进制文件转成结构体
	newStudyInfo := study.StudyInfo{}
	_ = proto.Unmarshal(marshal, &newStudyInfo)
	fmt.Printf("unmarshal result: %v\n", &newStudyInfo)
}
