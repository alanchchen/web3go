// Copyright (c) 2016, Alan Chen
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package web3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"

	"github.com/alanchchen/web3go/common"
	"github.com/alanchchen/web3go/provider"
	"github.com/tonnerre/golang-go.crypto/sha3"
)

var (
	big0    = big.NewInt(0)
	rat0    = big.NewRat(0, 1)
	unitMap = map[string]string{
		"noether":    "0",
		"wei":        "1",
		"kwei":       "1000",
		"Kwei":       "1000",
		"babbage":    "1000",
		"femtoether": "1000",
		"mwei":       "1000000",
		"Mwei":       "1000000",
		"lovelace":   "1000000",
		"picoether":  "1000000",
		"gwei":       "1000000000",
		"Gwei":       "1000000000",
		"shannon":    "1000000000",
		"nanoether":  "1000000000",
		"nano":       "1000000000",
		"szabo":      "1000000000000",
		"microether": "1000000000000",
		"micro":      "1000000000000",
		"finney":     "1000000000000000",
		"milliether": "1000000000000000",
		"milli":      "1000000000000000",
		"ether":      "1000000000000000000",
		"kether":     "1000000000000000000000",
		"grand":      "1000000000000000000000",
		"mether":     "1000000000000000000000000",
		"gether":     "1000000000000000000000000000",
		"tether":     "1000000000000000000000000000000",
	}
)

// Web3 Standard interface
// See https://github.com/ethereum/wiki/wiki/JavaScript-API#web3js-api-reference
type Web3 struct {
	provider       provider.Provider
	requestManager *requestManager
	Eth            Eth
	Net            Net
}

// NewWeb3 creates a new web3 object.
func NewWeb3(provider provider.Provider) *Web3 {
	requestManager := newRequestManager(provider)
	return &Web3{
		provider:       provider,
		requestManager: requestManager,
		Eth:            newEthAPI(requestManager),
		Net:            newNetAPI(requestManager)}
}

// IsConnected checks if a connection to a node exists.
func (web3 *Web3) IsConnected() bool {
	return true
}

// SetProvider sets provider.
func (web3 *Web3) SetProvider(provider provider.Provider) {
	web3.provider = provider
}

// CurrentProvider returns the current provider.
func (web3 *Web3) CurrentProvider() provider.Provider {
	return web3.provider
}

// Reset state of web3. Resets everything except manager. Uninstalls all
// filters. Stops polling. If keepSyncing is true, it will uninstall all
// filters, but will keep the web3.eth.IsSyncing() polls.
func (web3 *Web3) Reset(keepSyncing bool) {

}

// Sha3 returns Keccak-256 (not the standardized SHA3-256) of the given data.
func (web3 *Web3) Sha3(data string, options interface{}) string {
	opt := struct {
		Encoding string `json:"encoding"`
	}{
		"default",
	}

checkEncoding:
	switch options.(type) {
	case string:
		if err := json.Unmarshal([]byte(options.(string)), &opt); err != nil {
			return common.BytesToHex(web3.sha3Hash([]byte(data)))
		}
		break checkEncoding
	default:
		var err error
		var optBytes []byte
		if optBytes, err = json.Marshal(options); err != nil {
			return common.BytesToHex(web3.sha3Hash([]byte(data)))
		}

		if err = json.Unmarshal(optBytes, &opt); err != nil {
			return common.BytesToHex(web3.sha3Hash([]byte(data)))
		}
		break checkEncoding
	}

	if opt.Encoding == "hex" {
		return common.BytesToHex(web3.sha3Hash(common.HexToBytes(data)))
	}
	return common.BytesToHex(web3.sha3Hash([]byte(data)))
}

// ToHex converts any value into HEX.
func (web3 *Web3) ToHex(value interface{}) string {
	switch value.(type) {
	case bool:
		v := value.(bool)
		if v {
			return "0x1"
		}
		return "0x0"
	case string:
		jsonBytes, err := json.Marshal(value)
		if err != nil {
			return web3.FromDecimal(value)
		}
		unquoted, err := strconv.Unquote(string(jsonBytes))
		if err != nil {
			return common.BytesToHex(jsonBytes)
		}
		return common.BytesToHex([]byte(unquoted))
	case *big.Int:
		return web3.FromDecimal(value)
	default:
		jsonBytes, err := json.Marshal(value)
		if err != nil {
			return web3.FromDecimal(value)
		}
		return common.BytesToHex(jsonBytes)
	}
}

// ToASCII converts a HEX string into a ASCII string.
func (web3 *Web3) ToASCII(hexString string) string {
	return string(bytes.Trim(common.HexToBytes(hexString), "\x00"))
}

// FromASCII converts any ASCII string to a HEX string.
func (web3 *Web3) FromASCII(textString string, padding int) string {
	hex := ""
	for _, runeValue := range textString {
		hex += fmt.Sprintf("%x", runeValue)
	}

	l := len(hex)
	for i := 0; i < padding*2-l; i++ {
		hex += "0"
	}
	return "0x" + hex
}

