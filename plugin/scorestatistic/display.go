package scorestatistic

import (
	"bytes"
	"image"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

// 懒得写配置文件读取了，先写死
var m = config{
	mConfig: mConfig{
		displayPicture:       true,
		tailPictureInPicture: "./data/kaoyanscore/qq_group.png",
		fontPath:             "./data/kaoyanscore/DingTalk JinBuTi.ttf",
		headMsgInWebserver:   "欢迎加入知识图谱与类脑智能实验室~\n官网：http://kglab.hdu.edu.cn \n群号：438609367",
	},
}

type config struct {
	mConfig
	lastUpdateTime     time.Time
	lastTriggerTimeMap map[int64]time.Time // 每个群最后一次主动触发时间，用来限流
}

type mConfig struct {
	displayPicture       bool   // 是否将分析结果转换为图片发到群里（和webserver可以同时开启）
	tailPictureInPicture string // # 缀在分析结果图片后面的图片，可以放群二维码，留空则不缀图
	fontPath             string // 字体文件路径
	headMsgInWebserver   string // 在详细的统计信息后面附加的内容（位于分数段总体统计和过密分数段分析之间）
}

// 把字符串消息转成图片发出去
func sendGroupImgMsgFromStr(msg string, ctx *zero.Ctx) {
	var buf bytes.Buffer
	var tailPicture image.Image
	if m.tailPictureInPicture != "" {
		f, err := os.Open(m.tailPictureInPicture)
		defer f.Close()
		if err != nil {
			logrus.WithError(err).Error("open tail picture failed")
		} else {
			tailPicture, _, err = image.Decode(f)
			if err != nil {
				logrus.WithError(err).Error("decode tail picture failed")
			}
		}
	}
	if tailPicture != nil {
		err := String2PicWriterWithTailPicture(msg, m.fontPath, &buf, tailPicture)
		if err != nil {
			logrus.WithError(err).Error("write pic with tail picture failed")
		}
	} else {
		err := String2PicWriter(msg, m.fontPath, &buf)
		if err != nil {
			logrus.WithError(err).Error("write pic failed")
		}
	}
	ctx.SendGroupMessage(ctx.Event.GroupID, message.ImageBytes(buf.Bytes()))
}
