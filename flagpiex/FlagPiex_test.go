package flagpiex

import (
	"fmt"
	"github.com/lxn/win"
	"image/png"
	"os"
	"strconv"
	"testing"
	"time"
	"yys/getyyshwnd"
	"yys/yys_screenshot"
)

func Test_FlagPixel(t *testing.T){
	f :=new(FLagPiex)
	b :=f.FlagZhanDouJieMian()
	fmt.Println(b)
}


//截取当前图像大小方便取色
func Test_JieTuCaptureRect(T *testing.T){
	hd :=yys_screenshot.Yys_windows_screenshot{}
	img, _ := hd.YYS_Capture()//得到一个go类型的窗口句柄图像
	//ioutil.WriteFile("./output.png", img, 0666) //还原图像
	file, _ := os.Create(strconv.Itoa(int(time.Now().UnixNano()))+".png")
	png.Encode(file,img)
}

//手动取色
	func Test_find_Pixel(t *testing.T){
	//r :=yys_find_img.Result{}
	colorxy:=[][]int{
		{643,137},
		{947,137},
		{643,258},
		{947,258},
		{643,379},
		{947,379},
		{643,499},
		{947,499},
	}
	//r.Find_Pixel(xyp)

	hwnd := getyyshwnd.YYSHWND{}
	hdc :=win.GetDC(hwnd.Get_yys_hwnd())
	defer win.DeleteDC(hdc)
	for i,_ :=range colorxy{
		colorPixel :=win.GetPixel(hdc,int32(colorxy[i][0]),int32(colorxy[i][1]))
		fmt.Printf("{%d,%d,%d},\n",colorxy[i][0],colorxy[i][1],colorPixel)
		//c :=uint32(colorPixel)
		//if c != uint32(colorxy[i][2]) {
		//	return false
		//}
		//r.ColorrfeToRGB(c)
	}
}