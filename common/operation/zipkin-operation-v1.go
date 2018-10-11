package operation

const (
	ZipKinOperationV1Topic = "zipkin-operation-v1"
)

type ZipKinOperationV1 struct {
	Item []byte `json:"item"`
}
