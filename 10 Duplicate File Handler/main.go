package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Directory is not specified")
		return
	}

	fmt.Println("Enter file format:")
	fileFormat := input()
	fileFormat = strings.TrimPrefix(fileFormat, ".")

	sortOption := getSortOption()
	list := walkDir(fileFormat)
	printFiles(list, sortOption)

	if getConfirmation("Check for duplicates?") {
		fmt.Println()
		filesByOrder := handleDuplicates(list, sortOption)
		if getConfirmation("Delete files?") {
			deleteFiles(filesByOrder)
		}
	}
	fmt.Println("bye-bye.")
}

func deleteFiles(filesByOrder map[int]string) {
	var nums []int
	runOuterLoop := true
	for runOuterLoop {
		nums = make([]int, 0)
		fmt.Println("Enter file numbers to delete:")
		numsInput := input()
		if len(numsInput) == 0 {
			fmt.Println("Wrong format")
			continue
		}
		numsString := strings.Fields(numsInput)
		for _, s := range numsString {
			sInt, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Wrong format")
				runOuterLoop = true
				break
			}
			if _, ok := filesByOrder[sInt]; !ok {
				fmt.Println("Wrong format")
				runOuterLoop = true
				break
			}
			nums = append(nums, sInt)
			runOuterLoop = false
		}
	}

	totalFreedSpace := 0
	for _, num := range nums {
		filePath := filesByOrder[num]
		totalFreedSpace += fileSize(filePath)
		err := os.Remove(filePath)
		if err != nil {
			fmt.Println("Error deleting this file:", filePath)
			return
		}
	}
	fmt.Printf("Total freed up space: %d bytes\n", totalFreedSpace)
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func fileHash(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	hashAlg := md5.New()
	if _, err := io.Copy(hashAlg, file); err != nil {
		log.Fatal(err)
	}
	return hashAlg.Sum(nil)
}

func fileSize(filePath string) int {
	fi, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Could not obtain stat, handle error")
	}
	return int(fi.Size())
}

func getConfirmation(prompt string) bool {
	for {
		fmt.Println(prompt + " (yes|no)")
		response := strings.ToLower(input())
		switch response {
		case "yes":
			return true
		case "no":
			return false
		default:
			fmt.Println("Wrong option")
		}
	}
}

func getSortOption() int {
	fmt.Println("\nSize sorting options:\n1. Descending\n2. Ascending")
	for {
		fmt.Println()
		fmt.Println("Enter a sorting option:")
		sortOption, _ := strconv.Atoi(input())
		if sortOption == 2 || sortOption == 1 {
			return sortOption
		}
		fmt.Println("Wrong Option")
	}
}

func handleDuplicates(list map[int][]string, sortOption int) map[int]string {
	sizes := make(map[int]map[string][]string)
	for size, filePaths := range list {
		sizes[size] = make(map[string][]string)
		for _, filePath := range filePaths {
			hash := fmt.Sprintf("%x", fileHash(filePath))
			sizes[size][hash] = append(sizes[size][hash], filePath)
		}
	}

	cnt := 1
	filesByOrder := make(map[int]string)
	sortedKeys := getSortedKeys(sizes, sortOption)
	for _, size := range sortedKeys {
		writeSize := true
		for hash, filePaths := range sizes[size] {
			if len(filePaths) < 2 {
				continue
			}
			if writeSize {
				fmt.Println(size, "bytes")
				writeSize = false
			}
			fmt.Println("Hash:", hash)
			for _, filePath := range filePaths {
				fmt.Printf("%d. %s\n", cnt, filePath)
				filesByOrder[cnt] = filePath
				cnt++
			}
		}
		fmt.Println()
	}

	return filesByOrder
}

func printFiles(list map[int][]string, sortOption int) {
	fmt.Println()
	sortedKeys := getSortedKeys(list, sortOption)
	for _, size := range sortedKeys {
		fmt.Println(size, "bytes")
		for _, filePath := range list[size] {
			fmt.Println(filePath)
		}
		fmt.Println("")
	}
}

func getSortedKeys[T any](arr map[int]T, sortOption int) []int {
	sortedKeys := make([]int, 0)
	for size := range arr {
		sortedKeys = append(sortedKeys, size)
	}
	if sortOption == 1 {
		sort.Sort(sort.Reverse(sort.IntSlice(sortedKeys)))
	} else {
		sort.Ints(sortedKeys)
	}
	return sortedKeys
}

func walkDir(fileFormat string) map[int][]string {
	list := make(map[int][]string)
	rootDir := os.Args[1]
	err := filepath.WalkDir(
		rootDir,
		func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				log.Fatal("err is: ", err)
			}
			if entry.IsDir() {
				return nil
			}
			if len(fileFormat) == 0 {
				x, _ := entry.Info()
				size := int(x.Size())
				list[size] = append(list[size], path)
			} else if strings.TrimLeft(filepath.Ext(entry.Name()), ".") == fileFormat {
				x, _ := entry.Info()
				size := int(x.Size())
				list[size] = append(list[size], path)
			}
			return nil
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	return list
}
