package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"syscall"
	"time"
)

func daemon(nochdir, noclose int) int {
	var ret, ret2 uintptr
	var err syscall.Errno
	darwin := runtime.GOOS == "darwin"
	// 如果当前进程的ppid等于1，则表明已经是一个守护进程了，直接返回即可(already a daemon)
	if syscall.Getppid() == 1 {
		return 0
	}
	// 通过 syscall 包调用 fork 派生子进程
	ret, ret2, err = syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	if err != 0 {
		log.Println("syscall.RawSyscall")
		return -1
	}
	// ret2 表示子进程收到的 fork 返回（暂不完全确定？），正常派生情况下 ret2 应该等于0
	if ret2 < 0 {
		os.Exit(-1)
	}
	// 针对 darwin 系统的错误处理（没啥用，在macOS上还是跑不了）
	if darwin && ret2 == 1 {
		log.Printf("OS: %v, ret: %v, ret2: %v\n", darwin, ret, ret2)
		// ret = 0
	}
	// ret 表示fork成功执行后产生的子进程的pid（由父进程接收）
	if ret > 0 {
		os.Exit(0) // 退出当前的父进程
	}

	// create a new SID for the child process
	s_ret, s_errno := syscall.Setsid()
	if s_errno != nil {
		log.Printf("Error: syscall.Setsid errno: %d", s_errno)
	}
	if s_ret < 0 {
		log.Println("syscall.Setsid")
		return -1
	}

	// 调用 fork 派生子进程
	ret, ret2, err = syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	if err != 0 {
		log.Println("syscall.RawSyscall")
		return -1
	}
	// ret2 表示子进程收到的 fork 返回，正常派生情况下 ret2 应该等于0
	if ret2 < 0 {
		os.Exit(-1)
	}
	// ret 表示fork成功执行后产生的子进程的pid（由父进程接收）
	if ret > 0 {
		os.Exit(0) // 退出当前的父进程
	}

	if nochdir == 0 {
		os.Chdir("/")
	}
	/* Change the file mode mask */
	_ = syscall.Umask(0)

	if noclose == 0 {
		f, e := os.OpenFile("/dev/null", os.O_RDWR, 0)
		if e == nil {
			fd := f.Fd()
			syscall.Dup2(int(fd), int(os.Stdin.Fd()))
			syscall.Dup2(int(fd), int(os.Stdout.Fd()))
			syscall.Dup2(int(fd), int(os.Stderr.Fd()))
		}
	}
	return 0
}

func main() {
	daemon(0, 1)
	for {
		fmt.Println("hello")
		time.Sleep(2 * time.Second)
	}
}
