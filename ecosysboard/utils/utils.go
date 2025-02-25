/******************************************************************************
 * Copyright © 2013-2019 The Komodo Platform Developers.                      *
 *                                                                            *
 * See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 * the top-level directory of this distribution for the individual copyright  *
 * holder information and the developer policies on copyright and licensing.  *
 *                                                                            *
 * Unless otherwise agreed in a custom licensing agreement, no part of the    *
 * Komodo Platform software, including this file may be copied, modified,     *
 * propagated or distributed except according to the terms contained in the   *
 * LICENSE file                                                               *
 *                                                                            *
 * Removal or modification of this copyright notice is prohibited.            *
 *                                                                            *
 ******************************************************************************/

package utils

import (
	"io/ioutil"
	"os"
	"regexp"
)

func IsLookLikeAKomodoAddress(address string) bool {
	var re = regexp.MustCompile(`(?m)^[R][a-km-zA-HJ-NP-Z1-9]{27,34}$`)
	return re.MatchString(address)
}

func IsLookLikeABlock(input string) bool {
	var re = regexp.MustCompile(`(?m)^[0-9]*$`)
	return re.MatchString(input)
}

func IsLookLikeABlockHashOrTransactionId(input string) bool {
	var re = regexp.MustCompile(`(?m)^[0-9a-f]{64}$`)
	return re.MatchString(input)
}

func IsPathExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func DeleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func IOReadDir(root string, prefix string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		if len(prefix) > 0 {
			files = append(files, prefix+"/"+file.Name())
		} else {
			files = append(files, file.Name())
		}
	}
	return files, nil
}

//000000014c0797b609abef168e8df13c03b92415f3a9b00c9f583013b5824b06
//2909f0b98ca4c3812bb500bc79fa83d3b8c8159c8f4328ac80777968f59400c7
