package microClustering

import (
	"fmt"
	"testing"
)

func TestClassifier(t *testing.T) {

	//data := [][]float64{{2.0, 2.0}, {1.0, 3.0}, {2.0, 8.0}, {2.0, 9.0}, {3, 8}, {4, 6}, {4, 7}, {4, 9}, {5, 7}, {5, 8}, {5, 9}, {6, 4}, {7, 5}, {9, 4}}

	data := [][]float64{{3.0, 2.0, 1.0}, {3.1, 2.1, 1.0}, {3.0, 2.1, 1.0}, {2.9, 1.9, 1.0}, {3.0, 1.9, 1.0},
		{2.0, 4.0, 1.0}, {2.1, 4.1, 1.0}, {1.9, 4.1, 1.0}, {1.9, 3.9, 1.0}, {2.1, 3.9, 1.0}, {2.1, 4.0, 1.0}, {1.9, 4.0, 1.0}, {2.0, 4.0, 1.0},

		{7.0, 2.0, 2.0}, {6.9, 1.9, 2.0}, {7.1, 2.1, 2.0},
		{9.0, 2.0, 2.0}, {9.1, 2.1, 2.0}, {8.9, 2.1, 2.0}, {8.9, 1.9, 2.0}, {9.1, 1.9, 2.0}, {9.1, 2.0, 2.0}, {8.9, 2.0, 2.0}, {9.0, 2.0, 2.0},
		{8, 4, 2}, {8.1, 4.1, 2}, {7.9, 3.9, 2}, {7.9, 4.1, 2}, {8.1, 3.9, 2},
		{10, 4, 2}, {10.1, 4.1, 2}, {9.9, 3.9, 2}, {10.1, 3.9, 2}, {9.9, 4.1, 2}, {10.1, 4, 2}, {9.9, 4, 2}, {10, 4.1, 2}, {10, 3.9, 2}, {10, 4, 2},
		{9, 5, 2}, {9.1, 5.1, 2}, {8.9, 5.1, 2}, {8.9, 4.9, 2}, {9.1, 4.9, 2}, {9.1, 5, 2}, {8.9, 5, 2}, {9, 5.1, 2}, {9, 4.9, 2}, {9, 5, 2},

		{5, 9, 3}, {5.1, 9, 3}, {4.9, 9, 3}, {5, 9.1, 3}, {5, 8.9, 3},
		{6, 9, 3}, {6.1, 9.1, 3}, {6.1, 8.9, 3}, {5.9, 8.9, 3}, {5.9, 9.1, 3}, {5.9, 9, 3}, {6.1, 9, 3}, {6, 9, 3},
		{3, 10, 3}, {3, 10.1, 3}, {3.1, 10, 3},
		{5, 11, 3}, {5, 11.1, 3}, {5, 10.9, 3}, {5.1, 11, 3}, {4.9, 11, 3}, {5.1, 11.1, 3}, {4.9, 10.9, 3}, {5.1, 11.1, 3}, {4.9, 10.9, 3}, {5, 11, 3},
	}

	radius := 0.0
	min := 2
	label := 2
	SetDistanceFunction("cosinus")
	c := NewClassifier(label, radius, min, 1, 3.0)
	c.Verbose = 0
	c.Fit(data)

	test := [][]float64{{5, 3}, {6, 4}, {5, 6}}
	y := c.KNN(test, 3)
	fmt.Println("y=", y)
	//	c.Stats()

	js, err := c.ToJson()
	fmt.Println("err:", err)
	fmt.Println("js:", string(js))

	test = [][]float64{{5, 3}, {5, 6}, {6, 4}, {5, 6}}
	c2, err := NewClassifierFromJson(js)
	fmt.Println("err:", err)
	y2 := c2.KNN(test, 3)
	fmt.Println("y2=", y2)

}
