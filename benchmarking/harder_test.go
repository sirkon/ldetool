package main

import (
	"testing"

	"bytes"

	"regexp"

	"strconv"

	"github.com/stretchr/testify/require"
)

var sampleLines = []string{
	"[22653 22.09.2016 23:53:51] I MCHAT-ST txn_finalize(): TXN  0x357f4a000[674459048.ModAnketa() " +
		"reqid '1474577512273-1474577630:51#modChatAlpha.674459048@chat.agent' from " +
		"3:666366008(666366008@uin.icq) orgn 51:1474577512273-1474577630:51#modChatAlpha.674459048@chat.agent:]: " +
		"DONE (200[]) FLAGS[set:.......J.., unset:..........] " +
		"FIELDS[changed:.............] ank_ver[{6333262700747542897=22.09.16-23:53:51-189387121}] " +
		"list_ver[{6333246817957146840=22.09.16-22:52:13-188051672}] " +
		"name[🔥👑الصياد 👑خط احمر🚫❌🔫💣🔥] stamp[AoLBzdsoJtJ1LI6-RgY] regions[KW] " +
		"flags[0x10b(PLC....J..)] created[1474573933=22.09.16-22:52:13] creator[3:666366008] " +
		"avatars_lastcheck[1474577584] cavatar_lastmod[1474575351] " +
		"origin[0::] abuse_drugs[0] abuse_spam[0] abuse_porno[0] abuse_unknown[0] abuse_unknown[0]",
	"[22653 22.09.2016 23:53:51] I MCHAT-ST txn_finalize(): TXN  0x357f4a000[674459048.ModAnketa() " +
		"reqid '1474577512273-1474577630:51#modChatAlpha.674459048@chat.agent' from " +
		"3:666366008(666366008@uin.icq) orgn 51:1474577512273-1474577630:51#modChatAlpha.674459048@chat.agent:]: " +
		"DONE (200[]) FLAGS[set:.......J.., unset:..........] " +
		"FIELDS[changed:.............] ank_ver[{6333262700747542897=22.09.16-23:53:51-189387121}] " +
		"list_ver[{6333246817957146840=22.09.16-22:52:13-188051672}] " +
		"name[🔥👑الصياد 👑خط احمر🚫❌🔫💣🔥] stamp[AoLBzdsoJtJ1LI6-RgY] " +
		"flags[0x10b(PLC....J..)] created[1474573933=22.09.16-22:52:13] creator[3:666366008] " +
		"avatars_lastcheck[1474577584] cavatar_lastmod[1474575351] " +
		"origin[0::] abuse_drugs[0] abuse_spam[0] abuse_porno[0] abuse_unknown[0] abuse_unknown[0]",
	"[22653 22.09.2016 14:30:16] I MCHAT-ST txn_finalize(): TXN  0x1d2c42000[674451730.ModAnketa() " +
		"reqid '80405607-1474543815:1#modChatAlpha.674451730@chat.agent' " +
		"from 3:700337623(700337623@uin.icq) orgn 1:80405607-1474543815:1#modChatAlpha.674451730@chat.agent:]: " +
		"DONE (200[]) FLAGS[set:.........., unset:..........] FIELDS[changed:N............] " +
		"ank_ver[{6333117466407183233=22.09.16-14:30:16-168141697}] " +
		"list_ver[{6333107832793820511=22.09.16-13:52:53-166423903}] name[Việt Nam Zalo] " +
		"about[người Việt Nam] rules[ nói chuyện và chat video] " +
		"stamp[AoLBzaISFnNsHgQiXzQ] regions[US] flags[0x20b(PLC.....R.)] " +
		"created[1474541573=22.09.16-13:52:53] creator[3:700337623] " +
		"avatars_lastcheck[1474541573] origin[0::] " +
		"abuse_drugs[0] abuse_spam[0] abuse_porno[0] abuse_unknown[0] abuse_unknown[0]",
	`[3988519 30.05.2017 16:41:31] W MCHAT-ST txn_finalize(): TXN 0x1234d000[2000023207.ModAnketa() reqid ` +
		`'2837303925-1496151691:13#modChatAlpha.2000023207@chat.agent' from 3:1999585731(1999585731@uin.icq) ` +
		`orgn 13:2837303925-1496151691:13#modChatAlpha.2000023207@chat.agent:]: DONE (200[]) ` +
		`FLAGS[set:..........., unset:...........] FIELDS[changed:..............] ` +
		`ank_ver[{6425922582700258264=30.05.17-16:41:31-160728}] name[Cccc] about[Jdjdjdjdjjdr] nick[wwww] ` +
		`stamp[Aoe5190nWD89Ef-RCGQ] regions[RU] flags[0xb(PLC.......O)] ` +
		`created[1496151654=30.05.17-16:40:54] creator[3:1999585731] avatars_lastcheck[1496151655] ` +
		`cavatar_lastmod[1496151655] origin[0::] abuse drugs[0] abuse spam[0] abuse porno[0]` +
		`abuse violation[0] abuse other[0]`,
}
var lines [][]byte
var names = [][]byte{
	[]byte("🔥👑الصياد 👑خط احمر🚫❌🔫💣🔥"),
	[]byte("🔥👑الصياد 👑خط احمر🚫❌🔫💣🔥"),
	[]byte("Việt Nam Zalo"),
	[]byte("Cccc"),
}
var avatarsLastcheck = []int64{
	1474577584,
	1474577584,
	1474541573,
	1496151655,
}
var avatarsLastcheckLit = [][]byte{}

