package solutions

type Task struct {
	Payload []int `json:"payload"`
	Result  []int `json:"result"`
}

const (
	userName    string = "il_linberg"
	Port        string = ":1701"
	SolutionURL string = "https://kuvaev-ituniversity.vps.elewise.com"
	Rotation    string = "Циклическая ротация"
	FindUnique  string = "Чудные вхождения в массив"
	Cons        string = "Проверка последовательности"
	Missing     string = "Поиск отсутствующего элемента"
)
