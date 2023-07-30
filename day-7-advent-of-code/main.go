package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File interface {
  GetSize() int
  GetName() string
}

type CommonFile struct {
  Size int
  Name string
}

func (commonFile *CommonFile) GetSize() int {
  return commonFile.Size
}

func (commonFile *CommonFile) GetName() string {
  return commonFile.Name
}

type DirectoryFile struct {
  Name string
  Files map[string]File
  Parent *DirectoryFile
}

func (directoryFile *DirectoryFile) GetSize() int {
  sizeSum := 0
  for _, file := range directoryFile.Files {
    sizeSum += file.GetSize()
  }
  return sizeSum
}

func (directoryFile *DirectoryFile) GetName() string {
  return directoryFile.Name
}

func (directoryFile *DirectoryFile) GetSonDirRefOrPanic(son string) *DirectoryFile {
  newDir, ok := directoryFile.Files[son].(*DirectoryFile)
  if !ok {
    panic("CD to file")
  }
  return newDir
}

type CLIInput int
const (
  LS CLIInput = iota
  CD
  FILE
  DIR
  NONE
)

func GetCliInputTypeAndTokens(str string) (CLIInput, []string) {
  tokens := strings.SplitAfter(str, " ")
  for token := range tokens {
    tokens[token] = strings.TrimSpace(tokens[token])
  }
  if len(tokens) < 2 {
    return NONE, tokens
  }
  if tokens[0] == "$" {
    if tokens[1] == "ls" {
      return LS, tokens
    }
    return CD, tokens
  }
  if tokens[0] == "dir" {
    return DIR, tokens
  }
  return FILE, tokens
}

func NewDir(name string) *DirectoryFile {
  return &DirectoryFile{Name: name, Files: make(map[string]File, 0), Parent: nil}
}

func NewDirFromHere(name string, parent *DirectoryFile) *DirectoryFile {
  return &DirectoryFile{Name: name, Files: make(map[string]File, 0), Parent: parent}
}

func SumDirSizesBiggerThan(directory DirectoryFile, parameter int) int {
  sumOfSubDirsBigger := 0
  for _, file := range directory.Files {
    dir, isDirectory := file.(*DirectoryFile)
    if isDirectory {
      sumOfSubDirsBigger += SumDirSizesBiggerThan(*dir, parameter)
      dirSize := dir.GetSize()
      if dirSize < parameter {
        fmt.Print(file.GetName() + " is on with ")
        fmt.Printf("%d\n", dirSize)
        sumOfSubDirsBigger += dirSize
      }
    }
  }
  return sumOfSubDirsBigger
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetSizeOfSmallestToDelete(basedir DirectoryFile, needToFree int) int {
  smallestDirSizeToDelete := int(^uint(0)>>1)
  for _, file := range basedir.Files {
    dir, isDirectory := file.(*DirectoryFile)
    if isDirectory {
      dirSize := dir.GetSize()
      if dirSize < smallestDirSizeToDelete && dirSize > needToFree {
        smallestDirSizeToDelete = dirSize
      }
      smallestDirSizeToDelete = min(smallestDirSizeToDelete, GetSizeOfSmallestToDelete(*dir, needToFree))
    }
  }
  return smallestDirSizeToDelete
}

func AOCSmallestDirToDelete(filePath string) int {
  file, err := os.Open(filePath) 
  if err != nil {
    panic("can't open " + filePath)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  
  baseDir := NewDir("/")
  currentDir := baseDir
  for scanner.Scan() {
    commandType, tokens := GetCliInputTypeAndTokens(scanner.Text())

    switch(commandType) {
    case CD:
      dir_to_change := tokens[2] 
      if dir_to_change == ".." {
        currentDir = currentDir.Parent
      } else if dir_to_change == "/" {
        currentDir = baseDir
      } else if currentDir.Files[dir_to_change] == nil {
        currentDir.Files[dir_to_change] = NewDirFromHere(dir_to_change, currentDir)
        currentDir = currentDir.GetSonDirRefOrPanic(dir_to_change)
      } else if currentDir.Files[dir_to_change] != nil {
        currentDir = currentDir.GetSonDirRefOrPanic(dir_to_change)
      }
    case LS:
      //currentDir.Files = make(map[string]File, 0)
    case FILE:
      fileSize := tokens[0]
      fileName := tokens[1]
      convertedFileSize, err := strconv.Atoi(fileSize)
      if err != nil {
        panic("Can't parse convertedFileSize")
      }
      currentDir.Files[fileName] = &CommonFile{Name: fileName, Size: convertedFileSize}
    case DIR:
      fileName := tokens[1]
      currentDir.Files[fileName] = NewDirFromHere(fileName, currentDir)
    }
  }
  return GetSizeOfSmallestToDelete(*baseDir, 30000000 - 70000000 + baseDir.GetSize())
}

func main() {
  fmt.Println("EXAMPLE")
  fmt.Println(AOCSmallestDirToDelete("test.txt"))
  fmt.Println("-------")
  fmt.Println("INPUT")
  fmt.Println(AOCSmallestDirToDelete("input.txt"))
}
