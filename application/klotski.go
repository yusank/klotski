/*
* MIT License
*
* Copyright (c) 2017 Yusan Kurban.
*
* Permission is hereby granted, free of charge, to any person obtaining a copy of
* this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:
*
* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.
*
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
 */

/*
* Revision History
*     Initial: 2017/08/30          Yusan Kurban
 */

package main

type character struct {
	Name     string   //
	Position position //
	Cover    [][]int  //
	Control  string   // short name
}

type position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

var (
	chessBoard = [][]character{}
	cover22    = [][]int{}
	cover12    = [][]int{}
	cover21    = [][]int{}
	cover11    = [][]int{}
	empty      character
	caoCao     character
	zhangFei   character
	zhaoYun    character
	huangZhong character
	maChao     character
	guanYu     character
	zu         character
	bing       character
	shi        character
	zuo        character
	all        []*character // collection of character
)

func initCharacter() {
	cover22 = make([][]int, 2)
	for i := range cover22 {
		column := make([]int, 2)

		cover22[i] = column
	}

	cover12 = make([][]int, 2)
	for i := range cover12 {
		column := make([]int, 1)

		cover12[i] = column
	}

	cover21 = make([][]int, 1)
	column := make([]int, 2)
	cover21[0] = column

	cover11 = make([][]int, 1)
	column = make([]int, 1)
	cover11[0] = column

	caoCao = character{
		Name: "曹操",
		Position: position{
			X: 1,
			Y: 0,
		},
		Cover:   cover22,
		Control: "cc",
	}
	all = append(all, &caoCao)

	zhangFei = character{
		Name: "张飞",
		Position: position{
			X: 0,
			Y: 0,
		},
		Cover:   cover12,
		Control: "zf",
	}
	all = append(all, &zhangFei)

	zhaoYun = character{
		Name: "赵云",
		Position: position{
			X: 3,
			Y: 0,
		},
		Cover:   cover12,
		Control: "zy",
	}
	all = append(all, &zhaoYun)

	huangZhong = character{
		Name: "黄忠",
		Position: position{
			X: 3,
			Y: 2,
		},
		Cover:   cover12,
		Control: "hz",
	}
	all = append(all, &huangZhong)

	maChao = character{
		Name: "马超",
		Position: position{
			X: 0,
			Y: 2,
		},
		Cover:   cover12,
		Control: "mc",
	}
	all = append(all, &maChao)

	guanYu = character{
		Name: "关羽",
		Position: position{
			X: 1,
			Y: 2,
		},
		Cover:   cover21,
		Control: "gy",
	}
	all = append(all, &guanYu)

	zu = character{
		Name: "小卒",
		Position: position{
			X: 0,
			Y: 4,
		},
		Cover:   cover11,
		Control: "zu",
	}
	all = append(all, &zu)

	bing = character{
		Name: "小兵",
		Position: position{
			X: 1,
			Y: 3,
		},
		Cover:   cover11,
		Control: "bi",
	}
	all = append(all, &bing)

	shi = character{
		Name: "小士",
		Position: position{
			X: 2,
			Y: 3,
		},
		Cover:   cover11,
		Control: "sh",
	}
	all = append(all, &shi)

	zuo = character{
		Name: "小佐",
		Position: position{
			X: 3,
			Y: 4,
		},
		Cover:   cover11,
		Control: "zo",
	}
	all = append(all, &zuo)
}
