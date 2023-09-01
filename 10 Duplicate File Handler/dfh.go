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
	fileFormat := inputClosure()()
	fileFormat = strings.TrimPrefix(fileFormat, ".")
	fmt.Println()
	sortOption := getSortOption()
	list := walkDir(fileFormat)
	fmt.Println()
	printMap(list, sortOption)
	filesByOrder := handleDuplicates(list)
	deleteFiles(filesByOrder)
}

func deleteFiles(filesByOrder map[int]string) {
	input := inputClosure()
	for {
		fmt.Println("Delete files? (yes|no)")
		removeFiles := strings.ToLower(input())
		if removeFiles == "yes" {
			fmt.Println()
			break
		}
		if removeFiles == "no" {
			return
		}
		fmt.Println("Wrong option")
	}
	var nums []int
	runOuterLoop := true
	for runOuterLoop {
		nums = make([]int, 0)
		fmt.Println("Enter file numbers to delete:")
		numsInput := input()
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
		totalFreedSpace += getFileSize(filePath)
		err := os.Remove(filePath)
		if err != nil {
			fmt.Println("Error deleting this file:", filePath)
			return
		}
	}
	fmt.Printf("Total freed up space: %d bytes\n", totalFreedSpace)
}

func getFileSize(filePath string) int {
	fi, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Could not obtain stat, handle error")
	}
	return int(fi.Size())
}

func handleDuplicates(list map[int][]string) map[int]string {
	input := inputClosure()
	for {
		fmt.Println("Check for duplicates? (yes|no)")
		checkDuplicates := strings.ToLower(input())
		if checkDuplicates == "yes" {
			break
		}
		if checkDuplicates == "no" {
			return nil
		}
		fmt.Println("Wrong option")
		fmt.Println()
	}

	filesSizes := make(map[int]map[string][]string)
	filesByOrder := make(map[int]string)
	cnt := 1
	for fileSize, filePaths := range list {
		filesSizes[fileSize] = make(map[string][]string)
		for _, filePath := range filePaths {
			hash := fmt.Sprintf("%x", fileHash(filePath))
			filesSizes[fileSize][hash] = append(filesSizes[fileSize][hash], filePath)
			filesByOrder[cnt] = filePath
			cnt++
		}
	}

	cnt = 1
	for size, filesHash := range filesSizes {
		fmt.Println(size, "bytes")
		for hash, filePaths := range filesHash {
			fmt.Println(hash)
			for _, filePath := range filePaths {
				fmt.Printf("%d. %s\n", cnt, filePath)
				cnt++
			}
			fmt.Println()
		}
	}

	return filesByOrder
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

func getSortOption() int {
	input := inputClosure()
	fmt.Println("Size sorting options:\n1. Descending\n2. Ascending")
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

func inputClosure() func() string {
	reader := bufio.NewReader(os.Stdin)
	return func() string {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		return line
	}
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

func printMap(list map[int][]string, sortingOption int) {
	keys := make([]int, 0, len(list))
	for k := range list {
		keys = append(keys, k)
	}

	if sortingOption == 2 {
		sort.Ints(keys)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	}

	for _, size := range keys {
		fmt.Println(size, "bytes")
		for i := range list[size] {
			fmt.Println(list[size][i])
		}
		fmt.Println("")
	}
}
