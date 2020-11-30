/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http:www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// duplicateCmd represents the dublicate command
var isDelete bool

var duplicateCmd = &cobra.Command{
	Use:   "duplicate",
	Short: "Find duplicate and delete easily",
	Long:  `Find duplicator a simple CLI tool which is help you for find the duplicate files and delete them. `,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		getFiles(args[0])
	},
}

func init() {
	rootCmd.AddCommand(duplicateCmd)
	duplicateCmd.PersistentFlags().BoolVarP(&isDelete, "delete", "d", false, "verbose output")
	duplicateCmd.MarkFlagRequired("delete")

}
func getFiles(str string) {

	uniqueList := make(map[string]string)
	willDelete := make(map[string]string)

	files, err := ioutil.ReadDir(str)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			f, err := os.Open(str + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			h := sha256.New()
			if _, err := io.Copy(h, f); err != nil {
				log.Fatal(err)
			}
			//sha256 convert to string
			var key = hex.EncodeToString(h.Sum(nil))
			uniqueList[key] = f.Name()
			willDelete[f.Name()] = key
		}
	}

	/*
		İf a key exists in wilDelete(map) value,
		then just delete element in willDelete list,
	    For example let we just thing we have 4 jpeg files.
	    dream1.jpeg, dream2.jpeg,dream3.jpeg,dream4.jpeg
		I created two map 1- uniqueList 2-willDelete
		uniqueList's key is sha256 and value is path, willDelete is opposite.
		İf uniqueList's key which is sha256 contains willDelete's value which is same, then delete willDelete
		because we are saying, this file is unique, do not delete.
		After the uniqueList length, we have non-unique-files to delete.
	*/

	for a := range uniqueList {
		for b := range willDelete {
			if a == willDelete[b] {
				delete(willDelete, uniqueList[a])
			}
		}
	}

	// delete all none unique files
	for i := range willDelete {
		fmt.Println("Duplicate File :", i)
		if isDelete {
			os.Remove(i)
		}
	}
	fmt.Println("Total Scanned Files :", len(uniqueList)+len(willDelete))
	fmt.Println("Total Unique Files :", len(uniqueList))
	fmt.Println("Total Duplicate Files :", len(willDelete))
	if isDelete {
		fmt.Println("Total Deleted Files :", len(willDelete))
	}
}
