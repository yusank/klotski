/*
 * MIT License
 *
 * Copyright (c) 2017 Yusan Kurban.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
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
 * Revision History:
 *     Initial: 2017/08/31        Yusan Kurban
 */

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type gameSchedule struct {
	GameLoad []save `json:"gameload"`
}

type save struct {
	Name     string   `json:"name"`
	Position position `json:"position"`
}

const (
	timeFormat   = "2006-01-02T15:04:05"
	validPostfix = "kls"
)

func marshalAndSave() (string, error) {
	var (
		char  save
		chars []save
		sch   gameSchedule
	)
	name := fmt.Sprintf("klotski-%s.kls", time.Now().Format(timeFormat))
	file, err := os.Create(name)
	if err != nil {
		return "", err
	}
	defer file.Close()

	for _, c := range all {
		char.Name = c.Name
		char.Position = c.Position

		chars = append(chars, char)
	}

	sch.GameLoad = chars
	by, err := json.Marshal(sch)
	if err != nil {
		return "", err
	}

	_, err = file.Write(by)
	if err != nil {
		return "", err
	}

	return name, nil
}

func readFile(fileName string) error {
	var sch gameSchedule

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	body := make([]byte, 435)
	_, err = file.Read(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &sch)
	if err != nil {
		return err
	}

	for _, c := range all {
		for _, load := range sch.GameLoad {
			if c.Name == load.Name {
				c.Position.X = load.Position.X
				c.Position.Y = load.Position.Y
			}
		}
	}

	return nil
}
