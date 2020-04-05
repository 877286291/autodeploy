package utils

import (
	"fmt"
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

var (
	sftpClient   *sftp.Client
	auth         []ssh.AuthMethod
	addr         string
	clientConfig *ssh.ClientConfig
	sshClient    *ssh.Client
	err          error
)

func init() {
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(setting.FtpPassWord))

	clientConfig = &ssh.ClientConfig{
		User:            setting.FtpUserName,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //ssh.FixedHostKey(hostKey),
	}
	// connet to ssh
	addr = fmt.Sprintf("%s:%d", setting.FtpHost, 22)
	sshClient, err = ssh.Dial("tcp", addr, clientConfig)
	if err != nil {
		log.Println("连接FTP服务器时发生错误，请检查网络连接！")
		log.Println(err)
		os.Exit(1)
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		log.Println("连接FTP服务器时发生错误，请检查！")
		log.Println(err)
		os.Exit(1)
	}
}

func Upload(localPath string, remotePath string) (err error) {
	//获取路径的属性
	var s os.FileInfo
	s, err = os.Stat(localPath)
	if err != nil {
		//fmt.Println("文件路径不存在")
		return
	}
	_ = sftpClient.MkdirAll(remotePath)
	//判断是否是文件夹
	if s.IsDir() {
		uploadDirectory(localPath, remotePath)
	} else {
		uploadFile(localPath, remotePath)
	}
	return
}

func uploadFile(localFilePath string, remotePath string) {
	//打开本地文件流
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		fmt.Println("os.Open error : ", localFilePath)
		_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
		log.Fatal(err)
	}
	//关闭文件流
	defer srcFile.Close()
	//上传到远端服务器的文件名,与本地路径末尾相同
	var remoteFileName = path.Base(localFilePath)
	//打开远程文件,如果不存在就创建一个
	dstFile, err := sftpClient.Create(path.Join(remotePath, remoteFileName))
	if err != nil {
		fmt.Println("sftpClient.Create error : ", path.Join(remotePath, remoteFileName))
		_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
		log.Fatal(err)
	}
	//关闭远程文件
	defer dstFile.Close()
	//读取本地文件,写入到远程文件中
	fileInfo, _ := os.Stat(localFilePath)
	//文件大小
	fileSize := fileInfo.Size()
	//文件大于200M采用buffer
	if fileSize > 100*1024^2 {
		//一次复制10M
		buffer := make([]byte, 10*1024^2)
		_, err := io.CopyBuffer(dstFile, srcFile, buffer)
		if err != nil {
			_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
			log.Println(err)
		}
	} else {
		ff, err := ioutil.ReadAll(srcFile)
		if err != nil {
			fmt.Println("ReadAll error : ", localFilePath)
			_ = ioutil.WriteFile("err.log", []byte(err.Error()), 0777)
		}
		_, _ = dstFile.Write(ff)
	}
}

func uploadDirectory(localPath string, remotePath string) {
	//打开本地文件夹流
	localFiles, err := ioutil.ReadDir(localPath)
	if err != nil {
		log.Fatal("路径错误 ", err)
	}
	//先创建最外层文件夹
	_ = sftpClient.Mkdir(remotePath)
	//遍历文件夹内容
	for _, backupDir := range localFiles {
		localFilePath := path.Join(localPath, backupDir.Name())
		remoteFilePath := path.Join(remotePath, backupDir.Name())
		//判断是否是文件,是文件直接上传.是文件夹,先远程创建文件夹,再递归复制内部文件
		if backupDir.IsDir() {
			_ = sftpClient.Mkdir(remoteFilePath)
			uploadDirectory(localFilePath, remoteFilePath)
		} else {
			uploadFile(path.Join(localPath, backupDir.Name()), remotePath)
		}
	}
}
