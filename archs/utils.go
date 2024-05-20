package archs

func bigEndian16Bytes(encoding []byte) []byte {
    result := make([]byte, len(encoding))
    for i := 0; i < len(encoding); i++ {
        result[i] = encoding[i]
    }

    for i := 0; i < len(encoding); i += 16 {
        end := i + 16
        if end > len(encoding) {
            end = len(encoding)
        }

        for j := i; j < end; j += 2 {
            if j+1 < end {
                result[j], result[j+1] = result[j+1], result[j]
            }
        }
    }

    return result
}

func bigEndian32Bytes(encoding []byte) []byte {
    result := make([]byte, len(encoding))
    for i := 0; i < len(encoding); i++ {
        result[i] = encoding[i]
    }
    
    for i := 0; i < len(encoding); i += 32 {
        end := i + 32
        if end > len(encoding) {
            end = len(encoding)
        }
        
        for j := i; j < end; j += 4 {
            if j+3 < end {
                result[j], result[j+3] = result[j+3], result[j]
                result[j+1], result[j+2] = result[j+2], result[j+1]
            }
        }
    }
    
    return result
}