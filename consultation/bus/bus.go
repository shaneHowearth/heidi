package bus

type Carrier interface {
	Read(count int) ([][]byte, error)
}