// ToDecimal converts value to it"s decimal representation in string.
func (web3 *Web3) ToDecimal(value interface{}) string {
	n := web3.ToBigNumber(value)
	if n.IsInt() {
		return n.Num().String()
	}
	return n.String()
}

// FromDecimal converts value to it"s hex representation.
func (web3 *Web3) FromDecimal(value interface{}) string {
	number := web3.ToBigNumber(value)
	if number.IsInt() {
		result := number.Num().Text(16)

		if number.Cmp(rat0) < 0 {
			return "-0x" + result[1:]
		}
		return "0x" + result
	}

	v, _ := number.Float64()
	return fmt.Sprintf("%x", math.Float64bits(v))
}

// FromWei takes a number of wei and converts it to any other ether unit.
//
// Possible units are:
//   SI Short   SI Full        Effigy       Other
// - kwei       femtoether     babbage
// - mwei       picoether      lovelace
// - gwei       nanoether      shannon      nano
// - --         microether     szabo        micro
// - --         microether     szabo        micro
// - --         milliether     finney       milli
// - ether      --             --
// - kether                    --           grand
// - mether
// - gether
// - tether
func (web3 *Web3) FromWei(number string, unit string) string {
	num := web3.ToBigNumber(number)
	returnValue := num.Quo(num, web3.getValueOfUnit(unit))
	return returnValue.Num().String()
}

// ToWei takes a number of a unit and converts it to wei.
//
// Possible units are:
//   SI Short   SI Full        Effigy       Other
// - kwei       femtoether     babbage
// - mwei       picoether      lovelace
// - gwei       nanoether      shannon      nano
// - --         microether     szabo        micro
// - --         microether     szabo        micro
// - --         milliether     finney       milli
// - ether      --             --
// - kether                    --           grand
// - mether
// - gether
// - tether
func (web3 *Web3) ToWei(number interface{}, unit string) string {
	num := web3.ToBigNumber(number)
	returnValue := num.Mul(num, web3.getValueOfUnit(unit))
	return returnValue.Num().String()
}

// ToBigNumber takes an input and transforms it into an *big.Rat.
func (web3 *Web3) ToBigNumber(value interface{}) (result *big.Rat) {
	switch value.(type) {
	case *big.Rat:
		v := value.(*big.Rat)
		return v
	case *big.Int:
		v := value.(*big.Int)
		result = new(big.Rat)
		result.SetInt(v)
		return result
	case string:
		v := value.(string)
		i := new(big.Int)
		result = new(big.Rat)

		if strings.Index(v, "0x") == 0 || strings.Index(v, "-0x") == 0 {
			i.SetString(strings.Replace(v, "0x", "", -1), 16)
		} else {
			i.SetString(v, 10)
		}
		result.SetInt(i)
		return result
	}
	return result
}

// IsAddress checks if the given string is an address.
func (web3 *Web3) IsAddress(address string) bool {
	smallCapsMatcher := regexp.MustCompile("^(0x)?[0-9a-f]{40}$")
	smallCapsMatched := smallCapsMatcher.MatchString(address)
	allCapsMatcher := regexp.MustCompile("^(0x)?[0-9A-F]{40}$")
	allCapsMatched := allCapsMatcher.MatchString(address)
	if smallCapsMatched || allCapsMatched {
		return true
	}
	return web3.isChecksumAddress(address)
}

func (web3 *Web3) isChecksumAddress(address string) bool {
	addr := strings.Replace(address, "0x", "", -1)
	addressHash := web3.Sha3(strings.ToLower(addr), "")

	for i := 0; i < 40; i++ {
		d, err := strconv.ParseInt(string(addressHash[i]), 16, 32)
		if err != nil {
			return false
		}

		if d > 7 && strings.ToUpper(string(address[i])) == string(address[i]) ||
			d <= 7 && strings.ToLower(string(address[i])) == string(address[i]) {
			return false
		}
	}
	return true
}

func (web3 *Web3) sha3Hash(data ...[]byte) []byte {
	d := sha3.NewKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

func (web3 *Web3) getValueOfUnit(unit string) *big.Rat {
	u := strings.TrimSpace(unit)
	if u != "" {
		u = strings.ToLower(u)
	} else {
		u = "ether"
	}

	if unitValue, ok := unitMap[u]; ok {
		value := new(big.Int)
		value.SetString(unitValue, 10)
		returnValue := new(big.Rat)
		returnValue.SetInt(value)
		return returnValue
	}

	keys := make([]string, 0, len(unitMap))
	for k := range unitMap {
		keys = append(keys, k)
	}
	panic(fmt.Sprintf("This unit doesn\"t exists, please use the one of the following units, %v", keys))
}
