package main

import (
	"fmt"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

type httpServer struct {
	graphicParams []graphicParameter
}

type graphicParameter struct {
	tao   float64
	n     int
	layer []float64
}

func initialConditions(x float64) float64 {
	if x < 1 {
		return 3.
	} else if x == 1 {
		return 2.
	} else {
		return 1.
	}
}

func lineItemsByLayer(layer []float64) []opts.LineData {
	items := make([]opts.LineData, 0, len(layer))
	for _, l := range layer {
		items = append(items, opts.LineData{Value: l})
	}
	return items
}

func lineItemExactSolution(args []float64) []opts.LineData {
	items := make([]opts.LineData, 0, len(args))
	for _, l := range args {
		items = append(items, opts.LineData{Value: initialConditions(l - 1)})
	}
	return items
}

func (s *httpServer) ServerFunc(w http.ResponseWriter, _ *http.Request) {
	for _, p := range s.graphicParams {
		baseNumbers := make([]float64, p.n+1)
		for i := 0; i <= p.n; i++ {
			baseNumbers[i] = -10. + float64(i)*20./float64(p.n)
		}
		line := charts.NewLine()
		line.SetGlobalOptions(
			charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
			charts.WithTitleOpts(opts.Title{
				Title: fmt.Sprintf(
					"Решение уравнения для T = 1 на отрезке [-10; 10] с tao = %v, n = %v",
					p.tao,
					p.n,
				),
				Subtitle: "Синий - численное решение, голубой - точное",
			}),
		)
		line.SetXAxis(baseNumbers).
			AddSeries("", lineItemsByLayer(p.layer)).
			AddSeries("", lineItemExactSolution(baseNumbers)).
			SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
		line.Render(w)
	}
}

func main() {
	graphicParams := make([]graphicParameter, 0, 8)
	l := -10.
	r := +10.
	possibleN := []int{100, 1_000}
	for _, n := range possibleN {
		h := (r - l) / float64(n)
		a := 1.
		possibleTao := []float64{0.25 * h, 0.5 * h, h, 1.25 * h}
		for _, tao := range possibleTao {
			fmt.Printf("l: %v\nr: %v\nn: %v\nh: %v\na: %v\ntao: %v\n", l, r, n, h, a, tao)
			targetT := 1.
			layer := make([]float64, n+1)
			for i := 0; i < n+1; i++ {
				layer[i] = initialConditions(l + float64(i)*h)
			}

			for layerT := tao; layerT <= targetT; layerT += tao {
				u__j_sub_1__n := layer[0]
				k := a * tao / h

				for i := 1; i <= n; i++ {
					u__j__n := layer[i]
					resultU := u__j__n - k*(u__j__n-u__j_sub_1__n)
					layer[i] = resultU
					u__j_sub_1__n = u__j__n
				}
			}

			// file, err := os.OpenFile(
			// 	fmt.Sprintf("./n=%v;tao=%v.txt", n, tao),
			// 	os.O_CREATE|os.O_RDWR,
			// 	0644,
			// )
			// if err != nil {
			// 	fmt.Println("ERROR:" + err.Error())
			// 	return
			// }
			// for _, result := range layer {
			// 	fmt.Fprintf(file, "%v ", result)
			// }
			graphicParams = append(graphicParams, graphicParameter{
				tao:   tao,
				n:     n,
				layer: layer,
			})
		}
	}
	server := httpServer{
		graphicParams: graphicParams,
	}
	http.HandleFunc("/", server.ServerFunc)
	fmt.Println("Сервер готов. Для просмотра графиков пройдите по ссылке http://localhost:8069/")
	http.ListenAndServe(":8069", nil)
}
