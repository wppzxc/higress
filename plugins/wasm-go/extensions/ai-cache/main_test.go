package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"
	"time"
	"unicode"

	"github.com/tidwall/gjson"
)

func TestTrimChar(t *testing.T) {
	s := "Hello, World!   "
	fmt.Println("Original:", s)
	s = trimRightPunctAndSpace(s)
	fmt.Println("Trimmed:", s)
}

func TestSliceNil(t *testing.T) {
	jsonStr := `{"code":0,"request_id":"a684546a-1e23-936f-8d20-ff5b4e301f71","message":"Success"}`
	vector := new(VectorSearchResponse)
	json.Unmarshal([]byte(jsonStr), vector)
	if len(vector.Output) < 1 {
		fmt.Println("empty")
	} else {
		fmt.Println("not empty")
	}
}

func TestJsonArray(t *testing.T) {
	responseBody := []byte(`{"code":0,"request_id":"a684546a-1e23-936f-8d20-ff5b4e`)
	embeddingResultJson := gjson.ParseBytes(responseBody)
	embeddingArray := embeddingResultJson.Get("output.embeddings.0.embedding").Array()
	fmt.Printf("%#v", embeddingArray)
}

func TestCosine(t *testing.T) {
	vectorA := []float64{0.02247256240777982, -0.020711698613635976}
	vectorB := []float64{0.0380711, 0.0186769}

	distCosine := distCosine(vectorA, vectorB)
	fmt.Printf("Cosine Similarity: %f\n", distCosine)
}

func TestHanCount(t *testing.T) {
	str := "java呢？"
	var count = 0
	for _, c := range str {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Printf("count: '%d'", count)
}

func mySyncHttpRequest(path string, reqHeaders [][2]string, reqBody []byte) (int, []byte) {
	// 变量初始化
	var (
		finish = false
		result = ""
	)
	// 请求http接口
	go func(path string, reqHeaders [][2]string, reqBody []byte) {
		defer func() {
			finish = true
		}()
		time.Sleep(3 * time.Second)
		result = "ok"
	}(path, reqHeaders, reqBody)
	// 等待请求完成
	for {
		if !finish {
			continue
		}
		fmt.Println(result)
		break
	}
	return 0, []byte(result)
}

func TestSyncHttpRequest(t *testing.T) {
	mySyncHttpRequest("http://www.baidu.com", [][2]string{}, []byte{})
}

type Item struct {
	Value int
}

// 定义比较函数
func findSmallestThree(items []Item, compareValue int) []Item {
	// 筛选出小于比较值的元素
	var filteredItems []Item
	for _, item := range items {
		if item.Value < compareValue {
			filteredItems = append(filteredItems, item)
		}
	}

	// 按照Value排序
	sort.Slice(filteredItems, func(i, j int) bool {
		return filteredItems[i].Value < filteredItems[j].Value
	})

	// 获取最小的三个结果
	if len(filteredItems) > 3 {
		return filteredItems[:3]
	}
	return filteredItems
}

func TestSortSlice(t *testing.T) {
	items := []Item{
		{Value: 10},
		{Value: 5},
		{Value: 8},
		{Value: 3},
		{Value: 15},
		{Value: 1},
	}
	compareValue := 9

	// 调用比较函数
	result := findSmallestThree(items, compareValue)

	// 输出结果
	fmt.Println("Smallest three items less than", compareValue, ":")
	for _, item := range result {
		fmt.Println(item)
	}
}

// func TestLen(t *testing.T) {
// 	str := "地址是?"
// 	// 输出：字符串中的汉字数量: 3
// 	fmt.Println("字符串中的汉字数量:", countChineseChars(str))
// }
