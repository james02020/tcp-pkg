### 一个处理TCP分包、合包的实例（固定包头+包体协议）

#### 用法
    
    r, _ := tcp_package.NewReader(conn, maxBufferSize, headerLen, lengthOffset)

	r.Do()
	
	go func(acceptData chan string) {
    	for {
    		value, isOk := <-acceptData
    		if !isOk {
    			break
    		}
    		fmt.Println(value)
    	}
    }(r.Message)