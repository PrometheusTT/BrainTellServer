package services

//
//import (
//	"fmt"
//	"github.com/pytorch/tensor"
//	"github.com/pytorch/torch"
//	log "github.com/sirupsen/logrus"
//)
//
//func Use3DUNet() {
//	// 加载PyTorch模型
//	modelPath := "path/to/your/model.pth"
//	model := torch.NewModel(modelPath)
//
//	// 创建一个输入张量
//	input := tensor.New(
//		tensor.Of(tensor.Float64),        // 根据模型要求的数据类型选择
//		tensor.WithShape(1, 3, 224, 224), // 根据模型期望的输入形状选择
//	)
//
//	// 执行推理
//	output, err := model.Forward(input)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 处理输出
//	// 这里根据你的模型和任务来处理输出张量
//	fmt.Println("Model output:", output)
//}