func init() {
	totalLength := 0
	for _, line := range sampleLines {
		totalLength += len(line)
	}
	buf := make([]byte, totalLength)
	offset := 0
	for _, line := range sampleLines {
		copy(buf[offset:], line)
		next := offset + len(line)
		lines = append(lines, buf[offset:next])
		offset = next
	}
	for _, num := range avatarsLastcheck {
		data := strconv.FormatInt(num, 10)
		avatarsLastcheckLit = append(avatarsLastcheckLit, []byte(data))
	}
}

var ldeParser = &CRMod{}

func BenchmarkLDEComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j, line := range lines {
			if ok, err := ldeParser.Extract(line); !ok {
				if err != nil {
					b.Fatalf("%s on parsing >>\033[1m%s\033[0m", err, string(line))
				}
				require.NotNil(b, err)
			}
			if !bytes.Equal(names[j], ldeParser.Name) {
				require.Equal(b, string(names[j]), string(ldeParser.Name))
			}
			if avatarsLastcheck[j] != ldeParser.GetAvatarLastCheckValue() {
				require.Equal(b, avatarsLastcheck[j], ldeParser.AvatarLastCheck.Value)
			}
		}
	}
}

var regexComplex = regexp.MustCompile(
	`` +
		`\[\S* (?P<time>.*?)\]` +
		`[^[]*\[(?P<chatID>\d+)\.` +
		`.*?reqid '(?P<reqid>.*?)'` +
		`.*?from.*?\((?P<uin>.*?)\)` +
		`.*?FLAGS\[set:(?P<flagsSet>.*?),` +
		` unset:(?P<flagsUnset>.*?)\]` +
		` FIELDS\[changed:(?P<fieldsChanged>.*?)\]` +
		`(:? ank_ver\[(?P<ankVer>.*?)\])?` +
		`(:? list_ver\[(?P<listVer>.*?)\])?` +
		` name\[(?P<name>.*?)\]` +
		`(:? about\[(?P<about>.*?)\])?` +
		`(:? stamp\[(?P<stamp>.*?)\])?` +
		`(:? regions\[(?P<regions>.*?)\])?` +
		`(:? flags\[(?P<flags>.*?)\])?` +
		` created\[(?P<created>\d+)=.*?\]` +
		`(:? creator\[(?P<creator>.*?)\])?` +
		`(:?.*?avatars_lastcheck\[(?P<avatarsLastCheck>\d+)\])?` +
		`(:?.*?cavatar_lastmod\[(?P<cavatarLastMod>\d+)\])?` +
		` origin\[(?P<origin>.*?)\]` +
		` abuse.*drugs\[(?P<drugs>\d+)\]` +
		` abuse.*spam\[(?P<spam>\d+)\]` +
		` abuse.*porno\[(?P<pron>\d+)\]` +
		`(:? abuse.violation\[(?P<violation>\d+)\])?` +
		`(:? abuse.other\[(?P<abuseOther>\d+)\])?`,
)

func BenchmarkRegexComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j, line := range lines {
			data := regexComplex.FindSubmatch(line)
			if len(data) != 36 {
				require.Equal(b, 35, len(data))
			}
			name := data[12]
			if !bytes.Equal(names[j], name) {
				require.Equal(b, string(names[j]), string(name))
			}
			avLastCheck := data[25]
			if !bytes.Equal(avatarsLastcheckLit[j], avLastCheck) {
				require.Equal(b, string(avatarsLastcheckLit[j]), string(avLastCheck))
			}
		}
	}
}
