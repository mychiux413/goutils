package c

func GetSizeMB(data []byte) float64 {
	return float64(len(data)) / 1024 / 1024
}
