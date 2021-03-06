// Copyright 2017 ZhongAn Information Technology Services Co.,Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package expr

import (
	"strings"
	"testing"

	gcmn "github.com/dappledger/AnnChain/gemmill/modules/go-common"
)

func TestParse(t *testing.T) {
	testParse(t, `"foobar"`, `"foobar"`)
	testParse(t, "0x1234", "0x1234")
	testParse(t, "0xbeef", "0xBEEF")
	testParse(t, "xbeef", "xBEEF")
	testParse(t, "12345", "{i 12345}")
	testParse(t, "u64:12345", "{u64 12345}")
	testParse(t, "i64:12345", "{i64 12345}")
	testParse(t, "i64:-12345", "{i64 -12345}")
	testParse(t, "[1 u64:2]", "[{i 1},{u64 2}]")
	testParse(t, "[(1 2) (3 4)]", "[({i 1} {i 2}),({i 3} {i 4})]")
	testParse(t, "0x1234 1 u64:2 [3 4]", "(0x1234 {i 1} {u64 2} [{i 3},{i 4}])")
	testParse(t, "[(1 <sig:user1>)(2 <sig:user2>)][3 4]",
		"([({i 1} <sig:user1>),({i 2} <sig:user2>)] [{i 3},{i 4}])")
}

func testParse(t *testing.T, input string, expected string) {
	got, err := ParseReader(input, strings.NewReader(input))
	if err != nil {
		t.Error(err.Error())
		return
	}
	gotStr := gcmn.Fmt("%v", got)
	if gotStr != expected {
		t.Error(gcmn.Fmt("Expected %v, got %v", expected, gotStr))
	}
}

func TestBytes(t *testing.T) {
	testBytes(t, `"foobar"`, `0106666F6F626172`)
	testBytes(t, "0x1234", "01021234")
	testBytes(t, "0xbeef", "0102BEEF")
	testBytes(t, "xbeef", "BEEF")
	testBytes(t, "12345", "023039")
	testBytes(t, "u64:12345", "0000000000003039")
	testBytes(t, "i64:12345", "0000000000003039")
	testBytes(t, "i64:-12345", "FFFFFFFFFFFFCFC7")
	testBytes(t, "[1 u64:2]", "010201010000000000000002")
	testBytes(t, "[(1 2) (3 4)]", "01020101010201030104")
	testBytes(t, "0x1234 1 u64:2 [3 4]", "0102123401010000000000000002010201030104")
	testBytes(t, "[(1 <sig:user1>)(2 <sig:user2>)][3 4]",
		"0102010100010200010201030104")
}

func testBytes(t *testing.T, input string, expected string) {
	got, err := ParseReader(input, strings.NewReader(input))
	if err != nil {
		t.Error(err.Error())
		return
	}
	gotBytes, err := got.(Byteful).Bytes()
	if err != nil {
		t.Error(err.Error())
		return
	}
	gotHex := gcmn.Fmt("%X", gotBytes)
	if gotHex != expected {
		t.Error(gcmn.Fmt("Expected %v, got %v", expected, gotHex))
	}
}
