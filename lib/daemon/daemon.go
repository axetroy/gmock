package daemon

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	path "path/filepath"
)

type Action func() error

func getPidFilePath() (string, error) {
	cwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	executableFilePath, err := os.Executable()

	executableName := path.Base(executableFilePath)

	if err != nil {
		return "", err
	}

	return path.Join(cwd, executableName) + ".pid", nil
}

func Start(action Action, shouldRunInDaemon bool) error {
	if shouldRunInDaemon && os.Getppid() != 1 {
		pidFilePath, err := getPidFilePath()

		if err != nil {
			return err
		}

		if _, err := os.Stat(pidFilePath); err == nil {
			log.Fatalf("已经存在进程，如果进程没有启动，请删除文件 `%s`\n", pidFilePath)
		}

		// 将命令行参数中执行文件路径转换成可用路径
		filePath, _ := path.Abs(os.Args[0])
		cmd := exec.Command(filePath, os.Args[1:]...)
		// 将其他命令传入生成出的进程
		// cmd.Stdin = os.Stdin // 给新进程设置文件描述符，可以重定向到文件中
		//cmd.Stdout = ioutil.Discard
		//cmd.Stderr = ioutil.Discard
		// 开始执行新进程，不等待新进程退出
		if err := cmd.Start(); err != nil {
			return err
		} else {
			fmt.Printf("启动守护进程 %d. 文件 `%s`\n", cmd.Process.Pid, pidFilePath)
			os.Exit(0)
			return nil
		}
	} else {
		pidFilePath, err := getPidFilePath()

		if err != nil {
			return err
		}

		pid := os.Getpid()

		if err := ioutil.WriteFile(pidFilePath, []byte(fmt.Sprintf("%d", pid)), 0600); err != nil {
			return err
		}

		return action()
	}
}
