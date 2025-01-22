package luck

import (
    "image"
    "image/png"
    "image/draw"
	"github.com/fogleman/gg"
    "os"
    "log"
	"math/rand"
	"time"
	"strconv"
	"fmt"
)

func loadImage(filepath string) image.Image {
    file, err := os.Open(filepath)
    if err != nil {
        log.Fatalf("Failed to open file: %v", err)
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    if err != nil {
        log.Fatalf("Failed to decode image: %v", err)
    }
    return img
}

func get_time_seed() int64{
	now := time.Now()
	year, month, day := now.Date()
	seed := int64(year*10000 + int(month)*100 + day)
	return seed
}

var event_list = []string{
"打麻将",
"出勤",
"睡觉",
"打游戏",
"吃大餐",
"出去玩",
"访友",
"赶DDL",
"抽卡",
"听歌",
"看番",
"打代码"} 

var pic_num = 6

func Set_Pic_Num(num int) {
	pic_num = num
}

func Set_event_list(list []string) {
	event_list = list
}

func gen_luck(lid int) (string, string, string){
	id1 := rand.Intn(len(event_list))
	id2 := int((id1 + rand.Intn(len(event_list) - 1) + 1) % 12)
	good, bad := event_list[id1], event_list[id2]
	if lid < 2 {
		return "0.png", "万事皆宜", "行无禁忌"
	} else if lid < 4 {
		return "1.png", good, bad
	} else if lid < 6 {
		return "2.png", good, bad
	} else if lid < 8 {
		return "3.png", good, bad
	} else if lid < 11 {
		return "4.png", good, bad
	} else if lid < 12 {
		return "5.png", good, bad
	} else {
		return "6.png", "诸事不宜", "诸事不宜"
	}
}

func Gen_Pic(uid int64, bk_dir string, ft_dir string, ttf_path string, out_dir string) {
	rand.Seed(uid + get_time_seed())

	pid := strconv.Itoa(rand.Intn(pic_num) + 1) + ".png";
	lid, good, bad := gen_luck(rand.Intn(13))

	// fmt.Println(pid, lid)
	// fmt.Println(good, bad)

    background := loadImage(bk_dir + "/" + pid)
    overlay := loadImage(ft_dir + "/" + lid)

    output := image.NewRGBA(background.Bounds())

    draw.Draw(output, output.Bounds(), background, image.Point{0, 0}, draw.Src)

	offset := image.Pt(0, output.Bounds().Dy() - overlay.Bounds().Dy())
    draw.Draw(output, overlay.Bounds().Add(offset), overlay, image.Point{0, 0}, draw.Over)

	dc := gg.NewContextForImage(output)
	if err := dc.LoadFontFace(ttf_path, 36); err != nil {
		fmt.Println("加载字体失败:", err)
		return
	}

	dc.SetRGB(1, 0, 0)
	dc.DrawString("宜: " + good, 36, float64(output.Bounds().Dy() - 202))
	
	dc.SetRGB(0, 0, 0)
	dc.DrawString("忌: " + bad, 36, float64(output.Bounds().Dy() - 108))

    outFile, err := os.Create(out_dir + "/" +strconv.Itoa(int(uid))+".png")
    if err != nil {
        log.Fatalf("Failed to create output file: %v", err)
    }
    defer outFile.Close()

    if err := png.Encode(outFile, dc.Image()); err != nil {
        log.Fatalf("Failed to encode output image: %v", err)
    }
}