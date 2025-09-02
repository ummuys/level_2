package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

func main() {
	var err error // если мы сделали именно err := test(), то мы бы не прошли проверку
	// так как моржовый оператор сделал бы err как *customError, а не error (interface)
	err = test()
	// Мы возвращаем *customError, который nil
	// НО! так как customError реализует интерфейс error
	// то в интерфейс error вносятся такие значения: data -> nil, а вот type -> *customError
	// Поэтому сам err, который является интерфейсом, не пустой и мы проходим условие
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
