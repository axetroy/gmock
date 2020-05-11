package daemon

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/axetroy/gmock/internal/lib/fs"
)

type Action func() error

func getPidFilePath() (string, error) {
	cwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return path.Join(cwd, "gmock") + ".pid", nil
}

func Start(action Action, shouldRunInDaemon bool) error {
	if shouldRunInDaemon && os.Getppid() != 1 {
		// 将命令行参数中执行文件路径转换成可用路径
		filePath, _ := filepath.Abs(os.Args[0])

		args := os.Args[1:]

		var params []string

		for _, arg := range args {
			if arg != "-daemon" {
				params = append(params, arg)
			}
		}

		fmt.Println(filePath, params)

		cmd := exec.Command(filePath, params...)
		// 将其他命令传入生成出的进程
		//cmd.Stdin = os.Stdin // 给新进程设置文件描述符，可以重定向到文件中
		//cmd.Stdout = os.Stdout
		//cmd.Stderr = os.Stderr
		//cmd.Stdout = ioutil.Discard
		//cmd.Stderr = ioutil.Discard
		err := cmd.Start() // 开始执行新进程，不等待新进程退出
		return err
	} else {
		pidFilePath, err := getPidFilePath()

		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(pidFilePath, []byte(fmt.Sprintf("%d", os.Getpid())), os.ModePerm); err != nil {
			return err
		}

		log.Println(fmt.Sprintf("Generate the pid file %s", pidFilePath))

		return action()
	}
}

func Stop() error {
	pidFilePath, err := getPidFilePath()

	if err != nil {
		return err
	}

	if exist, err := fs.PathExists(pidFilePath); err != nil {
		return err
	} else if !exist {
		return nil
	}

	b, err := ioutil.ReadFile(pidFilePath)

	if err != nil {
		return nil
	}

	pidStr := string(b)

	pid, err := strconv.Atoi(pidStr)

	if err != nil {
		return err
	}

	ps, err := os.FindProcess(pid)

	if err != nil {
		return err
	}

	if err := ps.Signal(syscall.SIGTERM); err != nil {
		return err
	}

	psState, err := ps.Wait()

	if err != nil {
		return err
	}

	haveBeenKill := psState.Exited()

	if haveBeenKill {
		log.Printf("进程 %d 已结束.\n", psState.Pid())

		_ = os.RemoveAll(pidFilePath)
	} else {
		log.Printf("进程 %d 结束失败.\n", psState.Pid())
	}

	return nil
}
