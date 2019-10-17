package common

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// os.File Read 사용
// os.Open으로 파일을 열어서 가져온 os.File오브젝트의 Read함수를 사용
// Read함수 인자로 지정된 []byte에 내용이 담긴다.
func baseOsFileRead(filename string) {
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	buf := make([]byte, 64)
	for {
		n, err := fp.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buf))
	}
}

// ioutil.ReadFile 사용
// ioutil.ReadFile함수를 사용해서 파일을 읽는 처리를 하는데 한번에 전체의 데이터를
// 가져와서 이후 읽기 처리를 수행한다. 대용량의 파일사이즈인 경우 주의를 해서 사용해야 한다.
// 대신 코드량은 상당히 간단하다
func baseIoUtilReadFile(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

// bufio.NewReaderSize 사용
// bufio.NewReaderSize를 사용해서 1행씩 읽는다. 1행 사이즈 크기는 2번째 인자로
// 지정한 사이즈를 넘을 경우 isPrefix함수의 리턴값이 false 된다.
func baseBufioNewReaderSize(filename string) {
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 64)
	var tmp []byte
	for {
		line, isPrefix, err := reader.ReadLine() // size を超えるとisPrefixがfalse
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		tmp = append(tmp, line...)
		if !isPrefix {
			fmt.Println(string(tmp))
			tmp = nil
		}
	}
}

// bufio.NewScanner 사용
// bufio.NewScanner사용해서 파일을 1행씩 읽는다. bufio.NewReaderSize를 사용한것과
// 비교해 볼때 보다 간단하게 코드를 작성할 수 있다.
func baseBufioScanner(filename string) {
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// bufio.Reader 사용
func baseReadString(args []string) {
	var fp *os.File
	var err error

	if len(args) < 2 {
		fp = os.Stdin
	} else {
		fmt.Printf(">> read file: %s\n", args[1])
		fp, err = os.Open(args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	reader := bufio.NewReaderSize(fp, 4096)
	for line := ""; err == nil; line, err = reader.ReadString('\n') {
		fmt.Print(line)
	}

	if err != io.EOF {
		panic(err)
	}
}

func baseReadLine(args []string) {
	var fp *os.File
	var err error

	if len(args) < 2 {
		fp = os.Stdin
	} else {
		fmt.Printf(">> read file: %s\n", args[1])
		fp, err = os.Open(args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	reader := bufio.NewReaderSize(fp, 4096)
	for {
		line, _, err := reader.ReadLine()
		fmt.Println(string(line))
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
}

func baseScanner(args []string) {
	var fp *os.File
	var err error

	if len(args) < 2 {
		fp = os.Stdin
	} else {
		fp, err = os.Open(args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// GoLang은 csv파일을 다루는 Reader가 준비되어 있다. 이것을 사용하면 간단하게 csv을 읽을 수 있다.
// 구분자 지정은 Reader.Comma를 이용하여 설정한다. Reader.Comma의 기본값은 ',' 콤마이다.
func baseReadCsvFile(args []string) {
	var fp *os.File
	if len(os.Args) < 2 {
		fp = os.Stdin
	} else {
		var err error
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	reader := csv.NewReader(fp)
	reader.Comma = '\t'
	reader.LazyQuotes = true // ""(더블쿼테이션)을 엄격하게 체크하지 않는다
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println(record)
	}
}
