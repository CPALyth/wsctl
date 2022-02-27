package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	var file string
	flag.StringVar(&file, "f", "user.proto", "填写proto文件全名")
	flag.Parse()

	//cmd = exec.Command("protoc", "--go_out=./", "--go-grpc_out=./", file)
	//err = cmd.Run()
	//fmt.Println("生成pb.go:", err)
	//
	//time.Sleep(time.Second * 4)

	//var cmd *exec.Cmd
	//var err error
	//in := bytes.NewBuffer(nil)
	//cmd = exec.Command("sh")
	//cmd.Stdin = in
	//
	//command := fmt.Sprintf("protoc --go_out=./ --go-grpc_out=./ %s && "+
	//	"ls *.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'\n", file)
	//in.WriteString(command)
	//in.WriteString("exit\n")
	//
	//if err = cmd.Run(); err != nil {
	//	fmt.Println(err)
	//}

	cmd := exec.Command("ls")

	stdout, err := cmd.StdoutPipe()
	if err != nil { //获取输出对象，可以从该对象中读取输出结果
		log.Fatal(err)
	}
	defer stdout.Close() // 保证关闭输出流

	if err := cmd.Run(); err != nil { // 运行命令
		log.Fatal(err)
	}

	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
		log.Fatal(err)
	} else {
		log.Println(string(opBytes))
	}
}
