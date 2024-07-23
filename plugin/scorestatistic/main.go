package scorestatistic

import (
	"fmt"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/ctxext"
	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
)

func init() {
	en := control.Register("scorestatistic", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "根据群名片统计分数实时排名",
		Help:             "命令 :score 或 ：score 或 分数实时排名",
	})

	en.OnFullMatchGroup([]string{":score", "：score", "分数实时排名"}, zero.OnlyGroup).
		//SetBlock(true).
		Limit(ctxext.LimitByGroup).
		Handle(func(ctx *zero.Ctx) {
			fmt.Println("检测到分数实时排名请求")
			cards := getMemberCards(ctx)
			msg := generateScoreAnalyse(cards)
			sendGroupImgMsgFromStr(msg, ctx)
		})

	en.OnCommandGroup([]string{"score2", "分数排名", "分数统计"}, zero.OnlyGroup).
		//SetBlock(true).
		Limit(ctxext.LimitByGroup).
		Handle(func(ctx *zero.Ctx) {
			fmt.Println("检测到分数实时排名请求")
			cards := getMemberCards(ctx)
			msg := generateScoreAnalyse(cards)
			sendGroupImgMsgFromStr(msg, ctx)
		})

}

func getMemberCards(ctx *zero.Ctx) (cards []string) {
	members := ctx.GetThisGroupMemberList()
	members.ForEach(func(k, v gjson.Result) bool {
		cards = append(cards, v.Get("card").String())
		return true
	})
	return cards
}
