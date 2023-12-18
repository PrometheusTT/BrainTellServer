package utils

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var pluginCommands = map[string][]string{
	"convert_Data_Type":                  {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/data_type/Convert_8_16_32_bits_data/libdatatypeconvert.so", "-f", "dtc", "-i", "", "-o", "", "-p 16"},
	"image_up_sample":                    {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_geometry/image_resample/libresampleimage.so", "-f", "up_sample", "-i", "", "-o", "", "-p 1 1 1 1"},
	"Grayscale_Image_Distance_Transform": {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Grayscale_Image_Distance_Transform/libgsdt.so", "-f", "gsdt", "-i", "", "-o", ""},
	"Fast_Distance_Transform":            {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Fast_Distance_Transform/libfast_dt.so", "-f", "dt", "-i", "", "-o", ""},
	"Edge_Extraction_from_Mask_Image":    {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_edge_detection/Edge_Extraction_from_Mask_Image/libedge_of_maskimg.so", "-f", "find_edge", "-i", "", "-o", ""},
	"extract_Z_Slices":                   {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_geometry/extract_Z_Slices/libextractZSlices.so", "-f", "subzslices2stack", "-i", "", "-o", ""},
	"Gaussian_Filter":                    {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Gaussian_Filter/libgaussianfilter.so", "-f", "gf", "-i", "", "-o", ""},
	"convert_file_format":                {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/data_IO/convert_file_format/libconvert_file_format.so", "-f", "convert_format", "-i", "", "-o", ""},
	"Cell_Segmentation_GVF":              {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_segmentation/Cell_Segmentation_GVF/libgvf_cellseg.so", "-f", "gvf_segmentation", "-i", "", "-o", ""},
	"split_extract_channels":             {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/color_channel/split_extract_channels/libchannelsplit.so", "-f", "split_extract_channels", "-i", "", "-o", ""},
	"5D_Stack_Converter":                 {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/data_type/5D_Stack_Converter/libmovieZCswitch.so", "-f", "4D_to_5D", "-i", "", "-o", ""},
	"Simple_Adaptive_Thresholding":       {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_thresholding/Simple_Adaptive_Thresholding/libada_threshold.so", "-f", "adath", "-i", "", "-o", ""},
	"median_filter":                      {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/median_filter/libmedianfilter.so", "-f", "fixed_window", "-i", "", "-o", ""},
	"muitiscaleEnhancement":              {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_analysis/muitiscaleEnhancement/libmultiscaleEnhancement.so", "-f", "adaptive_auto", "-i", "", "-o", ""},
	//"Fast_Distance_Transform":            {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Fast_Distance_Transform/libfast_dt.so", "-f", "dt ", "-i", "", "-o", ""},
	//"Fast_Distance_Transform":            {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Fast_Distance_Transform/libfast_dt.so", "-f", "dt ", "-i", "", "-o", ""},
	//"Fast_Distance_TransfZxSorm":            {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Fast_Distance_Transform/libfast_dt.so", "-f", "dt ", "-i", "", "-o", ""},
	//"Fast_Distance_Transform":            {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Fast_Distance_Transform/libfast_dt.so", "-f", "dt ", "-i", "", "-o", ""},
	//"Fast_Distance_Transform":            {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Fast_Distance_Transform/libfast_dt.so", "-f", "dt ", "-i", "", "-o", ""},
	//"Fast_Distance_Transform":            {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Fast_Distance_Transform/libfast_dt.so", "-f", "dt ", "-i", "", "-o", ""},
	//"Fast_Distance_Transform":            {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Fast_Distance_Transform/libfast_dt.so", "-f", "dt ", "-i", "", "-o", ""},
	//"Fast_Distance_Transform":            {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Fast_Distance_Transform/libfast_dt.so", "-f", "dt ", "-i", "", "-o", ""},
	//"Fast_Distance_Transform":            {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Fast_Distance_Transform/libfast_dt.so", "-f", "dt ", "-i", "", "-o", ""},
	//"Fast_Distance_Transform":            {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Fast_Distance_Transform/libfast_dt.so", "-f", "dt ", "-i", "", "-o", ""},
	//"Fast_Distance_Transform":            {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Fast_Distance_Transform/libfast_dt.so", "-f", "dt ", "-i", "", "-o", ""},
	//"Fast_Distance_Transform":            {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Fast_Distance_Transform/libfast_dt.so", "-f", "dt ", "-i", "", "-o", ""},
	//"Fast_Distance_Transform":            {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Fast_Distance_Transform/libfast_dt.so", "-f", "dt ", "-i", "", "-o", ""},
	//"Fast_Distance_Transform":            {"-x", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/plugins/image_filters/Fast_Distance_Transform/libfast_dt.so", "-f", "dt ", "-i", "", "-o", ""},
}

// CMD 命令行结构体
type CMD struct {
	CmdName string `json:"cmdname"`
	Args    string `json:"args"`
}

// Img 图像数据结构体
type Img struct {
	Obj string `json:"obj"`
}

func GetPluginList() (string, error) {
	var pluginNames []string

	// 遍历插件Commands的keys
	for key := range pluginCommands {
		pluginNames = append(pluginNames, key)
	}
	pluginList, err := json.Marshal(pluginNames)
	if err != nil {
		return "", err
	}
	return string(pluginList), err
}

func ListImageFiles() (string, error) {
	directory := "/home/BrainTellServer/PluginData"
	var imageFiles []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isImageFile(info.Name()) {
			imageFiles = append(imageFiles, path)
		}
		return nil
	})
	imageList, err := json.Marshal(imageFiles)
	if err != nil {
		return "", err
	}
	return string(imageList), err
}

func isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".v3draw" || ext == ".tiff" || ext == ".tif" || ext == ".v3dpbd"
}

func FindCmdArgs(pluginName string) ([]string, error) {
	if args, exists := pluginCommands[pluginName]; exists {
		fmt.Printf("Commands for %s: %v\n", pluginName, args)
		return args, nil
	} else {
		return nil, fmt.Errorf("plugin %s not found", pluginName)
	}
}

func ConvertImageType(arg string) (string, error) {
	//ctx := context.TODO()
	//if err := availableImgProcess.Acquire(ctx, 1); err != nil {
	//	log.Infof("Failed to acquire semaphore: %v\n", err)
	//	return "", errors.New(fmt.Sprintf("Failed to acquire semaphore: %v", err))
	//}
	//defer availableImgProcess.Release(1)
	//Arg := strings.Replace(arg, "\\", "", -1)
	input := arg
	ConvertimageSavefile := "/home/BrainTellServer/tmppath/" + fmt.Sprintf("%s_%d.v3draw", "convert_data_type", time.Now().UnixNano()) //转换后输出地址
	args, err := FindCmdArgs("convert_Data_Type")
	if err != nil {
		log.WithFields(
			log.Fields{
				"event":  "convert image type",
				"status": "Failed",
				"args":   []string(args),
			}).Warnf("%s\n", err)
		return "", err
	}
	args[5] = input
	args[7] = ConvertimageSavefile

	cmd := exec.Command("xvfb-run", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/Vaa3D-x")
	cmd.Args = append(cmd.Args, args...)

	fmt.Println(cmd.String())
	//返回环境变量集合
	cmd.Env = append(os.Environ())

	out, err := cmd.Output()

	if err != nil {
		log.WithFields(
			log.Fields{
				"event":  "convert image type",
				"status": "Failed",
				"out":    string(out),
			}).Warnf("%s\n", err)
		return "", err
	}
	log.WithFields(
		log.Fields{
			"event":  "convert image type",
			"status": "Success",
			"out":    string(out),
		}).Infof("\n")
	return ConvertimageSavefile, nil
}

func V3DCommandLine(cmdName string, arg string) (string, error) {
	//ctx := context.TODO()
	//if err := availableImgProcess.Acquire(ctx, 1); err != nil {
	//	log.Infof("Failed to acquire semaphore: %v\n", err)
	//	return "", errors.New(fmt.Sprintf("Failed to acquire semaphore: %v", err))
	//}
	//defer availableImgProcess.Release(1)

	input := arg
	saveFile := "/home/BrainTellServer/tmppath/" + fmt.Sprintf("%s_%d.v3draw", cmdName, time.Now().UnixNano()) //转换后输出地址
	args, err := FindCmdArgs(cmdName)
	if err != nil {
		log.WithFields(
			log.Fields{
				"event":  "use plugin",
				"status": "Failed",
				"args":   []string(args),
			}).Warnf("%s\n", err)
		return "", err
	}
	args[5] = input
	args[7] = saveFile

	cmd := exec.Command("xvfb-run", "/home/BrainTellServer/Vaa3D-x-Qt6.1.3-version/Vaa3D-x")
	cmd.Args = append(cmd.Args, args...)

	log.Infoln(cmd.String())
	//返回环境变量集合
	cmd.Env = append(os.Environ())

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.WithFields(
			log.Fields{
				"event":  "use plugin",
				"status": "command execution failed",
				"error":  err.Error(),
			}).Errorf("\n")
		return "", err

	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.WithFields(
			log.Fields{
				"event":  "use plugin",
				"status": "command execution failed",
				"error":  err.Error(),
			}).Errorf("\n")
		return "", err

	}

	if err := cmd.Start(); err != nil {
		log.WithFields(
			log.Fields{
				"event":  "use plugin",
				"status": "command execution failed",
				"error":  err.Error(),
			}).Errorf("\n")
		return "", err

	}

	stdoutData, _ := ioutil.ReadAll(stdout)
	stderrData, _ := ioutil.ReadAll(stderr)

	if err := cmd.Wait(); err != nil {
		log.WithFields(
			log.Fields{
				"event":  "use plugin",
				"status": "command execution failed",
				"stderr": string(stderrData),
				"error":  err.Error(),
			}).Errorf("\n")
		return "", err
	}

	//out, err := cmd.Output()
	//if err != nil {
	//	log.WithFields(
	//		log.Fields{
	//			"event":  "use plugin",
	//			"status": "Failed",
	//			//"out":    string(out),
	//		}).Warnf("%s\n", err)
	//	return "", err
	//}

	log.WithFields(
		log.Fields{
			"event":  "use plugin",
			"status": "Success",
			"stdout": string(stdoutData),
			"stderr": string(stderrData),
		}).Infof("\n")
	//time.Sleep(time.Duration(5) * time.Second)
	return saveFile, nil
}
