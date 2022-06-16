package devicestatus

import (
	"net/http"
	"rc/param"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
)

//const unsigned int* map_bram_ctrl_address = spdcpldop_init();

//#cgo CFLAGS: -I. -I./include
//#cgo LDFLAGS: -L./lib -lmc3s029zesensorinfoget -lmc3s028zecpldcfg
//#include "libmc3s029zesensorinfoget.h"
//#include "libmc3s028zecpldcfg.h"
import (
	"C"
)

func DeviceStatusHandler(c *gin.Context) {
	voltage_data_tmp := float32(C.get_main_control_board_voltage1_vccaux())
	temperature_data_tmp := float32(C.get_main_control_board_temperature())
	total_memory := uint32(C.get_main_control_board_total_memory())
	free_memory := uint32(C.get_main_control_board_free_memory())
	memory_usage := uint32(float32(total_memory-free_memory) / float32(total_memory) * 100)
	//cpu_usage := int32(C.get_main_control_board_cpu_usage())
	percent, _ := cpu.Percent(time.Second, false)
	cpu_usage := int32(percent[0])
	devicedata := param.DeviceData{
		Voltage:     voltage_data_tmp,
		Temperature: temperature_data_tmp,
		MemoryUsage: memory_usage,
		CpuUsage:    cpu_usage,
	}
	c.JSON(http.StatusOK, devicedata)
}

func DeviceFPGAStatusHandler(c *gin.Context) {
	voltage_fpga_data_tmp := float32(C.get_pretreatment_board_voltage1(C.spdcpldop_init()))
	temperature_fpga_data_tmp := float32(C.get_pretreatment_board_temperature1(C.spdcpldop_init()))
	devicedata := param.DeviceFPGAData{
		Voltage:     voltage_fpga_data_tmp,
		Temperature: temperature_fpga_data_tmp,
	}
	c.JSON(http.StatusOK, devicedata)
}

func DeviceDSPStatusHandler(c *gin.Context) {
	voltage_fpga_data_tmp := float32(C.get_signal_processing_board_voltage1(C.spdcpldop_init()))
	temperature_fpga_data_tmp := float32(C.get_signal_processing_board_temperature1(C.spdcpldop_init()))
	devicedata := param.DeviceFPGAData{
		Voltage:     voltage_fpga_data_tmp,
		Temperature: temperature_fpga_data_tmp,
	}
	c.JSON(http.StatusOK, devicedata)
}
